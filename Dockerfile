# 第一阶段：构建前端
FROM hub.baidubce.com/library/node:18-alpine AS frontend-builder

# 设置国内镜像源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 设置 npm 国内镜像源
RUN npm config set registry https://registry.npmmirror.com

# 设置工作目录
WORKDIR /app/frontend

# 复制前端 package.json 和 package-lock.json 文件
COPY frontend/package*.json ./

# 安装依赖
RUN npm install

# 复制前端源代码
COPY frontend/ .

# 构建前端应用
RUN npm run build

# 第二阶段：构建后端
FROM hub.baidubce.com/library/golang:1.21-alpine AS backend-builder

# 设置国内镜像源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 安装构建依赖
RUN apk add --no-cache git gcc g++

# 设置工作目录
WORKDIR /app/backend

# 复制后端 go.mod 和 go.sum 文件
COPY backend/go.mod backend/go.sum ./

# 设置 Go 国内镜像源
ENV GOPROXY=https://goproxy.cn,direct

# 下载依赖
RUN go mod tidy

# 复制后端源代码
COPY backend/ .

# 构建后端应用
RUN go build -o cms .

# 第三阶段：运行阶段
FROM hub.baidubce.com/library/alpine:3.18

# 设置国内镜像源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 安装运行依赖
RUN apk add --no-cache tzdata nginx

# 设置时区
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone

# 创建必要的目录
RUN mkdir -p /app/backend/uploads /app/backend/template /app/frontend /etc/nginx/conf.d

# 从前端构建阶段复制构建结果
COPY --from=frontend-builder /app/frontend/dist /app/frontend

# 从后端构建阶段复制可执行文件和配置文件
COPY --from=backend-builder /app/backend/cms /app/backend/
COPY backend/config.yaml /app/backend/

# 复制 nginx 配置文件
COPY frontend/nginx.conf /etc/nginx/conf.d/default.conf

# 暴露端口
EXPOSE 80 8080

# 启动脚本
COPY start.sh /app/
RUN chmod +x /app/start.sh

# 启动应用
CMD ["/app/start.sh"]