<template>
  <div class="user-management">
    <el-tabs v-model="activeTab" type="border-card">
      <!-- 用户列表 -->
      <el-tab-pane label="用户列表" name="list">
        <el-card shadow="never">
          <template #header>
            <div class="card-header">
              <span>用户列表</span>
              <el-button type="primary" @click="handleAdd">
                <el-icon><PlusFilled /></el-icon> 添加用户
              </el-button>
            </div>
          </template>

          <el-table :data="userList" style="width: 100%" v-loading="loading" stripe>
            <el-table-column prop="id" label="ID" width="80" />
            <el-table-column prop="username" label="用户名" width="200" />
            <el-table-column prop="email" label="邮箱" width="250" />
            <el-table-column prop="role" label="角色" width="120">
              <template #default="{ row }">
                <el-tag :type="row.role === 'admin' ? 'danger' : 'primary'">
                  {{ row.role === 'admin' ? '管理员' : '普通用户' }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="status" label="状态" width="100">
              <template #default="{ row }">
                <el-tag :type="row.status === 1 ? 'success' : 'danger'">
                  {{ row.status === 1 ? '启用' : '禁用' }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="createdAt" label="创建时间" width="180" />
            <el-table-column label="操作" width="220" fixed="right">
              <template #default="{ row }">
                <el-button type="primary" size="small" @click="handleEdit(row)">
                  编辑
                </el-button>
                <el-button 
                  :type="row.status === 1 ? 'warning' : 'success'" 
                  size="small" 
                  @click="handleToggleStatus(row)"
                >
                  {{ row.status === 1 ? '禁用' : '启用' }}
                </el-button>
                <el-button type="danger" size="small" @click="handleDelete(row)" :disabled="row.id === 1">
                  删除
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-tab-pane>

      <!-- 个人信息 -->
      <el-tab-pane label="个人信息" name="profile">
        <el-card shadow="never">
          <template #header>
            <span>个人信息</span>
          </template>
          <el-form :model="profileForm" label-width="120px" style="max-width: 600px">
            <el-form-item label="用户名">
              <el-input v-model="profileForm.username" disabled />
            </el-form-item>
            <el-form-item label="昵称">
              <el-input v-model="profileForm.nickname" placeholder="请输入昵称" />
            </el-form-item>
            <el-form-item label="邮箱">
              <el-input v-model="profileForm.email" placeholder="请输入邮箱" />
            </el-form-item>
            <el-form-item label="手机号">
              <el-input v-model="profileForm.phone" placeholder="请输入手机号" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="saveProfile">保存信息</el-button>
              <el-button @click="resetProfile">重置</el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-tab-pane>

      <!-- 修改密码 -->
      <el-tab-pane label="修改密码" name="password">
        <el-card shadow="never">
          <template #header>
            <span>修改密码</span>
          </template>
          <el-form :model="passwordForm" :rules="passwordRules" ref="passwordFormRef" label-width="120px" style="max-width: 600px">
            <el-form-item label="当前密码" prop="currentPassword">
              <el-input v-model="passwordForm.currentPassword" type="password" placeholder="请输入当前密码" show-password />
            </el-form-item>
            <el-form-item label="新密码" prop="newPassword">
              <el-input v-model="passwordForm.newPassword" type="password" placeholder="请输入新密码" show-password />
            </el-form-item>
            <el-form-item label="确认密码" prop="confirmPassword">
              <el-input v-model="passwordForm.confirmPassword" type="password" placeholder="请再次输入新密码" show-password />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="changePassword">修改密码</el-button>
              <el-button @click="resetPassword">重置</el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-tab-pane>
    </el-tabs>

    <!-- 添加/编辑用户对话框 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form :model="form" :rules="rules" ref="formRef" label-width="80px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="form.username" placeholder="请输入用户名" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="form.email" placeholder="请输入邮箱" />
        </el-form-item>
        <el-form-item label="密码" prop="password" v-if="!form.id">
          <el-input v-model="form.password" type="password" placeholder="请输入密码" show-password />
        </el-form-item>
        <el-form-item label="角色" prop="role">
          <el-select v-model="form.role" placeholder="请选择角色" style="width: 100%">
            <el-option label="管理员" value="admin" />
            <el-option label="普通用户" value="user" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-switch v-model="form.status" :active-value="1" :inactive-value="0" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const activeTab = ref('list')
const loading = ref(false)
const userList = ref([])
const dialogVisible = ref(false)
const dialogTitle = ref('')
const formRef = ref(null)

const form = ref({
  id: null,
  username: '',
  email: '',
  password: '',
  role: 'user',
  status: 1
})

const rules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 20, message: '用户名长度在 3-20 个字符', trigger: 'blur' }
  ],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度至少 6 位', trigger: 'blur' }
  ],
  role: [
    { required: true, message: '请选择角色', trigger: 'change' }
  ]
}

// 个人信息表单
const profileForm = ref({
  username: localStorage.getItem('username') || 'admin',
  nickname: '',
  email: '',
  phone: ''
})

// 密码修改表单
const passwordForm = ref({
  currentPassword: '',
  newPassword: '',
  confirmPassword: ''
})

const passwordFormRef = ref(null)

const passwordRules = {
  currentPassword: [
    { required: true, message: '请输入当前密码', trigger: 'blur' }
  ],
  newPassword: [
    { required: true, message: '请输入新密码', trigger: 'blur' },
    { min: 6, message: '密码长度至少 6 位', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请再次输入新密码', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (value !== passwordForm.value.newPassword) {
          callback(new Error('两次输入的密码不一致'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ]
}

// 保存个人信息
const saveProfile = () => {
  ElMessage.success('个人信息保存成功')
}

// 重置个人信息
const resetProfile = () => {
  profileForm.value = {
    username: localStorage.getItem('username') || 'admin',
    nickname: '',
    email: '',
    phone: ''
  }
}

// 修改密码
const changePassword = async () => {
  if (!passwordFormRef.value) return
  
  await passwordFormRef.value.validate(async (valid) => {
    if (valid) {
      // TODO: 调用后端 API 修改密码
      ElMessage.success('密码修改成功，请重新登录')
      passwordForm.value = {
        currentPassword: '',
        newPassword: '',
        confirmPassword: ''
      }
      // 退出登录
      setTimeout(() => {
        localStorage.removeItem('token')
        localStorage.removeItem('username')
        window.location.href = '/login'
      }, 1500)
    }
  })
}

// 重置密码表单
const resetPassword = () => {
  passwordForm.value = {
    currentPassword: '',
    newPassword: '',
    confirmPassword: ''
  }
  if (passwordFormRef.value) {
    passwordFormRef.value.clearValidate()
  }
}

// 模拟用户数据
const loadUsers = () => {
  loading.value = true
  setTimeout(() => {
    userList.value = [
      {
        id: 1,
        username: 'admin',
        email: 'admin@example.com',
        role: 'admin',
        status: 1,
        createdAt: '2024-01-01 10:00:00'
      },
      {
        id: 2,
        username: 'user1',
        email: 'user1@example.com',
        role: 'user',
        status: 1,
        createdAt: '2024-01-02 10:00:00'
      },
      {
        id: 3,
        username: 'user2',
        email: 'user2@example.com',
        role: 'user',
        status: 0,
        createdAt: '2024-01-03 10:00:00'
      }
    ]
    loading.value = false
  }, 300)
}

const handleAdd = () => {
  dialogTitle.value = '添加用户'
  form.value = {
    id: null,
    username: '',
    email: '',
    password: '',
    role: 'user',
    status: 1
  }
  dialogVisible.value = true
}

const handleEdit = (row) => {
  dialogTitle.value = '编辑用户'
  form.value = {
    ...row,
    password: ''
  }
  dialogVisible.value = true
}

const handleSubmit = () => {
  if (!formRef.value) return
  formRef.value.validate((valid) => {
    if (valid) {
      if (form.value.id) {
        // 编辑用户
        const index = userList.value.findIndex(u => u.id === form.value.id)
        if (index !== -1) {
          userList.value[index] = {
            ...userList.value[index],
            username: form.value.username,
            email: form.value.email,
            role: form.value.role,
            status: form.value.status
          }
        }
        ElMessage.success('更新成功')
      } else {
        // 添加用户
        const newUser = {
          id: userList.value.length + 1,
          username: form.value.username,
          email: form.value.email,
          role: form.value.role,
          status: form.value.status,
          createdAt: new Date().toLocaleString('zh-CN')
        }
        userList.value.push(newUser)
        ElMessage.success('添加成功')
      }
      dialogVisible.value = false
    }
  })
}

const handleDelete = (row) => {
  if (row.id === 1) {
    ElMessage.warning('不能删除管理员账户')
    return
  }
  
  ElMessageBox.confirm('确定要删除该用户吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    userList.value = userList.value.filter(u => u.id !== row.id)
    ElMessage.success('删除成功')
  })
}

const handleToggleStatus = (row) => {
  const newStatus = row.status === 1 ? 0 : 1
  const action = newStatus === 1 ? '启用' : '禁用'
  
  ElMessageBox.confirm(`确定要${action}该用户吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    row.status = newStatus
    ElMessage.success(`${action}成功`)
  })
}

onMounted(() => {
  loadUsers()
})
</script>

<style scoped>
.user-management {
  height: 100%;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

:deep(.el-table) {
  font-size: 14px;
}

:deep(.el-table th) {
  background-color: #f5f7fa;
  color: #606266;
  font-weight: 600;
}

:deep(.el-table--striped .el-table__row--striped td) {
  background: #fafafa;
}

:deep(.el-table td) {
  color: #303133 !important;
}
</style>
