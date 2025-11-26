<template>
  <div class="profile-container">
    <!-- 页面头部 -->
    <div class="page-header">
      <div class="header-left">
        <h1 class="page-title">个人资料</h1>
        <p class="page-description">管理你的个人信息和账户设置</p>
      </div>
    </div>

    <el-row :gutter="20">
      <!-- 基本信息卡片 -->
      <el-col :span="16">
        <el-card class="profile-card">
          <template #header>
            <div class="card-header">
              <span>基本信息</span>
              <!-- <el-button
                type="primary"
                size="small"
                :icon="Edit"
                @click="showEditDialog = true"
              >
                编辑资料
              </el-button> -->
            </div>
          </template>

          <div class="profile-info">
            <div class="avatar-section">
              <el-avatar :size="80" :icon="UserFilled" />
            </div>
            <div class="info-section">
              <div class="info-item">
                <label>用户名</label>
                <span>{{ userStore.userInfo?.username }}</span>
              </div>
              <div class="info-item">
                <label>邮箱</label>
                <span>{{ userStore.userInfo?.email }}</span>
              </div>
              <div class="info-item">
                <label>注册时间</label>
                <span>{{ formatDate(userStore.userInfo?.created_at) }}</span>
              </div>
              <div class="info-item">
                <label>最后更新</label>
                <span>{{ formatDate(userStore.userInfo?.updated_at) }}</span>
              </div>
            </div>
          </div>
        </el-card>

        <!-- 修改密码卡片 -->
        <el-card class="password-card">
          <template #header>
            <span>修改密码</span>
          </template>

          <el-form
            ref="passwordFormRef"
            :model="passwordForm"
            :rules="passwordRules"
            label-width="100px"
            @submit.prevent="handleChangePassword"
          >
            <el-form-item label="当前密码" prop="old_password">
              <el-input
                v-model="passwordForm.old_password"
                type="password"
                placeholder="请输入当前密码"
                show-password
                clearable
              />
            </el-form-item>

            <el-form-item label="新密码" prop="new_password">
              <el-input
                v-model="passwordForm.new_password"
                type="password"
                placeholder="请输入新密码"
                show-password
                clearable
              />
            </el-form-item>

            <el-form-item label="确认密码" prop="confirm_password">
              <el-input
                v-model="passwordForm.confirm_password"
                type="password"
                placeholder="请再次输入新密码"
                show-password
                clearable
                @keyup.enter="handleChangePassword"
              />
            </el-form-item>

            <el-form-item>
              <el-button
                type="primary"
                :loading="passwordLoading"
                @click="handleChangePassword"
              >
                修改密码
              </el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>

      <!-- 账户统计 -->
      <el-col :span="8">
        <el-card class="stats-card">
          <template #header>
            <span>账户统计</span>
          </template>

          <div class="stats-content">
            <div class="stat-item">
              <div class="stat-icon">
                <el-icon><List /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ (taskStore.tasks || []).length }}</div>
                <div class="stat-label">总任务数</div>
              </div>
            </div>

            <div class="stat-item">
              <div class="stat-icon completed">
                <el-icon><CircleCheck /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ completedTasksCount }}</div>
                <div class="stat-label">已完成</div>
              </div>
            </div>

            <div class="stat-item">
              <div class="stat-icon categories">
                <el-icon><Collection /></el-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">{{ taskStore.categories.length }}</div>
                <div class="stat-label">分类数</div>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 编辑资料对话框 -->
    <EditProfileDialog
      v-model="showEditDialog"
      @success="handleEditSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
import {
  Edit,
  UserFilled,
  List,
  CircleCheck,
  Collection
} from '@element-plus/icons-vue'
import dayjs from 'dayjs'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { useTaskStore } from '@/stores/task'
import { authApi } from '@/api/auth'
import type { ChangePasswordRequest } from '@/api/types'
import EditProfileDialog from '@/components/user/EditProfileDialog.vue'
import { ElMessage } from 'element-plus'

const router = useRouter()
const userStore = useUserStore()
const taskStore = useTaskStore()

// 表单引用
const passwordFormRef = ref<FormInstance>()

// 响应式数据
const showEditDialog = ref(false)
const passwordLoading = ref(false)
// confirmPassword 已移动到 passwordForm 对象中

// 修改密码表单
const passwordForm = reactive<ChangePasswordRequest & { confirm_password: string }>({
  old_password: '',
  new_password: '',
  confirm_password: ''
})

// 计算属性
const completedTasksCount = computed(() =>
  (taskStore.tasks || []).filter(task => task.status === 1).length
)

// 确认密码验证器
const validateConfirmPassword = (rule: any, value: string, callback: any) => {
  if (value === '') {
    callback(new Error('请再次输入新密码'))
  } else if (value !== passwordForm.new_password) {
    callback(new Error('两次输入的密码不一致'))
  } else {
    callback()
  }
}

// 表单验证规则
const passwordRules: FormRules = {
  old_password: [
    { required: true, message: '请输入当前密码', trigger: 'blur' }
  ],
  new_password: [
    { required: true, message: '请输入新密码', trigger: 'blur' },
    { min: 6, max: 20, message: '密码长度为 6-20 个字符', trigger: 'blur' }
  ],
  confirm_password: [
    { required: true, message: '请确认新密码', trigger: 'blur' },
    { validator: validateConfirmPassword, trigger: 'blur' }
  ]
}

// 格式化日期
const formatDate = (dateStr?: string) => {
  if (!dateStr) return '-'
  return dayjs(dateStr).format('YYYY-MM-DD HH:mm')
}

// 处理修改密码
const handleChangePassword = async () => {
  if (!passwordFormRef.value) return

  try {
    await passwordFormRef.value.validate()
    passwordLoading.value = true

    await authApi.changePassword({
      old_password: passwordForm.old_password,
      new_password: passwordForm.new_password
    })

    // 重置表单
    passwordForm.old_password = ''
    passwordForm.new_password = ''
    passwordForm.confirm_password = ''
    passwordFormRef.value.clearValidate()

    ElMessage.success('密码修改成功，请重新登录')

    // 修改密码后自动登出
    setTimeout(async () => {
      await userStore.logout(false) // 不显示登出消息，避免重复
      // 跳转到登录页面
      router.push('/login')
    }, 1500) // 延迟1.5秒，让用户看到成功消息
  } catch (error) {
    console.error('Change password error:', error)
  } finally {
    passwordLoading.value = false
  }
}

// 处理编辑成功
const handleEditSuccess = () => {
  showEditDialog.value = false
  userStore.fetchUserInfo()
}

// 页面加载时获取数据
onMounted(async () => {
  await Promise.all([
    userStore.fetchUserInfo(),
    taskStore.fetchTasks({ page: 1, size: 100 }),
    taskStore.fetchCategories()
  ])
})
</script>

<style scoped>
.profile-container {
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

.profile-card {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.profile-info {
  display: flex;
  gap: 24px;
}

.avatar-section {
  flex-shrink: 0;
}

.info-section {
  flex: 1;
}

.info-item {
  display: flex;
  align-items: center;
  padding: 12px 0;
  border-bottom: 1px solid #f0f0f0;
}

.info-item:last-child {
  border-bottom: none;
}

.info-item label {
  width: 100px;
  font-size: 14px;
  color: #909399;
  font-weight: 500;
}

.info-item span {
  font-size: 14px;
  color: #303133;
}

.password-card {
  margin-bottom: 20px;
}

.stats-card {
  height: fit-content;
}

.stats-content {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 16px;
  background-color: #f8f9fa;
  border-radius: 8px;
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: 8px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 20px;
  flex-shrink: 0;
}

.stat-icon.completed {
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
}

.stat-icon.categories {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
}

.stat-info {
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
</style>