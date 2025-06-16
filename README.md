# project
这是一个测试的项目个人使用

## elasticsearch
pkg目前封装了es的操作
具体使用方式可以参考 "common/cmd/createEs/main.go"

## 通义千问
目前项目中还有关于阿里云通义千问的调用
代码地址为 "api/handle/Alichatgpt.go"
路由地址为 "api/main.go“
关于流式输出的文件会在 "api/files" 

## mongodb
文档编写了go操作mongodb的增删查改
代码地址为 "common/cmd/mongo/main.go"

## jwt
文档有一个简易的jwt生成token
代码地址为 "common/pkr/jwt.go"

## go项目编译(在cmd进行实现)
### 交叉编译到Linux
set GOOS=linux
set GOARCH=amd64
go build -o app-linux

### 交叉编译到Mac
set GOOS=darwin
set GOARCH=amd64
go build -o app-mac