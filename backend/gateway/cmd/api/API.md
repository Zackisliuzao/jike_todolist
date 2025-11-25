# TODO List Gateway API 文档

## 基础信息

- **Base URL**: `http://localhost:8888`
- **API Version**: `v1`
- **Content-Type**: `application/json`
- **Schemes**: `https`

---

## 认证相关 API

### 1. 用户注册
**POST** `/api/v1/auth/register`

**描述**: 注册新用户账号

**请求参数**:
```json
{
  "username": "string",      // 用户名 (必填, 3-50字符)
  "email": "string",         // 邮箱 (必填, 有效邮箱格式)
  "password": "string",      // 密码 (必填, 最少6位)
  "nickname": "string"       // 昵称 (必填, 最多50字符)
}
```

**返回值**:
```json
{
  "user": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "nickname": "测试用户",
    "avatar": "",
    "status": 1,
    "last_login_at": "2025-11-25T22:00:00Z",
    "created_at": "2025-11-25T21:00:00Z"
  },
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
  "expires_at": 1732579200
}
```

---

### 2. 用户登录
**POST** `/api/v1/auth/login`

**描述**: 用户名密码登录

**请求参数**:
```json
{
  "username": "string",      // 用户名 (必填)
  "password": "string"       // 密码 (必填)
}
```

**返回值**: 同用户注册

---

### 3. 刷新令牌
**POST** `/api/v1/auth/refresh`

**描述**: 使用刷新令牌获取新的访问令牌

**请求参数**:
```json
{
  "refresh_token": "string"  // 刷新令牌 (必填)
}
```

**返回值**: 同用户注册

---

### 4. 用户登出
**POST** `/api/v1/auth/logout`

**描述**: 用户登出，使令牌失效

**请求头**:
- `Authorization: Bearer <access_token>`

**返回值**:
```json
{
  "code": 0,
  "message": "登出成功",
  "data": null
}
```

---

## 用户管理 API (需要JWT认证)

### 5. 获取用户信息
**GET** `/api/v1/user/profile`

**描述**: 获取当前用户的详细信息

**请求头**:
- `Authorization: Bearer <access_token>`

**返回值**:
```json
{
  "code": 0,
  "message": "获取用户信息成功",
  "data": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "nickname": "测试用户",
    "avatar": "",
    "status": 1,
    "last_login_at": "2025-11-25T22:00:00Z",
    "created_at": "2025-11-25T21:00:00Z"
  }
}
```

---

### 6. 更新用户信息
**PUT** `/api/v1/user/profile`

**描述**: 更新当前用户的信息

**请求头**:
- `Authorization: Bearer <access_token>`

**请求参数**:
```json
{
  "nickname": "string",      // 昵称 (可选)
  "avatar": "string"         // 头像URL (可选)
}
```

**返回值**:
```json
{
  "code": 0,
  "message": "更新成功",
  "data": null
}
```

---

### 7. 修改密码
**PUT** `/api/v1/user/password`

**描述**: 修改用户密码

**请求头**:
- `Authorization: Bearer <access_token>`

**请求参数**:
```json
{
  "old_password": "string",  // 原密码 (必填)
  "new_password": "string"   // 新密码 (必填, 最少6位)
}
```

**返回值**:
```json
{
  "code": 0,
  "message": "密码修改成功",
  "data": null
}
```

---

## 任务管理 API (需要JWT认证)

### 8. 创建任务
**POST** `/api/v1/tasks`

**描述**: 创建新任务

**请求头**:
- `Authorization: Bearer <access_token>`

**请求参数**:
```json
{
  "title": "string",         // 任务标题 (必填, 最多255字符)
  "description": "string",   // 任务描述 (必填)
  "category": "string",      // 分类 (可选, 默认"default")
  "priority": 1,             // 优先级 (可选, 1:低, 2:中, 3:高, 默认1)
  "due_date": "string"       // 截止日期 (可选, ISO格式)
}
```

**返回值**:
```json
{
  "code": 0,
  "message": "创建成功",
  "data": null
}
```

---

### 9. 获取任务列表
**GET** `/api/v1/tasks`

**描述**: 获取分页的任务列表，支持过滤

**请求头**:
- `Authorization: Bearer <access_token>`

**查询参数**:
- `page` (int, 可选): 页码, 默认1
- `size` (int, 可选): 每页数量, 默认10
- `category` (string, 可选): 按分类过滤
- `status` (int, 可选): 按状态过滤 (0:待办, 1:完成, 2:进行中)
- `sort_by` (string, 可选): 排序字段, 默认"created_at"

**返回值**:
```json
{
  "code": 0,
  "message": "获取成功",
  "data": {
    "tasks": [
      {
        "id": 1,
        "user_id": 1,
        "title": "学习Go",
        "description": "完成Go-Zero教程",
        "category": "学习",
        "priority": 2,
        "status": 0,
        "due_date": "2025-12-01T00:00:00Z",
        "created_at": "2025-11-25T21:00:00Z",
        "updated_at": "2025-11-25T21:00:00Z",
        "completed_at": null
      }
    ],
    "total": 1
  }
}
```

---

### 10. 获取任务详情
**GET** `/api/v1/tasks/{id}`

**描述**: 根据ID获取任务详情

**请求头**:
- `Authorization: Bearer <access_token>`

**路径参数**:
- `id` (int): 任务ID

**返回值**:
```json
{
  "code": 0,
  "message": "获取成功",
  "data": {
    "id": 1,
    "user_id": 1,
    "title": "学习Go",
    "description": "完成Go-Zero教程",
    "category": "学习",
    "priority": 2,
    "status": 0,
    "due_date": "2025-12-01T00:00:00Z",
    "created_at": "2025-11-25T21:00:00Z",
    "updated_at": "2025-11-25T21:00:00Z",
    "completed_at": null
  }
}
```

---

### 11. 更新任务
**PUT** `/api/v1/tasks/{id}`

**描述**: 更新任务信息

**请求头**:
- `Authorization: Bearer <access_token>`

**路径参数**:
- `id` (int): 任务ID

**请求参数**:
```json
{
  "title": "string",         // 任务标题 (可选)
  "description": "string",   // 任务描述 (可选)
  "category": "string",      // 分类 (可选)
  "priority": 1,             // 优先级 (可选, 1-3)
  "due_date": "string"       // 截止日期 (可选)
}
```

**返回值**:
```json
{
  "code": 0,
  "message": "更新成功",
  "data": null
}
```

---

### 12. 更新任务状态
**PATCH** `/api/v1/tasks/{id}/status`

**描述**: 更新任务完成状态

**请求头**:
- `Authorization: Bearer <access_token>`

**路径参数**:
- `id` (int): 任务ID

**请求参数**:
```json
{
  "status": 1               // 状态 (必填, 0:待办, 1:完成)
}
```

**返回值**:
```json
{
  "code": 0,
  "message": "状态更新成功",
  "data": null
}
```

---

### 13. 删除任务
**DELETE** `/api/v1/tasks/{id}`

**描述**: 删除任务（软删除）

**请求头**:
- `Authorization: Bearer <access_token>`

**路径参数**:
- `id` (int): 任务ID

**返回值**:
```json
{
  "code": 0,
  "message": "删除成功",
  "data": null
}
```

---

### 14. 搜索任务
**GET** `/api/v1/tasks/search`

**描述**: 根据关键词搜索任务

**请求头**:
- `Authorization: Bearer <access_token>`

**查询参数**:
- `keyword` (string, 必填): 搜索关键词
- `page` (int, 可选): 页码, 默认1
- `size` (int, 可选): 每页数量, 默认10

**返回值**: 同获取任务列表

---

### 15. 批量更新任务状态
**PATCH** `/api/v1/tasks/batch/status`

**描述**: 批量更新多个任务的状态

**请求头**:
- `Authorization: Bearer <access_token>`

**请求参数**:
```json
{
  "task_ids": [1, 2, 3],    // 任务ID数组 (必填)
  "status": 1               // 状态 (必填, 0:待办, 1:完成)
}
```

**返回值**:
```json
{
  "code": 0,
  "message": "批量更新成功",
  "data": null
}
```

---

### 16. 批量删除任务
**DELETE** `/api/v1/tasks/batch`

**描述**: 批量删除多个任务

**请求头**:
- `Authorization: Bearer <access_token>`

**请求参数**:
```json
{
  "task_ids": [1, 2, 3]     // 任务ID数组 (必填)
}
```

**返回值**:
```json
{
  "code": 0,
  "message": "批量删除成功",
  "data": null
}
```

---

## 分类管理 API (需要JWT认证)

### 17. 创建分类
**POST** `/api/v1/categories`

**描述**: 创建新的任务分类

**请求头**:
- `Authorization: Bearer <access_token>`

**请求参数**:
```json
{
  "name": "string",          // 分类名称 (必填, 最多50字符)
  "color": "#1890ff"         // 分类颜色 (可选, 默认"#1890ff")
}
```

**返回值**:
```json
{
  "code": 0,
  "message": "创建成功",
  "data": null
}
```

---

### 18. 获取分类列表
**GET** `/api/v1/categories`

**描述**: 获取用户的分类列表

**请求头**:
- `Authorization: Bearer <access_token>`

**返回值**:
```json
{
  "code": 0,
  "message": "获取成功",
  "data": {
    "categories": [
      {
        "id": 1,
        "user_id": 1,
        "name": "工作",
        "color": "#ff4d4f",
        "created_at": "2025-11-25T21:00:00Z",
        "updated_at": "2025-11-25T21:00:00Z"
      }
    ]
  }
}
```

---

### 19. 更新分类
**PUT** `/api/v1/categories/{id}`

**描述**: 更新分类信息

**请求头**:
- `Authorization: Bearer <access_token>`

**路径参数**:
- `id` (int): 分类ID

**请求参数**:
```json
{
  "name": "string",          // 分类名称 (可选)
  "color": "#1890ff"         // 分类颜色 (可选)
}
```

**返回值**:
```json
{
  "code": 0,
  "message": "更新成功",
  "data": null
}
```

---

### 20. 删除分类
**DELETE** `/api/v1/categories/{id}`

**描述**: 删除分类

**请求头**:
- `Authorization: Bearer <access_token>`

**路径参数**:
- `id` (int): 分类ID

**返回值**:
```json
{
  "code": 0,
  "message": "删除成功",
  "data": null
}
```

---

## 统计 API (需要JWT认证)

### 21. 获取概览统计
**GET** `/api/v1/stats/overview`

**描述**: 获取用户的任务概览统计

**请求头**:
- `Authorization: Bearer <access_token>`

**返回值**:
```json
{
  "code": 0,
  "message": "获取成功",
  "data": {
    "total_tasks": 100,
    "completed_tasks": 80,
    "pending_tasks": 15,
    "completion_rate": 0.8
  }
}
```

---

### 22. 获取完成统计
**GET** `/api/v1/stats/completion`

**描述**: 获取详细的完成情况统计

**请求头**:
- `Authorization: Bearer <access_token>`

**返回值**:
```json
{
  "code": 0,
  "message": "获取成功",
  "data": {
    "daily_completion": [
      {
        "date": "2025-11-25",
        "completed": 5,
        "total": 8
      }
    ],
    "category_stats": [
      {
        "category": "工作",
        "count": 20,
        "rate": 0.75
      }
    ],
    "priority_stats": [
      {
        "priority": 3,
        "count": 10,
        "rate": 0.9
      }
    ]
  }
}
```

---

## 错误码说明

| 错误码 | 说明 |
|--------|------|
| 0 | 成功 |
| 401 | 未认证或令牌无效 |
| 403 | 权限不足 |
| 404 | 资源不存在 |
| 422 | 请求参数错误 |
| 500 | 服务器内部错误 |

---

## 请求示例

### 注册用户
```bash
curl -X POST http://localhost:8888/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "123456",
    "nickname": "测试用户"
  }'
```

### 创建任务
```bash
curl -X POST http://localhost:8888/api/v1/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
  -d '{
    "title": "学习Go-Zero",
    "description": "完成官方教程",
    "category": "学习",
    "priority": 2
  }'
```

### 获取任务列表
```bash
curl -X GET "http://localhost:8888/api/v1/tasks?page=1&size=10" \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

---

## 注意事项

1. **认证**: 除认证相关接口外，所有接口都需要在请求头中携带有效的 JWT Token
2. **分页**: 列表接口支持分页，默认每页10条记录
3. **时间格式**: 所有时间字段均使用 ISO 8601 格式 (YYYY-MM-DDTHH:mm:ssZ)
4. **用户隔离**: 所有数据操作都基于当前登录用户，确保数据安全
5. **软删除**: 任务删除采用软删除方式，数据不会物理删除

---

## 更新日志

- **v1.0** (2025-11-25): 初始版本发布，包含完整的用户认证、任务管理、分类管理和统计功能