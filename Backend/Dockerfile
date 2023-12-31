# 使用官方的Go映像作为基础映像
FROM golang:1.19 AS builder
# 设置Go代理为Goproxy
ENV GOPROXY=https://goproxy.io,direct
# 设置工作目录
WORKDIR /app
# 复制go mod和sum文件并下载所有依赖项
COPY go.mod go.sum ./
RUN go mod download
# 复制源代码到容器中
COPY . .
# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 使用debian作为基础映像
FROM debian:latest

# 从builder阶段复制执行文件到当前阶段
COPY --from=builder /app/main /main
COPY ./config/config.ini /config/config.ini

# 设置debian容器时间
ENV TZ=Asia/Shanghai

# 设置应用程序为Docker容器的默认执行命令
CMD ["/main"]