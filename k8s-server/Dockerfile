# 使用官方的 Golang 镜像作为基础镜像
FROM arm64v8/golang:1.21

# 设置必要的环境变量
# ENV GO111MODULE=on \
    # CGO_ENABLED='1' \
    # GOOS='darwin' \
    # GOARCH='arm64' \
    # GOPROXY='https://goproxy.cn,direct'

# 在容器内创建一个目录来存放我们的应用代码
RUN mkdir /app

# 将工作目录切换到 /app
WORKDIR /app

# 将当前目录下的所有文件拷贝到 /app 目录
COPY k8s-server .

# 编译 Go 应用程序
# RUN go build -o k8sserver .

# 暴露 8080 端口
EXPOSE 9090 8082

# 运行应用程序
CMD ["./k8s-server"]

