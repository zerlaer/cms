@echo off
title CMS - Stop Services

echo ============================================
echo    CMS - Stop Services
echo ============================================
echo.

echo Stopping backend...
taskkill /F /FI "WindowTitle eq Backend*" 2>nul

echo Stopping frontend...
taskkill /F /FI "WindowTitle eq Frontend*" 2>nul

echo.
echo ============================================
echo    Services stopped
echo ============================================
pause
