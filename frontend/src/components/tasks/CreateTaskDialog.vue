<template>
  <el-dialog
    v-model="visible"
    title="创建任务"
    width="500px"
    :before-close="handleClose"
  >
    <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="80px"
      @submit.prevent="handleSubmit"
    >
      <el-form-item label="任务标题" prop="title">
        <el-input
          v-model="form.title"
          placeholder="请输入任务标题"
          clearable
        />
      </el-form-item>

      <el-form-item label="任务描述" prop="description">
        <el-input
          v-model="form.description"
          type="textarea"
          :rows="3"
          placeholder="请输入任务描述（可选）"
        />
      </el-form-item>

      <el-form-item label="分类" prop="category">
        <el-select
          v-model="form.category"
          placeholder="选择分类"
          clearable
          style="width: 100%"
        >
          <el-option
            v-for="category in taskStore.categories"
            :key="category.id"
            :label="category.name"
            :value="category.name"
          />
        </el-select>
      </el-form-item>

      <el-form-item label="优先级" prop="priority">
        <el-select
          v-model="form.priority"
          placeholder="选择优先级"
          style="width: 100%"
        >
          <el-option label="低" :value="TaskPriority.LOW" />
          <el-option label="中" :value="TaskPriority.MEDIUM" />
          <el-option label="高" :value="TaskPriority.HIGH" />
        </el-select>
      </el-form-item>

      <el-form-item label="截止时间" prop="due_date">
        <el-date-picker
          v-model="form.due_date"
          type="datetime"
          placeholder="选择截止时间（可选）"
          format="YYYY-MM-DD HH:mm"
          value-format="YYYY-MM-DD HH:mm:ss"
          style="width: 100%"
        />
      </el-form-item>
    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button type="primary" :loading="loading" @click="handleSubmit">
          创建
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive, computed, watch } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
import { useTaskStore } from '@/stores/task'
import { TaskPriority, type CreateTaskRequest } from '@/api/types'

interface Props {
  modelValue: boolean
}

interface Emits {
  (e: 'update:modelValue', value: boolean): void
  (e: 'success'): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const taskStore = useTaskStore()

// 表单引用
const formRef = ref<FormInstance>()

// 加载状态
const loading = ref(false)

// 对话框可见性
const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

// 表单数据
const form = reactive<CreateTaskRequest>({
  title: '',
  description: '',
  category: '',
  priority: TaskPriority.MEDIUM,
  due_date: ''
})

// 表单验证规则
const rules: FormRules = {
  title: [
    { required: true, message: '请输入任务标题', trigger: 'blur' },
    { min: 1, max: 100, message: '标题长度为 1-100 个字符', trigger: 'blur' }
  ],
  description: [
    { max: 500, message: '描述长度不能超过 500 个字符', trigger: 'blur' }
  ]
}

// 重置表单
const resetForm = () => {
  form.title = ''
  form.description = ''
  form.category = ''
  form.priority = TaskPriority.MEDIUM
  form.due_date = ''
  formRef.value?.clearValidate()
}

// 处理关闭
const handleClose = () => {
  resetForm()
  visible.value = false
}

// 处理提交
const handleSubmit = async () => {
  if (!formRef.value) return

  try {
    await formRef.value.validate()
    loading.value = true

    await taskStore.createTask(form)
    emit('success')
  } catch (error) {
    console.error('Create task error:', error)
  } finally {
    loading.value = false
  }
}

// 监听对话框打开，获取分类数据
watch(visible, (newVal) => {
  if (newVal) {
    taskStore.fetchCategories()
  }
})
</script>

<style scoped>
.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
</style>