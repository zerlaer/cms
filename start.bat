@echo off
title CMS - Start Services

echo ============================================
echo    CMS - Start Services
echo ============================================
echo.

echo [1/3] Starting backend service...
cd /d %~dp0backend
start "Backend" cmd /k "cd /d %~dp0backend && go run main.go"

echo Waiting for backend...
timeout /t 5 /nobreak >nul

echo [2/3] Starting frontend service...
cd /d %~dp0frontend
start "Frontend" cmd /k "cd /d %~dp0frontend && npm run dev"

echo.
echo ============================================
echo    Services started!
echo ============================================
echo.
echo    Backend:  http://localhost:8080
echo    Frontend: http://localhost:3000
echo.
echo    Press any key to exit...
echo ============================================
pause >nul
