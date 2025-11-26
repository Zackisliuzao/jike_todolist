import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authApi } from '@/api/auth'
import type { LoginRequest, RegisterRequest, UserInfo } from '@/api/types'
import { ElMessage } from 'element-plus'

export const useUserStore = defineStore('user', () => {
  // 状态
  const token = ref<string>(localStorage.getItem('token') || '')
  const refreshToken = ref<string>(localStorage.getItem('refreshToken') || '')
  const userInfo = ref<UserInfo | null>(null)

  // 计算属性
  const isAuthenticated = computed(() => !!token.value)

  // 登录
  const login = async (credentials: LoginRequest) => {
    try {
      const response = await authApi.login(credentials)
      const { user, access_token, refresh_token } = response.data

      token.value = access_token
      refreshToken.value = refresh_token
      userInfo.value = user

      // 持久化存储
      localStorage.setItem('token', access_token)
      localStorage.setItem('refreshToken', refresh_token)
      localStorage.setItem('userInfo', JSON.stringify(user))

      // 登录成功消息现在由axios拦截器处理
      return Promise.resolve()
    } catch (error) {
      return Promise.reject(error)
    }
  }

  // 注册
  const register = async (userData: RegisterRequest) => {
    try {
      const response = await authApi.register(userData)
      const { user, access_token, refresh_token } = response.data

      token.value = access_token
      refreshToken.value = refresh_token
      userInfo.value = user

      // 持久化存储
      localStorage.setItem('token', access_token)
      localStorage.setItem('refreshToken', refresh_token)
      localStorage.setItem('userInfo', JSON.stringify(user))

      // 注册成功消息现在由axios拦截器处理
      return Promise.resolve()
    } catch (error) {
      return Promise.reject(error)
    }
  }

  // 登出
  const logout = async () => {
    try {
      await authApi.logout()
    } catch (error) {
      console.error('Logout error:', error)
    } finally {
      // 清除状态
      token.value = ''
      refreshToken.value = ''
      userInfo.value = null

      // 清除本地存储
      localStorage.removeItem('token')
      localStorage.removeItem('refreshToken')
      localStorage.removeItem('userInfo')

      ElMessage.success('已退出登录')
    }
  }

  // 获取用户信息
  const fetchUserInfo = async () => {
    try {
      const response = await authApi.getUserInfo()
      userInfo.value = response.data
      localStorage.setItem('userInfo', JSON.stringify(response.data))
      return Promise.resolve()
    } catch (error) {
      return Promise.reject(error)
    }
  }

  // 刷新token
  const refreshAccessToken = async () => {
    try {
      const response = await authApi.refreshToken(refreshToken.value)
      token.value = response.data.token
      localStorage.setItem('token', response.data.token)
      return Promise.resolve()
    } catch (error) {
      // 刷新失败，清除登录状态
      logout()
      return Promise.reject(error)
    }
  }

  // 初始化用户信息
  const initUserInfo = () => {
    const storedUserInfo = localStorage.getItem('userInfo')
    if (storedUserInfo) {
      try {
        userInfo.value = JSON.parse(storedUserInfo)
      } catch (error) {
        console.error('Failed to parse user info:', error)
        localStorage.removeItem('userInfo')
      }
    }
  }

  // 更新用户信息
  const updateUserInfo = (newUserInfo: Partial<UserInfo>) => {
    if (userInfo.value) {
      userInfo.value = { ...userInfo.value, ...newUserInfo }
      localStorage.setItem('userInfo', JSON.stringify(userInfo.value))
    }
  }

  return {
    token,
    refreshToken,
    userInfo,
    isAuthenticated,
    login,
    register,
    logout,
    fetchUserInfo,
    refreshAccessToken,
    initUserInfo,
    updateUserInfo
  }
})