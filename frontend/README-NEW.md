# 极客TODO - 前端应用

基于Vue3 + TypeScript + Element Plus的现代化TODO管理应用，根据PRD要求实现。

## 🚀 快速开始

### 1. 安装依赖

```bash
pnpm install
```

### 2. 启动开发服务器

```bash
pnpm dev
```

应用将在 http://localhost:3000 启动

### 3. 构建生产版本

```bash
pnpm build
```

## 🛠 技术栈

- **框架**: Vue 3 (Composition API)
- **语言**: TypeScript
- **构建工具**: Vite
- **UI组件库**: Element Plus
- **状态管理**: Pinia
- **路由**: Vue Router 4
- **HTTP客户端**: Axios
- **样式**: CSS3 + Element Plus
- **日期处理**: Day.js
- **包管理器**: pnpm

## 📋 功能特性

### ✅ 已实现功能

#### 用户认证
- [x] 用户注册 - 支持用户名、邮箱、密码验证
- [x] 用户登录 - JWT认证，自动token管理
- [x] 用户登出 - 清除认证状态
- [x] 路由守卫 - 保护需要认证的页面
- [x] 自动登录状态维护

#### 任务管理
- [x] 任务创建 - 支持标题、描述、分类、优先级、截止时间
- [x] 任务列表 - 分页展示，支持状态和分类筛选
- [x] 任务编辑 - 部分字段更新功能
- [x] 任务删除 - 带确认提示的安全删除
- [x] 任务状态管理 - 待办/进行中/已完成三种状态
- [x] 任务详情 - 完整的任务信息展示
- [x] 任务搜索 - 实时搜索任务标题
- [x] 任务统计 - 待办、进行中、已完成数量统计

#### 分类管理
- [x] 创建分类 - 自定义分类名称和颜色
- [x] 分类列表 - 展示所有用户分类
- [x] 分类编辑 - 修改分类名称和颜色
- [x] 分类删除 - 带确认提示的删除功能

#### 个人资料
- [x] 基本信息展示 - 用户名、邮箱、注册时间等
- [x] 编辑个人资料 - 修改用户名和邮箱
- [x] 修改密码 - 安全的密码更新功能
- [x] 账户统计 - 任务数量、分类数量等统计信息

#### 界面设计
- [x] 响应式设计 - 适配不同屏幕尺寸
- [x] 现代化UI - Element Plus组件库
- [x] 侧边栏导航 - 可折叠的导航菜单
- [x] 中文本地化 - 完整的中文界面
- [x] 加载状态 - 友好的加载提示
- [x] 错误处理 - 完善的错误提示

## 📁 项目结构

```
frontend/
├── src/
│   ├── api/                    # API接口
│   │   ├── auth.ts            # 认证相关API
│   │   ├── task.ts            # 任务相关API
│   │   ├── types.ts           # 类型定义
│   │   └── request.ts         # HTTP请求封装
│   ├── components/            # 通用组件
│   │   ├── tasks/            # 任务组件
│   │   ├── categories/       # 分类组件
│   │   └── user/             # 用户组件
│   ├── router/                # 路由配置
│   │   └── index.ts
│   ├── stores/                # 状态管理
│   │   ├── user.ts           # 用户状态
│   │   └── task.ts           # 任务状态
│   ├── views/                 # 页面组件
│   │   ├── auth/             # 认证页面
│   │   ├── layout/           # 布局组件
│   │   ├── tasks/            # 任务页面
│   │   ├── categories/       # 分类页面
│   │   ├── user/             # 用户页面
│   │   └── error/            # 错误页面
│   ├── App.vue               # 根组件
│   └── main.ts               # 入口文件
├── package.json              # 项目配置
├── vite.config.ts            # Vite配置
└── README.md                 # 项目说明
```

## 🔧 API配置

前端默认连接到后端API：`http://localhost:8888/api`

如需修改，请编辑 `.env` 文件：
```env
VITE_API_BASE_URL=http://localhost:8888/api
```

## 🎯 核心页面

1. **登录页面** (`/login`) - 用户登录
2. **注册页面** (`/register`) - 用户注册
3. **任务列表** (`/dashboard/tasks`) - 任务管理主界面
4. **任务详情** (`/dashboard/tasks/:id`) - 单个任务详情
5. **分类管理** (`/dashboard/categories`) - 分类管理
6. **个人资料** (`/dashboard/profile`) - 用户信息管理

## 🔄 数据流程

1. **认证流程**
   - 用户登录/注册 → 获取JWT Token → 存储到localStorage → 自动添加到请求头

2. **任务管理**
   - 创建/编辑任务 → 调用后端API → 更新本地状态 → 刷新列表

3. **状态管理**
   - Pinia管理全局状态 → localStorage持久化 → 页面刷新自动恢复

## 🎨 UI特性

- **Element Plus组件** - 现代化的UI组件
- **响应式设计** - 适配手机、平板、桌面
- **暗色主题支持** - Element Plus内置主题
- **图标支持** - 完整的Element Plus图标库
- **加载状态** - 骨架屏和加载动画
- **错误提示** - 友好的错误信息展示

## 🔒 安全特性

- **JWT认证** - 安全的用户认证机制
- **路由守卫** - 保护需要认证的页面
- **请求拦截** - 自动添加认证token
- **响应处理** - 统一的错误处理和token刷新
- **表单验证** - 前端数据验证
- **密码加密** - 安全的密码传输

## 🚀 部署说明

### 开发环境
```bash
pnpm dev
```

### 生产构建
```bash
pnpm build
```

### 预览构建结果
```bash
pnpm preview
```

---

🎯 **根据PRD要求，完整实现了Vue3组合式API + TypeScript + pnpm的技术栈，提供了完整的TODO管理功能。**