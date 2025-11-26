# 极客待办 (GeekTODO)

> 基于 Go-zero + Vue3 + TypeScript 的现代化微服务待办事项管理系统

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)
![Vue Version](https://img.shields.io/badge/vue-3.5+-green.svg)
![TypeScript](https://img.shields.io/badge/typescript-5.9+-blue.svg)

## 📖 项目简介

极客待办是一个功能完整的现代化待办事项管理系统，采用微服务架构设计，提供高效的任务管理和分类功能。项目前后端完全分离，具备完善的用户认证体系和响应式用户界面。

### ✨ 核心特性

- 🔐 **完整认证体系**: JWT + Refresh Token 机制
- 🏗️ **微服务架构**: Go-zero + gRPC + etcd 服务发现
- 📱 **响应式设计**: 支持桌面端和移动端
- 🗂️ **任务分类管理**: 自定义分类和颜色标识
- ⚡ **高性能**: Redis 缓存 + 数据库优化
- 🛡️ **类型安全**: 全栈 TypeScript 支持
- 🎨 **现代UI**: Element Plus 组件库
- 🔒 **数据隔离**: 基于用户ID的严格权限控制

## 🏛️ 系统架构

### 整体架构图

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   前端 Vue3     │    │  API 网关        │    │   微服务集群    │
│   :3000         │◄──►│   :8888         │◄──►│  User :9000     │
│                 │    │                 │    │  Task :9001     │
│ - 任务管理       │    │ - 路由转发       │    │                 │
│ - 分类管理       │    │ - JWT验证       │    │ - gRPC通信      │
│ - 用户界面       │    │ - CORS处理      │    │ - etcd发现      │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                                        │
                       ┌─────────────────┐              │
                       │   基础设施       │◄─────────────┘
                       │                 │
                       │ MySQL :3306     │
                       │ Redis :6379     │
                       │ etcd  :2379     │
                       └─────────────────┘
```

### 技术栈

#### 后端技术栈
- **框架**: [Go-zero](https://go-zero.dev/) v1.9.3 - 云原生微服务框架
- **数据库**: MySQL 8.0+ - 主数据存储
- **缓存**: Redis - 分布式缓存
- **服务发现**: etcd - 微服务注册与发现
- **通信**: gRPC - 服务间通信协议
- **认证**: JWT - JSON Web Token
- **API文档**: Swagger - 自动生成的API文档

#### 前端技术栈
- **框架**: Vue 3.5.22 + TypeScript 5.9.0
- **构建工具**: Vite 7.1.11
- **UI组件**: Element Plus 2.9.7
- **状态管理**: Pinia 3.0.3
- **路由**: Vue Router 4.6.3
- **HTTP客户端**: Axios 1.8.2
- **开发工具**: ESLint + Prettier + Vitest

## 📁 项目结构

```
jike_todolist/
├── backend/                    # 后端微服务
│   ├── common/                # 公共模块
│   │   ├── ctxdata/          # 上下文数据管理
│   │   └── tool/             # 工具函数
│   ├── gateway/              # API网关服务
│   │   └── cmd/api/          # HTTP API服务 (端口:8888)
│   ├── task/                 # 任务管理服务
│   │   └── cmd/rpc/          # gRPC服务 (端口:9001)
│   └── user/                 # 用户管理服务
│       └── cmd/rpc/          # gRPC服务 (端口:9000)
├── frontend/                  # 前端Vue3应用
│   ├── src/
│   │   ├── api/              # API接口定义
│   │   ├── components/       # 可复用组件
│   │   │   ├── categories/   # 分类相关组件
│   │   │   ├── tasks/        # 任务相关组件
│   │   │   └── user/         # 用户相关组件
│   │   ├── router/           # 路由配置
│   │   ├── stores/           # Pinia状态管理
│   │   ├── views/            # 页面视图
│   │   │   ├── auth/         # 认证页面
│   │   │   ├── tasks/        # 任务管理页面
│   │   │   ├── categories/   # 分类管理页面
│   │   │   └── user/         # 用户设置页面
│   │   └── types.ts          # TypeScript类型定义
│   ├── public/               # 静态资源
│   └── dist/                 # 构建产物
├── docs/                      # 项目文档
├── prd/                       # 产品需求文档
├── .gitignore                 # Git忽略配置
├── docker-compose.yml        # Docker编排文件
└── README.md                  # 项目说明文档
```

## 🚀 快速开始

### 环境要求

- Go 1.21+
- Node.js 18+
- MySQL 8.0+
- Redis 6.0+
- etcd 3.5+

### 1. 克隆项目

```bash
git clone https://github.com/your-username/jike_todolist.git
cd jike_todolist
```

### 2. 数据库初始化

```sql
# 创建数据库
CREATE DATABASE todolist CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

# 创建用户表
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    nickname VARCHAR(50),
    avatar VARCHAR(255),
    status INT DEFAULT 1 COMMENT '1:正常 0:禁用',
    last_login_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

# 创建任务表
CREATE TABLE tasks (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    category VARCHAR(50) DEFAULT 'default',
    priority INT DEFAULT 1 COMMENT '1:低 2:中 3:高',
    status INT DEFAULT 0 COMMENT '0:未完成 1:已完成 2:已删除',
    due_date DATE,
    completed_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_user_id (user_id),
    INDEX idx_status (status),
    INDEX idx_category (category)
);

# 创建分类表
CREATE TABLE categories (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    name VARCHAR(50) NOT NULL,
    color VARCHAR(7) DEFAULT '#1890ff',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY uk_user_category (user_id, name)
);
```

### 3. 启动基础服务

```bash
# 启动 etcd
etcd --data-dir=/tmp/etcd-data --listen-peer-urls http://localhost:2380 --listen-client-urls http://localhost:2379 --advertise-client-urls http://localhost:2379

# 启动 Redis
redis-server

# 启动 MySQL
# 确保 MySQL 运行在 localhost:3306
```

### 4. 启动后端服务

```bash
# 启动用户服务
cd backend/user/cmd/rpc
go run user.go

# 启动任务服务
cd backend/task/cmd/rpc
go run task.go

# 启动API网关
cd backend/gateway/cmd/api
go run gateway.go
```

### 5. 启动前端服务

```bash
cd frontend
pnpm install
pnpm dev
```

### 6. 访问应用

- 前端应用: http://localhost:3000
- API网关: http://localhost:8888
- Swagger文档: http://localhost:8888/swagger/

## 📖 API 文档

### 认证接口

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | `/api/v1/auth/register` | 用户注册 |
| POST | `/api/v1/auth/login` | 用户登录 |
| POST | `/api/v1/auth/logout` | 用户登出 |
| POST | `/api/v1/auth/refresh` | 刷新Token |

### 用户管理

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | `/api/v1/user/profile` | 获取用户信息 |
| PUT | `/api/v1/user/profile` | 更新用户信息 |
| PUT | `/api/v1/user/password` | 修改密码 |

### 任务管理

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | `/api/v1/tasks` | 获取任务列表 |
| POST | `/api/v1/tasks` | 创建任务 |
| GET | `/api/v1/tasks/{id}` | 获取任务详情 |
| PUT | `/api/v1/tasks/{id}` | 更新任务 |
| DELETE | `/api/v1/tasks/{id}` | 删除任务 |
| PATCH | `/api/v1/tasks/{id}/status` | 更新任务状态 |

### 分类管理

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | `/api/v1/categories` | 获取分类列表 |
| POST | `/api/v1/categories` | 创建分类 |
| PUT | `/api/v1/categories/{id}` | 更新分类 |
| DELETE | `/api/v1/categories/{id}` | 删除分类 |

## 🔧 配置说明

### 后端配置

#### Gateway API 配置 (`backend/gateway/cmd/api/etc/gateway-api.yaml`)

```yaml
Name: gateway-api
Host: 0.0.0.0
Port: 8888

# RPC服务配置
UserRpcClient:
  Etcd:
    Hosts:
    - 127.0.0.1:2379
    Key: user.rpc

TaskRpcClient:
  Etcd:
    Hosts:
    - 127.0.0.1:2379
    Key: task.rpc

# JWT配置
JwtAuth:
  AccessSecret: your-access-secret
  AccessExpire: 31536000
```

#### 数据库配置

修改对应服务的配置文件中的数据库连接信息：

```yaml
DataSource: username:password@tcp(localhost:3306)/todolist
```

### 前端配置

#### 环境变量配置

```bash
# .env.development
VITE_API_BASE_URL=http://localhost:8888/api/v1
VITE_APP_TITLE=极客待办
```

## 🧪 开发指南

### 后端开发

#### 添加新的API接口

1. 在 `backend/gateway/cmd/api/internal/handler/` 中添加处理器
2. 在 `backend/gateway/cmd/api/internal/logic/` 中实现业务逻辑
3. 在 `backend/gateway/cmd/api/internal/types/` 中定义类型
4. 在 `backend/gateway/cmd/api/internal/handler/routes.go` 中注册路由

#### 添加新的RPC服务

1. 在 `backend/{service}/cmd/rpc/desc/` 中定义 proto 文件
2. 生成 gRPC 代码: `protoc --go_out=. --go-grpc_out=. desc/{service}.proto`
3. 实现业务逻辑: `backend/{service}/cmd/rpc/internal/logic/`
4. 启动服务: `go run {service}.go`

### 前端开发

#### 添加新页面

1. 在 `frontend/src/views/` 中创建页面组件
2. 在 `frontend/src/router/index.ts` 中添加路由配置
3. 在 `frontend/src/api/` 中添加API接口定义
4. 在 `frontend/src/stores/` 中添加状态管理

#### 代码规范

- 使用 TypeScript 进行类型检查
- 遵循 ESLint 和 Prettier 配置
- 组件命名使用 PascalCase
- 文件命名使用 kebab-case

## 📊 系统监控

### 日志查看

```bash
# Gateway服务日志
tail -f backend/gateway/cmd/api/logs/gateway-api.log

# 用户服务日志
tail -f backend/user/cmd/rpc/logs/user.log

# 任务服务日志
tail -f backend/task/cmd/rpc/logs/task.log
```

### 性能监控

- **数据库监控**: 使用 `SHOW PROCESSLIST` 查看连接状态
- **Redis监控**: 使用 `redis-cli monitor` 监控命令执行
- **API监控**: 查看网关日志监控请求响应时间

## 🚀 部署指南

### Docker 部署

```bash
# 构建镜像
docker-compose build

# 启动服务
docker-compose up -d

# 查看服务状态
docker-compose ps
```

### 生产环境配置

1. **安全配置**:
   - 更换默认JWT密钥
   - 配置HTTPS证书
   - 设置防火墙规则

2. **性能优化**:
   - 数据库索引优化
   - Redis连接池配置
   - 负载均衡配置

3. **监控告警**:
   - 应用性能监控
   - 错误日志告警
   - 资源使用监控

## 🤝 贡献指南

1. Fork 项目
2. 创建特性分支: `git checkout -b feature/new-feature`
3. 提交更改: `git commit -m 'Add new feature'`
4. 推送分支: `git push origin feature/new-feature`
5. 提交 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 🙏 致谢

- [Go-zero](https://go-zero.dev/) - 云原生微服务框架
- [Vue.js](https://vuejs.org/) - 渐进式JavaScript框架
- [Element Plus](https://element-plus.org/) - Vue 3 UI组件库
- [MySQL](https://www.mysql.com/) - 关系型数据库
- [Redis](https://redis.io/) - 内存数据库

