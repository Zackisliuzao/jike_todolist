import { defineStore } from 'pinia'
import { ref } from 'vue'
import { taskApi } from '@/api/task'
import type {
  Task,
  CreateTaskRequest,
  UpdateTaskRequest,
  GetTasksRequest,
  GetTasksResponse,
  Category,
  CreateCategoryRequest,
  TaskStatus
} from '@/api/types'
import { ElMessage } from 'element-plus'

export const useTaskStore = defineStore('task', () => {
  // 状态
  const tasks = ref<Task[]>([])
  const categories = ref<Category[]>([])
  const loading = ref(false)
  const total = ref(0)

  // 获取任务列表
  const fetchTasks = async (params?: GetTasksRequest) => {
    try {
      loading.value = true
      const response = await taskApi.getTasks({
        page: 1,
        size: 20,
        ...params
      })
      tasks.value = response.data.tasks
      total.value = response.data.total
      return Promise.resolve()
    } catch (error) {
      return Promise.reject(error)
    } finally {
      loading.value = false
    }
  }

  // 创建任务
  const createTask = async (data: CreateTaskRequest) => {
    try {
      const response = await taskApi.createTask(data)
      if (response && response.data) {
        tasks.value.unshift(response.data)
        ElMessage.success('任务创建成功')
        return Promise.resolve(response.data)
      } else {
        throw new Error('创建任务时API返回数据格式错误')
      }
    } catch (error) {
      return Promise.reject(error)
    }
  }

  // 获取单个任务
  const getTask = async (id: number): Promise<Task> => {
    try {
      const response = await taskApi.getTask(id)
      return response.data
    } catch (error) {
      return Promise.reject(error)
    }
  }

  // 更新任务
  const updateTask = async (id: number, data: UpdateTaskRequest) => {
    try {
      const response = await taskApi.updateTask(id, data)
      if (response && response.data) {
        const index = tasks.value.findIndex(task => task && task.id === id)
        if (index !== -1) {
          tasks.value[index] = response.data
        }
        ElMessage.success('任务更新成功')
        return Promise.resolve(response.data)
      } else {
        throw new Error('更新任务时API返回数据格式错误')
      }
    } catch (error) {
      return Promise.reject(error)
    }
  }

  // 删除任务
  const deleteTask = async (id: number) => {
    try {
      await taskApi.deleteTask(id)
      const index = tasks.value.findIndex(task => task.id === id)
      if (index !== -1) {
        tasks.value.splice(index, 1)
      }
      total.value = Math.max(0, total.value - 1)
      ElMessage.success('任务删除成功')
      return Promise.resolve()
    } catch (error) {
      return Promise.reject(error)
    }
  }

  // 更新任务状态
  const updateTaskStatus = async (id: number, status: number) => {
    try {
      const response = await taskApi.updateTaskStatus(id, status)
      // 验证响应数据
      if (response && response.data) {
        const index = tasks.value.findIndex(task => task && task.id === id)
        if (index !== -1) {
          tasks.value[index] = response.data
        }
        ElMessage.success('任务状态更新成功')
        return Promise.resolve(response.data)
      } else {
        throw new Error('API返回数据格式错误')
      }
    } catch (error) {
      return Promise.reject(error)
    }
  }

  // 获取分类列表
  const fetchCategories = async () => {
    try {
      const response = await taskApi.getCategories()
      categories.value = (response.data || []).filter(Boolean)
      return Promise.resolve()
    } catch (error) {
      categories.value = []  // 失败时设置为空数组
      return Promise.reject(error)
    }
  }

  // 创建分类
  const createCategory = async (data: CreateCategoryRequest) => {
    try {
      const response = await taskApi.createCategory(data)
      categories.value.push(response.data)
      ElMessage.success('分类创建成功')
      return Promise.resolve(response.data)
    } catch (error) {
      return Promise.reject(error)
    }
  }

  // 更新分类
  const updateCategory = async (id: number, data: CreateCategoryRequest) => {
    try {
      const response = await taskApi.updateCategory(id, data)
      const index = categories.value.findIndex(cat => cat.id === id)
      if (index !== -1) {
        categories.value[index] = response.data
      }
      ElMessage.success('分类更新成功')
      return Promise.resolve(response.data)
    } catch (error) {
      return Promise.reject(error)
    }
  }

  // 删除分类
  const deleteCategory = async (id: number) => {
    try {
      await taskApi.deleteCategory(id)
      const index = categories.value.findIndex(cat => cat.id === id)
      if (index !== -1) {
        categories.value.splice(index, 1)
      }
      ElMessage.success('分类删除成功')
      return Promise.resolve()
    } catch (error) {
      return Promise.reject(error)
    }
  }

  // 获取状态文本
  const getStatusText = (status: number): string => {
    switch (status) {
      case 0: // TODO: 未完成
        return '未完成'
      case 1: // COMPLETED: 已完成
        return '已完成'
      case 2: // DELETED: 已删除
        return '已删除'
      default:
        return '未知'
    }
  }

  // 获取优先级文本
  const getPriorityText = (priority: number): string => {
    switch (priority) {
      case 1: // LOW
        return '低'
      case 2: // MEDIUM
        return '中'
      case 3: // HIGH
        return '高'
      default:
        return '未知'
    }
  }

  // 获取优先级颜色
  const getPriorityColor = (priority: number): string => {
    switch (priority) {
      case 1: // LOW
        return '#909399'
      case 2: // MEDIUM
        return '#e6a23c'
      case 3: // HIGH
        return '#f56c6c'
      default:
        return '#909399'
    }
  }

  return {
    tasks,
    categories,
    loading,
    total,
    fetchTasks,
    createTask,
    getTask,
    updateTask,
    deleteTask,
    updateTaskStatus,
    fetchCategories,
    createCategory,
    updateCategory,
    deleteCategory,
    getStatusText,
    getPriorityText,
    getPriorityColor
  }
})