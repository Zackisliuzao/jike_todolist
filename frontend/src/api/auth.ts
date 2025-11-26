import request from './request'
import {
  type ApiResponse,
  type LoginRequest,
  type RegisterRequest,
  type LoginResponse,
  type UserInfo,
  type ChangePasswordRequest,
  type UpdateUserRequest
} from './types'

export const authApi = {
  // 用户注册
  register: (data: RegisterRequest): Promise<ApiResponse<LoginResponse>> => {
    return request.post('/auth/register', data)
  },

  // 用户登录
  login: (data: LoginRequest): Promise<ApiResponse<LoginResponse>> => {
    return request.post('/auth/login', data)
  },

  // 用户登出
  logout: (): Promise<ApiResponse> => {
    return request.post('/auth/logout')
  },

  // 刷新token
  refreshToken: (refreshToken: string): Promise<ApiResponse<{ token: string }>> => {
    return request.post('/auth/refresh-token', { refresh_token: refreshToken })
  },

  // 获取用户信息
  getUserInfo: (): Promise<ApiResponse<UserInfo>> => {
    return request.get('/user/profile')
  },

  // 修改密码
  changePassword: (data: ChangePasswordRequest): Promise<ApiResponse> => {
    return request.put('/user/password', data)
  },

  // 更新用户信息
  updateUser: (data: UpdateUserRequest): Promise<ApiResponse<UserInfo>> => {
    return request.put('/user/profile', data)
  }
}