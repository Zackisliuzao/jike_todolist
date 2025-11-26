<template>
  <el-dialog
    v-model="visible"
    title="编辑分类"
    width="400px"
    :before-close="handleClose"
  >
    <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="80px"
      @submit.prevent="handleSubmit"
      v-if="category"
    >
      <el-form-item label="分类名称" prop="name">
        <el-input
          v-model="form.name"
          placeholder="请输入分类名称"
          clearable
        />
      </el-form-item>

      <el-form-item label="分类颜色" prop="color">
        <div class="color-picker-wrapper">
          <el-color-picker
            v-model="form.color"
            :predefine="predefineColors"
            show-alpha
          />
          <span v-if="form.color" class="color-preview">
            当前颜色：{{ form.color }}
          </span>
        </div>
      </el-form-item>
    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button type="primary" :loading="loading" @click="handleSubmit">
          保存
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive, computed, watch } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
import type { Category } from '@/api/types'

interface Props {
  modelValue: boolean
  category: Category | null
}

interface Emits {
  (e: 'update:modelValue', value: boolean): void
  (e: 'success'): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

// 表单引用
const formRef = ref<FormInstance>()

// 加载状态
const loading = ref(false)

// 对话框可见性
const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

// 预定义颜色
const predefineColors = [
  '#1890ff',
  '#52c41a',
  '#faad14',
  '#f5222d',
  '#722ed1',
  '#13c2c2',
  '#eb2f96',
  '#fa8c16',
  '#2f54eb',
  '#a0d911'
]

// 表单数据
const form = reactive({
  name: '',
  color: '#1890ff'
})

// 表单验证规则
const rules: FormRules = {
  name: [
    { required: true, message: '请输入分类名称', trigger: 'blur' },
    { min: 1, max: 20, message: '分类名称长度为 1-20 个字符', trigger: 'blur' }
  ],
  color: [
    { required: true, message: '请选择分类颜色', trigger: 'change' }
  ]
}

// 初始化表单数据
const initForm = () => {
  if (props.category) {
    form.name = props.category.name
    form.color = props.category.color
  }
}

// 重置表单
const resetForm = () => {
  form.name = ''
  form.color = '#1890ff'
  formRef.value?.clearValidate()
}

// 处理关闭
const handleClose = () => {
  resetForm()
  visible.value = false
}

// 处理提交
const handleSubmit = async () => {
  if (!formRef.value || !props.category) return

  try {
    await formRef.value.validate()
    loading.value = true

    // 这里应该调用更新分类的API
    // await taskStore.updateCategory(props.category.id, form)
    emit('success')
  } catch (error) {
    console.error('Update category error:', error)
  } finally {
    loading.value = false
  }
}

// 监听分类变化，初始化表单
watch(() => props.category, () => {
  if (props.category) {
    initForm()
  }
}, { immediate: true })
</script>

<style scoped>
.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

.color-picker-wrapper {
  display: flex;
  align-items: center;
  gap: 12px;
}

.color-preview {
  font-size: 14px;
  color: #606266;
}
</style>