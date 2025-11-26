// API响应类型
export interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
}

// 用户相关类型
export interface UserInfo {
  id: number
  username: string
  email: string
  created_at: string
  updated_at: string
}

export interface LoginRequest {
  username: string
  password: string
}

export interface RegisterRequest {
  username: string
  email: string
  password: string
}

export interface LoginResponse {
  user: UserInfo
  access_token: string
  refresh_token: string
  expires_at: number
}

export interface ChangePasswordRequest {
  old_password: string
  new_password: string
}

export interface UpdateUserRequest {
  username?: string
  email?: string
}

// 任务相关类型
export enum TaskPriority {
  LOW = 1,
  MEDIUM = 2,
  HIGH = 3
}

export enum TaskStatus {
  TODO = 0,        // 未完成
  COMPLETED = 1,   // 已完成
  DELETED = 2      // 已删除
}

export interface Task {
  id: number
  user_id: number
  title: string
  description: string
  category: string
  priority: number
  status: TaskStatus
  due_date: string
  created_at: string
  updated_at: string
  completed_at: string
}

export interface CreateTaskRequest {
  title: string
  description?: string
  category?: string
  priority?: number
  due_date?: string
}

export interface UpdateTaskRequest {
  title?: string
  description?: string
  category?: string
  priority?: number
  due_date?: string
}

export interface GetTasksRequest {
  page?: number
  size?: number
  category?: string
  status?: TaskStatus
}

export interface GetTasksResponse {
  tasks: Task[]
  total: number
}

// 分类相关类型
export interface Category {
  id: number
  user_id: number
  name: string
  color: string
  created_at: string
  updated_at: string
}

export interface CreateCategoryRequest {
  name: string
  color?: string
}