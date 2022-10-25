# grpc_demo

## 1. 安装protoc,下载对应系统的版本

```text
https://github.com/protocolbuffers/protobuf/releases
```


## 2. 安装protoc-gen-go和protoc-gen-go-grpc

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest //安装grpc-gateway
```

## 3. 同目录下下载utools工具包
```text
git clone https://github.com/hwdhy/utools.git
```

## 4. 项目启动

- 本地创建pgsql, 创建hwdhy数据库
```shell
docker run --name mypostgres -d -p 5432:5432 -e POSTGRES_PASSWORD=123456 postgres
```
- make run 启动服务（make 需提前安装 无make使用go run ./cmd/serve/main.go启动）

```shell
make run_grpc
```