<template>
  <div class="categories-container">
    <!-- 页面头部 -->
    <div class="page-header">
      <div class="header-left">
        <h1 class="page-title">分类管理</h1>
        <p class="page-description">管理任务分类，更好地组织你的任务</p>
      </div>
      <div class="header-right">
        <el-button type="primary" :icon="Plus" @click="showCreateDialog = true">
          创建分类
        </el-button>
      </div>
    </div>

    <!-- 分类列表 -->
    <div class="categories-grid">
      <el-card v-loading="taskStore.loading">
        <div v-if="taskStore.categories.length === 0" class="empty-state">
          <el-empty description="暂无分类">
            <el-button type="primary" @click="showCreateDialog = true">
              创建第一个分类
            </el-button>
          </el-empty>
        </div>

        <div v-else class="category-list">
          <div
            v-for="category in taskStore.categories"
            :key="category.id"
            class="category-item"
          >
            <div class="category-left">
              <div
                class="category-color"
                :style="{ backgroundColor: category.color }"
              />
              <div class="category-content">
                <h3 class="category-name">{{ category.name }}</h3>
                <p class="category-date">
                  创建于 {{ formatDate(category.created_at) }}
                </p>
              </div>
            </div>
            <div class="category-right">
              <el-dropdown @command="(cmd: string) => handleAction(cmd, category)">
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

    <!-- 创建分类对话框 -->
    <CreateCategoryDialog
      v-model="showCreateDialog"
      @success="handleCreateSuccess"
    />

    <!-- 编辑分类对话框 -->
    <EditCategoryDialog
      v-model="showEditDialog"
      :category="currentCategory"
      @success="handleEditSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessageBox } from 'element-plus'
import { ElMessage } from 'element-plus'
import {
  Plus,
  MoreFilled,
  Edit,
  Delete
} from '@element-plus/icons-vue'
import dayjs from 'dayjs'
import { useTaskStore } from '@/stores/task'
import type { Category } from '@/api/types'
import CreateCategoryDialog from '@/components/categories/CreateCategoryDialog.vue'
import EditCategoryDialog from '@/components/categories/EditCategoryDialog.vue'

const taskStore = useTaskStore()

// 响应式数据
const showCreateDialog = ref(false)
const showEditDialog = ref(false)
const currentCategory = ref<Category | null>(null)

// 格式化日期
const formatDate = (dateStr: string) => {
  return dayjs(dateStr).format('YYYY-MM-DD HH:mm')
}

// 处理操作
const handleAction = (command: string, category: Category) => {
  switch (command) {
    case 'edit':
      currentCategory.value = category
      showEditDialog.value = true
      break
    case 'delete':
      handleDelete(category)
      break
  }
}

// 处理删除
const handleDelete = async (category: Category) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除分类"${category.name}"吗？删除后该分类下的任务将不再显示分类信息。`,
      '提示',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )

    // 这里应该调用删除分类的API
    // await taskStore.deleteCategory(category.id)
    ElMessage.success('删除成功')
    await taskStore.fetchCategories()
  } catch (error) {
    // 用户取消操作
  }
}

// 处理创建成功
const handleCreateSuccess = () => {
  showCreateDialog.value = false
  taskStore.fetchCategories()
}

// 处理编辑成功
const handleEditSuccess = () => {
  showEditDialog.value = false
  currentCategory.value = null
  taskStore.fetchCategories()
}

// 页面加载时获取数据
onMounted(async () => {
  await taskStore.fetchCategories()
})
</script>

<style scoped>
.categories-container {
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

.categories-grid {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.empty-state {
  padding: 60px 0;
}

.category-list {
  padding: 8px 0;
}

.category-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 20px;
  border-bottom: 1px solid #f0f0f0;
  transition: background-color 0.2s;
}

.category-item:hover {
  background-color: #f8f9fa;
}

.category-item:last-child {
  border-bottom: none;
}

.category-left {
  display: flex;
  align-items: center;
  gap: 12px;
  flex: 1;
}

.category-color {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  flex-shrink: 0;
}

.category-content {
  flex: 1;
}

.category-name {
  font-size: 16px;
  font-weight: 500;
  color: #303133;
  margin-bottom: 4px;
  line-height: 1.4;
}

.category-date {
  font-size: 12px;
  color: #909399;
  line-height: 1.4;
}

.category-right {
  display: flex;
  align-items: center;
}
</style>