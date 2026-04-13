package main

import (
	"cms/internal/config"
	"cms/internal/database"
	"cms/internal/handlers"
	"cms/internal/middleware"
	"cms/internal/models"
	"cms/internal/service"
	"cms/internal/utils"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	if err := config.Init("config.yaml"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	utils.InitJWTSecret()

	cfg := config.GetConfig()

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zapcore.InfoLevel,
	)

	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()

	if err := database.Init(); err != nil {
		logger.Fatal("Database connection failed " + err.Error())
	}

	db := database.GetDB()

	db.AutoMigrate(&models.Component{}, &models.StockRecord{}, &models.User{})

	componentService := service.NewComponentService(db)
	excelService := service.NewExcelService(db)
	userService := service.NewUserService(db)

	if err := userService.InitDefaultUser(); err != nil {
		logger.Error("Failed to init default user", zap.Error(err))
	}

	componentHandler := handlers.NewComponentHandler(componentService)
	excelHandler := handlers.NewExcelHandler(excelService)
	userHandler := handlers.NewUserHandler(userService)

	gin.SetMode(cfg.Server.Mode)
	r := gin.New()
	r.Use(middleware.Logger(logger))
	r.Use(middleware.CORS())

	api := r.Group("/api")
	{
		api.POST("/auth/login", userHandler.Login)

		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			components := protected.Group("/components")
			{
				components.GET("", componentHandler.GetList)
				components.GET("/:id", componentHandler.GetByID)
				components.POST("", componentHandler.Create)
				components.PUT("/:id", componentHandler.Update)
				components.DELETE("/:id", componentHandler.Delete)
				components.POST("/:id/stock-in", componentHandler.StockIn)
				components.POST("/:id/stock-out", componentHandler.StockOut)
				components.GET("/:id/records", componentHandler.GetStockRecords)
			}

			excel := protected.Group("/excel")
			{
				excel.GET("/export", excelHandler.Export)
				excel.POST("/import", excelHandler.Import)
			}

			users := protected.Group("/users")
			{
				users.GET("", userHandler.GetList)
				users.POST("", userHandler.Create)
				users.PUT("/:id", userHandler.Update)
				users.DELETE("/:id", userHandler.Delete)
				users.POST("/:id/change-password", userHandler.ChangePassword)
			}
		}
	}

	r.Static("/uploads", "./uploads")
	r.NoRoute(func(c *gin.Context) {
		c.Status(404)
	})

	port := fmt.Sprintf(":%d", cfg.Server.Port)
	logger.Info("Server started, port" + port)

	if err := r.Run(port); err != nil {
		logger.Fatal("Server startup failed " + err.Error())
	}
}
