# Mac 下构建 Linux 可执行程序
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build