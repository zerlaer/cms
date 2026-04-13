#!/bin/sh

# 等待 MySQL 就绪
echo "Waiting for MySQL to be ready..."
until mysqladmin ping -h mysql -u root -pmima41 --silent; do
  echo "MySQL is unavailable - sleeping"
  sleep 2
done
echo "MySQL is up and running!"

# 启动 Nginx
nginx

# 启动后端服务
cd /app/backend && ./cms