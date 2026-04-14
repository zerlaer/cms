@echo off
echo ========================================
echo 电子元件管理系统 - 初始化脚本
echo ========================================
echo.

echo [1/2] 安装后端依赖...
cd backend
call go mod tidy
if errorlevel 1 (
    echo 错误：安装后端依赖失败
    pause
    exit /b 1
)
echo 后端依赖安装完成
cd ..
echo.

echo [2/2] 安装前端依赖...
cd frontend
call npm install
if errorlevel 1 (
    echo 错误：安装前端依赖失败
    pause
    exit /b 1
)
echo 前端依赖安装完成
cd ..
echo.

echo ========================================
echo 初始化完成！
echo 请执行以下操作：
echo 1. 创建数据库：mysql -u root -p < backend/schema.sql
echo 2. 修改配置文件：backend/config.yaml
echo 3. 运行启动脚本：start.bat
echo ========================================
echo.
pause
