<template>
  <div class="tasks-container">
    <!-- 页面头部 -->
    <div class="page-header">
      <div class="header-left">
        <h1 class="page-title">任务列表</h1>
        <p class="page-description">管理你的所有任务</p>
      </div>
      <div class="header-right">
        <el-button type="primary" :icon="Plus" @click="showCreateDialog = true">
          创建任务
        </el-button>
      </div>
    </div>

    <!-- 筛选器 -->
    <div class="filters-card">
      <el-row :gutter="16">
        <el-col :span="6">
          <el-select
            v-model="filters.status"
            placeholder="任务状态"
            clearable
            @change="handleFilterChange"
          >
            <el-option label="全部" value="" />
            <el-option label="未完成" :value="0" />
            <el-option label="已完成" :value="1" />
            <!-- <el-option label="已删除" :value="2" /> -->
          </el-select>
        </el-col>
        <el-col :span="6">
          <el-select
            v-model="filters.category"
            placeholder="分类"
            clearable
            @change="handleFilterChange"
          >
            <el-option label="全部" value="" />
            <el-option
              v-for="category in taskStore.categories"
              :key="category.id"
              :label="category.name"
              :value="category.name"
            />
          </el-select>
        </el-col>
        <el-col :span="12">
          <el-input
            v-model="searchKeyword"
            placeholder="搜索任务标题..."
            clearable
            @input="handleSearch"
          >
            <template #prefix>
              <el-icon><Search /></el-icon>
            </template>
          </el-input>
        </el-col>
      </el-row>
    </div>

    <!-- 任务统计 -->
    <div class="stats-row">
      <div class="stat-card">
        <div class="stat-icon todo">
          <el-icon><Clock /></el-icon>
        </div>
        <div class="stat-content">
          <div class="stat-number">{{ todoCount }}</div>
          <div class="stat-label">待办</div>
        </div>
      </div>
      <!-- <div class="stat-card">
        <div class="stat-icon progress">
          <el-icon><Loading /></el-icon>
        </div>
        <div class="stat-content">
          <div class="stat-number">{{ inProgressCount }}</div>
          <div class="stat-label">进行中</div>
        </div>
      </div> -->
      <div class="stat-card">
        <div class="stat-icon completed">
          <el-icon><CircleCheck /></el-icon>
        </div>
        <div class="stat-content">
          <div class="stat-number">{{ completedCount }}</div>
          <div class="stat-label">已完成</div>
        </div>
      </div>
    </div>

    <!-- 任务列表 -->
    <div class="tasks-list">
      <el-card v-loading="taskStore.loading">
        <div v-if="filteredTasks.length === 0" class="empty-state">
          <el-empty description="暂无任务">
            <el-button type="primary" @click="showCreateDialog = true">
              创建第一个任务
            </el-button>
          </el-empty>
        </div>

        <div v-else class="task-items">
          <div
            v-for="task in filteredTasks"
            :key="task.id"
            class="task-item"
            @click="viewTask(task.id)"
          >
            <div class="task-left">
              <el-checkbox
                :model-value="task.status === 1"
                @change="handleStatusChange(task.id, $event)"
                @click.stop
              />
              <div class="task-content">
                <h3 class="task-title" :class="{ completed: task.status === 1 }">
                  {{ task.title }}
                </h3>
                <p v-if="task.description" class="task-description">
                  {{ task.description }}
                </p>
                <div class="task-meta">
                  <el-tag
                    v-if="task.category"
                    size="small"
                    type="info"
                    class="mr-8"
                  >
                    {{ task.category }}
                  </el-tag>
                  <el-tag
                    :color="taskStore.getPriorityColor(task.priority)"
                    size="small"
                    effect="dark"
                    class="mr-8"
                  >
                    {{ taskStore.getPriorityText(task.priority) }}
                  </el-tag>
                  <span class="task-date">
                    {{ formatDate(task.created_at) }}
                  </span>
                </div>
              </div>
            </div>
            <div class="task-right">
              <el-dropdown @command="(cmd) => handleTaskAction(cmd, task)" @click.stop>
                <el-button text :icon="MoreFilled" />
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item command="edit">
                      <el-icon><Edit /></el-icon>
                      编辑
                    </el-dropdown-item>
                    <el-dropdown-item command="delete" divided>
                      <el-icon><Delete /></el-icon>
                      删除
                    </el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </div>
          </div>
        </div>
      </el-card>
    </div>

    <!-- 创建任务对话框 -->
    <CreateTaskDialog
      v-model="showCreateDialog"
      @success="handleCreateSuccess"
    />

    <!-- 编辑任务对话框 -->
    <EditTaskDialog
      v-model="showEditDialog"
      :task="currentTask"
      @success="handleEditSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessageBox } from 'element-plus'
import {
  Plus,
  Search,
  Clock,
  Loading,
  CircleCheck,
  MoreFilled,
  Edit,
  Delete
} from '@element-plus/icons-vue'
import dayjs from 'dayjs'
import { useTaskStore } from '@/stores/task'
import type { Task } from '@/api/types'
import CreateTaskDialog from '@/components/tasks/CreateTaskDialog.vue'
import EditTaskDialog from '@/components/tasks/EditTaskDialog.vue'

const router = useRouter()
const taskStore = useTaskStore()

// 响应式数据
const showCreateDialog = ref(false)
const showEditDialog = ref(false)
const currentTask = ref<Task | null>(null)
const searchKeyword = ref('')

// 筛选器
const filters = reactive({
  status: '' as string,
  category: ''
})

// 过滤后的任务
const filteredTasks = computed(() => {
  let tasks = (taskStore.tasks || []).filter(task => task) // 首先过滤掉undefined任务

  console.log('Original tasks count:', tasks.length)
  console.log('Current filters:', filters.status, filters.category)

  // 状态筛选 - 处理空字符串、null、undefined
  if (filters.status !== '' && filters.status != null && filters.status !== undefined) {
    tasks = tasks.filter(task => task.status === filters.status)
    console.log('After status filter:', tasks.length)
  }

  // 分类筛选
  if (filters.category) {
    tasks = tasks.filter(task => task.category === filters.category)
  }

  // 搜索过滤
  if (searchKeyword.value) {
    tasks = tasks.filter(task =>
      task.title.toLowerCase().includes(searchKeyword.value.toLowerCase())
    )
  }

  console.log('Final filtered tasks count:', tasks.length)
  return tasks
})

// 任务统计
const todoCount = computed(() =>
  (taskStore.tasks || []).filter(task => task && task.status === 0).length
)

const inProgressCount = computed(() =>
  0 // 没有进行中状态，数据库中没有这个状态
)

const completedCount = computed(() =>
  (taskStore.tasks || []).filter(task => task && task.status === 1).length
)

// 格式化日期
const formatDate = (dateStr: string) => {
  return dayjs(dateStr).format('MM-DD HH:mm')
}

// 处理筛选变化
const handleFilterChange = () => {
  // 筛选会自动通过计算属性生效
  console.log('Filter changed:', filters.status, filters.category)
  console.log('Task store count:', taskStore.tasks?.length)
}

// 处理搜索
const handleSearch = () => {
  // 搜索会自动通过计算属性生效
}

// 获取任务列表
const fetchTasks = async () => {
  try {
    await taskStore.fetchTasks({ page: 1, size: 100 })
  } catch (error) {
    console.error('Fetch tasks error:', error)
  }
}

// 处理任务状态变化
const handleStatusChange = async (taskId: number, checked: boolean) => {
  const newStatus = checked ? 1 : 0
  try {
    await taskStore.updateTaskStatus(taskId, newStatus)
  } catch (error) {
    console.error('Update task status error:', error)
  }
}

// 查看任务详情
const viewTask = (taskId: number) => {
  router.push(`/dashboard/tasks/${taskId}`)
}

// 处理任务操作
const handleTaskAction = async (command: string, task: Task) => {
  switch (command) {
    case 'edit':
      currentTask.value = task
      showEditDialog.value = true
      break
    case 'delete':
      try {
        await ElMessageBox.confirm(
          `确定要删除任务"${task.title}"吗？`,
          '提示',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }
        )
        await taskStore.deleteTask(task.id)
      } catch (error) {
        // 用户取消操作
      }
      break
  }
}

// 处理创建成功
const handleCreateSuccess = () => {
  showCreateDialog.value = false
  fetchTasks()
}

// 处理编辑成功
const handleEditSuccess = () => {
  showEditDialog.value = false
  currentTask.value = null
  fetchTasks()
}

// 页面加载时获取数据
onMounted(async () => {
  await Promise.all([
    fetchTasks(),
    taskStore.fetchCategories()
  ])
})
</script>

<style scoped>
.tasks-container {
  padding: 20px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.header-left {
  flex: 1;
}

.page-title {
  font-size: 24px;
  font-weight: 600;
  color: #303133;
  margin-bottom: 4px;
}

.page-description {
  font-size: 14px;
  color: #909399;
}

.filters-card {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  padding: 20px;
  margin-bottom: 20px;
}

.stats-row {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 20px;
  margin-bottom: 20px;
}

.stat-card {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  padding: 20px;
  display: flex;
  align-items: center;
  gap: 16px;
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  color: white;
}

.stat-icon.todo {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.stat-icon.progress {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
}

.stat-icon.completed {
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
}

.stat-content {
  flex: 1;
}

.stat-number {
  font-size: 24px;
  font-weight: 600;
  color: #303133;
  line-height: 1;
}

.stat-label {
  font-size: 14px;
  color: #909399;
  margin-top: 4px;
}

.tasks-list {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.empty-state {
  padding: 60px 0;
}

.task-items {
  padding: 8px 0;
}

.task-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 20px;
  border-bottom: 1px solid #f0f0f0;
  cursor: pointer;
  transition: background-color 0.2s;
}

.task-item:hover {
  background-color: #f8f9fa;
}

.task-item:last-child {
  border-bottom: none;
}

.task-left {
  display: flex;
  align-items: flex-start;
  gap: 12px;
  flex: 1;
}

.task-content {
  flex: 1;
  min-width: 0;
}

.task-title {
  font-size: 16px;
  font-weight: 500;
  color: #303133;
  margin-bottom: 4px;
  line-height: 1.4;
}

.task-title.completed {
  text-decoration: line-through;
  color: #909399;
}

.task-description {
  font-size: 14px;
  color: #606266;
  margin-bottom: 8px;
  line-height: 1.4;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.task-meta {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
}

.task-date {
  font-size: 12px;
  color: #909399;
}

.task-right {
  display: flex;
  align-items: center;
}

.mr-8 {
  margin-right: 8px;
}
</style>