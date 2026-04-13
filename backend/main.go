package main

import (
	"cms/internal/config"
	"cms/internal/database"
	"cms/internal/handlers"
	"cms/internal/middleware"
	"cms/internal/service"
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

	cfg := config.GetConfig()

	// 配置优化的日志格式，使用自定义的控制台编码器
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

	// 创建核心和日志器
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zapcore.InfoLevel,
	)

	// 启用 caller 信息，显示文件路径和行号
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()

	if err := database.Init(); err != nil {
		logger.Fatal("数据库连接失败 " + err.Error())
	}

	db := database.GetDB()

	componentService := service.NewComponentService(db)
	excelService := service.NewExcelService(db)

	componentHandler := handlers.NewComponentHandler(componentService)
	excelHandler := handlers.NewExcelHandler(excelService)

	gin.SetMode(cfg.Server.Mode)
	r := gin.New()
	r.Use(middleware.Logger(logger))
	r.Use(middleware.CORS())

	api := r.Group("/api")
	{
		components := api.Group("/components")
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

		excel := api.Group("/excel")
		{
			excel.GET("/export", excelHandler.Export)
			excel.POST("/import", excelHandler.Import)
		}
	}

	r.Static("/uploads", "./uploads")

	port := fmt.Sprintf(":%d", cfg.Server.Port)
	logger.Info("服务器启动，端口" + port)

	if err := r.Run(port); err != nil {
		logger.Fatal("服务器启动失败 " + err.Error())
	}
}
