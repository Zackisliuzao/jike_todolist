// 简化的请求封装，避免axios类型问题

// 基础请求函数
export const apiRequest = async (url: string, options: RequestInit = {}) => {
  try {
    const token = localStorage.getItem('token')

    const response = await fetch(`/api${url}`, {
      headers: {
        'Content-Type': 'application/json',
        ...(token && { Authorization: `Bearer ${token}` }),
        ...options.headers,
      },
      ...options,
    })

    const data = await response.json()

    if (!response.ok) {
      throw new Error(data.message || '请求失败')
    }

    return data
  } catch (error) {
    console.error('API请求错误:', error)
    throw error
  }
}

// GET请求
export const get = (url: string) => apiRequest(url, { method: 'GET' })

// POST请求
export const post = (url: string, data: any) =>
  apiRequest(url, {
    method: 'POST',
    body: JSON.stringify(data),
  })

// PUT请求
export const put = (url: string, data: any) =>
  apiRequest(url, {
    method: 'PUT',
    body: JSON.stringify(data),
  })

// DELETE请求
export const del = (url: string) => apiRequest(url, { method: 'DELETE' })