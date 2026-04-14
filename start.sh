#!/bin/sh

# 安装curl用于健康检查
apk add --no-cache curl

# 等待 MySQL 就绪
echo "等待MySQL数据库就绪..."
until mysqladmin ping -h mysql -u root -pmima41 --silent; do
  echo "MySQL数据库不可用 -等待中"
  sleep 2
done
echo "MySQL数据库已启动并运行!"

# 启动 Nginx
nginx

# 启动后端服务
cd /app/backend && ./cms