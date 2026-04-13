# 第一阶段：构建阶段
FROM golang:1.21-alpine AS builder

# 设置国内镜像源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 安装构建依赖
RUN apk add --no-cache git gcc g++

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 设置 Go 国内镜像源
ENV GOPROXY=https://goproxy.cn,direct

# 下载依赖
RUN go mod tidy

# 复制源代码
COPY . .

# 构建应用
RUN go build -o cms-backend .

# 第二阶段：运行阶段
FROM alpine:3.18

# 设置国内镜像源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 安装运行依赖
RUN apk add --no-cache tzdata

# 设置时区
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone

# 创建工作目录
WORKDIR /app

# 从构建阶段复制可执行文件
COPY --from=builder /app/cms-backend /app/
# 复制配置文件
COPY config.yaml /app/
# 创建必要的目录
RUN mkdir -p /app/uploads /app/template

# 暴露端口
EXPOSE 8080

# 启动应用
CMD ["./cms-backend"]