<template>
  <div class="task-detail-container">
    <div class="page-header">
      <el-button :icon="ArrowLeft" @click="goBack">返回</el-button>
      <h1 class="page-title">任务详情</h1>
    </div>

    <div v-loading="loading" class="task-content">
      <el-card v-if="task">
        <div class="task-header">
          <div class="task-title-section">
            <h2 class="task-title">{{ task.title }}</h2>
            <div class="task-meta">
              <el-tag
                :color="taskStore.getPriorityColor(task.priority)"
                effect="dark"
                size="small"
              >
                {{ taskStore.getPriorityText(task.priority) }}
              </el-tag>
              <el-tag
                :type="getStatusType(task.status)"
                size="small"
              >
                {{ taskStore.getStatusText(task.status) }}
              </el-tag>
              <el-tag v-if="task.category" type="info" size="small">
                {{ task.category }}
              </el-tag>
            </div>
          </div>

          <div class="task-actions">
            <el-button
              :type="task.status === 2 ? 'warning' : 'success'"
              @click="toggleTaskStatus"
            >
              {{ task.status === 2 ? '标记为未完成' : '标记为已完成' }}
            </el-button>
            <el-button @click="showEditDialog = true">
              编辑
            </el-button>
            <el-button
              type="danger"
              @click="handleDelete"
            >
              删除
            </el-button>
          </div>
        </div>

        <el-divider />

        <div class="task-body">
          <div class="task-section">
            <h3>任务描述</h3>
            <p class="task-description">
              {{ task.description || '暂无描述' }}
            </p>
          </div>

          <div class="task-section">
            <h3>时间信息</h3>
            <div class="time-info">
              <div class="time-item">
                <label>创建时间：</label>
                <span>{{ formatDateTime(task.created_at) }}</span>
              </div>
              <div class="time-item">
                <label>更新时间：</label>
                <span>{{ formatDateTime(task.updated_at) }}</span>
              </div>
              <div v-if="task.due_date" class="time-item">
                <label>截止时间：</label>
                <span>{{ formatDateTime(task.due_date) }}</span>
              </div>
              <div v-if="task.completed_at" class="time-item">
                <label>完成时间：</label>
                <span>{{ formatDateTime(task.completed_at) }}</span>
              </div>
            </div>
          </div>
        </div>
      </el-card>

      <el-empty v-else description="任务不存在" />
    </div>

    <!-- 编辑任务对话框 -->
    <EditTaskDialog
      v-model="showEditDialog"
      :task="task"
      @success="handleEditSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessageBox } from 'element-plus'
import { ArrowLeft } from '@element-plus/icons-vue'
import dayjs from 'dayjs'
import { useTaskStore } from '@/stores/task'
import { TaskStatus, type Task } from '@/api/types'
import EditTaskDialog from '@/components/tasks/EditTaskDialog.vue'

const route = useRoute()
const router = useRouter()
const taskStore = useTaskStore()

// 响应式数据
const loading = ref(false)
const task = ref<Task | null>(null)
const showEditDialog = ref(false)

// 获取状态类型
const getStatusType = (status: TaskStatus) => {
  switch (status) {
    case TaskStatus.TODO:
      return 'info'
    case TaskStatus.IN_PROGRESS:
      return 'warning'
    case TaskStatus.COMPLETED:
      return 'success'
    default:
      return 'info'
  }
}

// 格式化日期时间
const formatDateTime = (dateStr: string) => {
  return dayjs(dateStr).format('YYYY-MM-DD HH:mm:ss')
}

// 返回上一页
const goBack = () => {
  router.go(-1)
}

// 获取任务详情
const fetchTask = async () => {
  const taskId = Number(route.params.id)
  if (isNaN(taskId)) {
    router.push('/dashboard/tasks')
    return
  }

  try {
    loading.value = true
    task.value = await taskStore.getTask(taskId)
  } catch (error) {
    console.error('Fetch task error:', error)
    router.push('/dashboard/tasks')
  } finally {
    loading.value = false
  }
}

// 切换任务状态
const toggleTaskStatus = async () => {
  if (!task.value) return

  const newStatus = task.value.status === TaskStatus.COMPLETED
    ? TaskStatus.TODO
    : TaskStatus.COMPLETED

  try {
    await taskStore.updateTaskStatus(task.value.id, newStatus)
    task.value.status = newStatus
  } catch (error) {
    console.error('Update task status error:', error)
  }
}

// 处理删除
const handleDelete = async () => {
  if (!task.value) return

  try {
    await ElMessageBox.confirm(
      `确定要删除任务"${task.value.title}"吗？`,
      '提示',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )

    await taskStore.deleteTask(task.value.id)
    router.push('/dashboard/tasks')
  } catch (error) {
    // 用户取消操作
  }
}

// 处理编辑成功
const handleEditSuccess = async () => {
  showEditDialog.value = false
  await fetchTask()
}

// 页面加载时获取任务详情
onMounted(() => {
  fetchTask()
})
</script>

<style scoped>
.task-detail-container {
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;
}

.page-header {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 24px;
}

.page-title {
  font-size: 24px;
  font-weight: 600;
  color: #303133;
  margin: 0;
}

.task-content {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.task-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 20px;
}

.task-title-section {
  flex: 1;
}

.task-title {
  font-size: 24px;
  font-weight: 600;
  color: #303133;
  margin-bottom: 12px;
  line-height: 1.4;
}

.task-meta {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.task-actions {
  display: flex;
  gap: 8px;
  flex-shrink: 0;
}

.task-body {
  padding-top: 20px;
}

.task-section {
  margin-bottom: 32px;
}

.task-section:last-child {
  margin-bottom: 0;
}

.task-section h3 {
  font-size: 16px;
  font-weight: 600;
  color: #303133;
  margin-bottom: 16px;
}

.task-description {
  font-size: 14px;
  color: #606266;
  line-height: 1.6;
  white-space: pre-wrap;
  word-wrap: break-word;
}

.time-info {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 16px;
}

.time-item {
  display: flex;
  align-items: center;
  font-size: 14px;
}

.time-item label {
  color: #909399;
  font-weight: 500;
  margin-right: 8px;
  min-width: 80px;
}

.time-item span {
  color: #303133;
}

@media (max-width: 768px) {
  .task-header {
    flex-direction: column;
    gap: 16px;
    align-items: stretch;
  }

  .task-actions {
    justify-content: flex-start;
  }

  .time-info {
    grid-template-columns: 1fr;
  }
}
</style>