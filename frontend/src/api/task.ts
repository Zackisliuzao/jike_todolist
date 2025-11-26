import request from './request'
import {
  type ApiResponse,
  type Task,
  type CreateTaskRequest,
  type UpdateTaskRequest,
  type GetTasksRequest,
  type GetTasksResponse,
  type Category,
  type CreateCategoryRequest,
  TaskStatus
} from './types'

export const taskApi = {
  // 创建任务
  createTask: (data: CreateTaskRequest): Promise<ApiResponse<Task>> => {
    return request.post('/tasks', data)
  },

  // 获取任务列表
  getTasks: (params: GetTasksRequest): Promise<ApiResponse<GetTasksResponse>> => {
    return request.get('/tasks', { params })
  },

  // 获取单个任务
  getTask: (id: number): Promise<ApiResponse<Task>> => {
    return request.get(`/tasks/${id}`)
  },

  // 更新任务
  updateTask: (id: number, data: UpdateTaskRequest): Promise<ApiResponse<Task>> => {
    return request.put(`/tasks/${id}`, data)
  },

  // 删除任务
  deleteTask: (id: number): Promise<ApiResponse<{ success: boolean }>> => {
    return request.delete(`/tasks/${id}`)
  },

  // 更新任务状态
  updateTaskStatus: (id: number, status: TaskStatus): Promise<ApiResponse<Task>> => {
    return request.put(`/tasks/${id}/status`, { status })
  },

  // 获取分类列表
  getCategories: (): Promise<ApiResponse<Category[]>> => {
    return request.get('/categories')
  },

  // 创建分类
  createCategory: (data: CreateCategoryRequest): Promise<ApiResponse<Category>> => {
    return request.post('/categories', data)
  }
}