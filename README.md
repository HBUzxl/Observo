# Observo

项目架构
```
observo/
├── cmd/                  # 各种程序入口（agent / server）
│   ├── agent/            # 单机版 agent（指标采集 + 上报/展示）
│   │   └── main.go
│   └── server/           # 多机聚合服务端（接收 agent 上报 + 提供仪表盘）
│       └── main.go
│
├── internal/             # 项目核心逻辑（模块封装，外部不可依赖）
│   ├── metrics/          # 系统资源采集逻辑（CPU/内存/磁盘等）
│   │   └── collector.go
│   ├── transport/        # Agent 和 Server 的通信（HTTP / gRPC）
│   │   ├── http.go
│   │   └── grpc.go
│   ├── dashboard/        # 本地仪表盘的 HTTP 渲染逻辑
│   │   └── handler.go
│   ├── storage/          # 多机模式下的临时存储（内存 / Redis / BoltDB）
│   │   └── memory_store.go
│   └── config/           # 配置文件加载（使用 viper / koanf）
│       └── config.go
│
├── api/                  # OpenAPI spec 或 proto 文件（如使用 gRPC）
│   └── observometrics.proto
│
├── web/                  # 仪表盘前端（Vue/React/HTML + JS）
│   ├── index.html
│   ├── assets/
│   └── js/
│
├── scripts/              # 自动部署脚本 / 开发工具脚本
│   └── install_agent.sh
│
├── configs/              # 示例配置文件（YAML）
│   ├── agent.yaml
│   └── server.yaml
│
├── pkg/                  # 可被复用的通用工具（日志、中间件等）
│   ├── logger/
│   └── utils/
│
├── Dockerfile            # 用于容器部署
├── Makefile              # 编译、打包、构建命令（或使用 Taskfile）
├── README.md
└── go.mod
```