# 使用官方Go基础镜像
FROM golang:latest AS builder

# 设置工作目录
WORKDIR /app

# 将项目源码复制到容器内工作目录
COPY . .

# 编译Go应用，假设主包在main.go中，生成名为my-go-app的可执行文件
RUN go build -o my-go-app main.go

# 使用精简的基础镜像作为运行时环境
FROM alpine:latest

# 设置必要的环境变量（如有需要）
ENV APP_ENV=production

# 复制编译好的可执行文件到新的镜像
COPY --from=builder /app/my-go-app /usr/local/bin/

# 指定容器启动时运行的命令
CMD ["my-go-app"]