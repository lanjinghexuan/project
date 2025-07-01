# 项目说明文档

## 一、项目简介

本项目为个人测试与学习用途，基于 Go 语言实现，集成了 Elasticsearch、MongoDB、MinIO、JWT、Redis、阿里云通义千问等多种常用后端技术，支持多种数据存储、AI接口调用、缓存、文件上传等功能，适合后端开发者参考和二次开发。

---

## 二、目录结构

```
project/
├── api/                # 后端API服务
│   ├── handle/         # 业务处理（如用户、商品分类、AI接口等）
│   ├── request/        # 请求参数结构体
│   ├── server/         # gRPC服务实现
│   ├── router/         # 路由注册
│   └── main.go         # API服务入口
├── client/             # gRPC客户端
│   ├── server/         # 客户端gRPC服务实现
│   └── main.go         # 客户端入口
├── cmd/                # 命令行工具
│   └── console/        # 控制台相关命令
├── common/             # 公共模块
│   ├── cmd/            # 各类命令工具（如ES、Mongo、通义千问）
│   ├── config/         # 配置文件与结构体
│   ├── gload/          # 全局变量与初始化
│   ├── init/           # 初始化相关
│   ├── model/          # 数据模型
│   ├── pkr/            # 工具库（如缓存、ES、MinIO、JWT等）
│   └── proto/          # gRPC协议文件
├── go.mod/go.sum       # Go依赖管理
└── README.md           # 项目说明文档
```

---

## 三、主要功能模块

### 1. Elasticsearch
- 封装了ES的索引、文档增删查改、批量操作等功能。
- 参考：`common/pkr/elasticsearch.go`、`common/pkr/aiElasticsearch.go`
- 示例：`common/cmd/createEs/main.go`

### 2. MongoDB
- 提供了MongoDB的基本操作示例（增删查改）。
- 参考：`common/pkr/mongodb.go`
- 示例：`common/cmd/mongo/main.go`

### 3. MinIO
- 支持文件上传，适合对象存储场景。
- 参考：`common/pkr/minio.go`

### 4. JWT
- 简易的JWT生成与解析，适合接口鉴权。
- 参考：`common/pkr/jwt.go`

### 5. 缓存
- 提供了基于Redis的接口缓存方法，支持自定义key、过期时间、强制刷新等。
- 参考：`common/pkr/cache.go`、`common/pkr/aicache.go`

### 6. 阿里云通义千问
- 支持AI对话、流式输出等功能。
- 参考：`api/handle/Alichatgpt.go`
- 路由注册：`api/main.go`
- 文件输出：`api/files/`

### 7. gRPC服务
- 商品分类、用户等服务的gRPC接口实现。
- 参考：`api/server/`、`client/server/`
- 协议文件：`common/proto/`

---

## 四、依赖环境

- Go 1.24+
- Elasticsearch 7/8
- MongoDB
- MinIO
- Redis
- 其他依赖见`go.mod`

---

## 五、编译与运行

### 1. 交叉编译
- 编译为Linux可执行文件：
  ```shell
  set GOOS=linux
  set GOARCH=amd64
  go build -o app-linux
  ```
- 编译为Mac可执行文件：
  ```shell
  set GOOS=darwin
  set GOARCH=amd64
  go build -o app-mac
  ```

### 2. 启动API服务
```shell
cd api
go run main.go
```

### 3. 启动gRPC客户端
```shell
cd client
go run main.go
```

### 4. 运行命令行工具
- Elasticsearch示例：`go run common/cmd/createEs/main.go`
- MongoDB示例：`go run common/cmd/mongo/main.go`
- 通义千问示例：`go run common/cmd/tongyiqianwen/qianwen.go`

---

## 六、配置说明

- 配置文件路径：`common/config/expland.config.yaml`
- 支持MySQL、Redis、Elasticsearch、MongoDB、MinIO、阿里云通义千问等配置项。
- 结构体定义见：`common/config/config.go`

---

## 七、常见问题与建议

- 各模块均有详细代码注释，适合参考和二次开发。
- 如需扩展新功能，建议在`common/pkr/`或`api/handle/`下新增模块。
- 配置文件注意与实际环境保持一致。

---

如需进一步完善文档或有其他定制需求，请随时告知！