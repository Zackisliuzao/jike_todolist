import axios from 'axios'
import { ElMessage } from 'element-plus'
import router from '@/router'
import { useUserStore } from '@/stores/user'

// 创建axios实例
const request = axios.create({
  baseURL: '/api/v1',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 请求拦截器
request.interceptors.request.use(
  (config) => {
    const userStore = useUserStore()
    const token = userStore.token

    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }

    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
request.interceptors.response.use(
  (response: any) => {
    // 检查response.data是否存在
    if (!response.data) {
      // 静默处理空数据，不显示错误消息
      return {
        code: 200,
        message: 'success',
        data: null
      }
    }

    // 处理字符串"null"的情况
    if (typeof response.data === 'string' && response.data === 'null') {
      return {
        code: 200,
        message: 'success',
        data: null
      }
    }

    // 后端现在统一返回 {code, message, data} 格式
    const { code, message, data } = response.data

    if (code === 200) {
      return response.data
    } else {
      ElMessage.error(message || '请求失败')
      return Promise.reject(new Error(message || '请求失败'))
    }
  },
  (error: any) => {
    if (error.response?.status === 401) {
      const userStore = useUserStore()
      userStore.logout()
      router.push('/login')
      ElMessage.error('登录已过期，请重新登录')
    } else if (error.response?.status === 403) {
      ElMessage.error('权限不足')
    } else if (error.response?.status === 500) {
      ElMessage.error('服务器错误')
    } else {
      ElMessage.error(error.message || '网络错误')
    }

    return Promise.reject(error)
  }
)

export default request