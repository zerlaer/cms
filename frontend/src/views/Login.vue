<template>
  <div class="login-container">
    <div class="login-box">
      <div class="login-header">
        <div class="logo-container">
                  <div class="feather-chip-logo">
                    <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="#ffffff" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <rect x="4" y="4" width="16" height="16" rx="2" ry="2"></rect>
                      <rect x="9" y="9" width="6" height="6"></rect>
                      <line x1="9" y1="1" x2="9" y2="4"></line>
                      <line x1="15" y1="1" x2="15" y2="4"></line>
                      <line x1="9" y1="20" x2="9" y2="23"></line>
                      <line x1="15" y1="20" x2="15" y2="23"></line>
                      <line x1="20" y1="9" x2="23" y2="9"></line>
                      <line x1="20" y1="14" x2="23" y2="14"></line>
                      <line x1="1" y1="9" x2="4" y2="9"></line>
                      <line x1="1" y1="14" x2="4" y2="14"></line>
                    </svg>
                  </div>
                </div>
        <h1>电子元件管理系统</h1>
        <p>Component Management System</p>
      </div>
      
      <el-form
        ref="loginFormRef"
        :model="loginForm"
        :rules="loginRules"
        class="login-form"
      >
        <el-form-item prop="username">
                  <el-input
                    v-model="loginForm.username"
                    placeholder="请输入用户名"
                    size="large"
                    :class="{'input-focus': usernameFocus}"
                    @focus="usernameFocus = true"
                    @blur="usernameFocus = false"
                  >
                    <template #prefix>
                      <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#606266" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="el-input__icon">
                        <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                        <circle cx="12" cy="7" r="4"></circle>
                      </svg>
                    </template>
                  </el-input>
                </el-form-item>
                
                <el-form-item prop="password">
                  <el-input
                    v-model="loginForm.password"
                    type="password"
                    placeholder="请输入密码"
                    size="large"
                    show-password
                    @keyup.enter="handleLogin"
                    :class="{'input-focus': passwordFocus}"
                    @focus="passwordFocus = true"
                    @blur="passwordFocus = false"
                  >
                    <template #prefix>
                      <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#606266" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="el-input__icon">
                        <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
                        <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
                      </svg>
                    </template>
                  </el-input>
                </el-form-item>
        
        <el-form-item>
          <el-button
            type="primary"
            size="large"
            class="login-button"
            :loading="loading"
            @click="handleLogin"
          >
            登录
          </el-button>
        </el-form-item>
      </el-form>
      
      <div class="login-tip">
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { login } from '@/api/user'

const router = useRouter()
const loading = ref(false)

const loginForm = ref({
  username: '',
  password: ''
})

const usernameFocus = ref(false)
const passwordFocus = ref(false)

const loginRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度至少 6 位', trigger: 'blur' }
  ]
}

const loginFormRef = ref(null)

const handleLogin = async () => {
  if (!loginFormRef.value) return
  
  await loginFormRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      
      try {
        const res = await login(loginForm.value)
        localStorage.setItem('token', res.token)
        localStorage.setItem('isLoggedIn', 'true')
        localStorage.setItem('username', res.user.username)
        localStorage.setItem('userInfo', JSON.stringify(res.user))
        
        ElMessage.success('登录成功')
        
        setTimeout(() => {
          router.push('/')
        }, 500)
      } catch (error) {
        console.error('Login error:', error)
        // 错误已经在request拦截器中处理
      } finally {
        loading.value = false
      }
    }
  })
}
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f5f7fa;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}

.login-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(74, 111, 165, 0.1) 0%, rgba(44, 62, 80, 0.1) 100%);
}

.login-box {
  width: 400px;
  padding: 40px;
  background: #ffffff;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
  border: 1px solid #e8e8e8;
  z-index: 1;
}

.login-box:hover {
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.15);
}

.login-header {
  text-align: center;
  margin-bottom: 32px;
}

.login-header h1 {
  font-size: 22px;
  font-weight: 600;
  margin: 0 0 8px;
  color: #304156;
}

.login-header p {
  font-size: 14px;
  color: #bfcbd9;
  margin: 0;
  opacity: 0.8;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(-20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.logo-container {
  width: 80px;
  height: 80px;
  margin: 0 auto 24px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: visible;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.feather-chip-logo {
  width: 64px;
  height: 64px;
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
}

.feather-chip-logo svg {
  width: 40px;
  height: 40px;
  stroke: #ffffff;
  stroke-width: 2;
  stroke-linecap: round;
  stroke-linejoin: round;
  fill: none;
}

@keyframes pulse {
  0% {
    box-shadow: 0 10px 30px rgba(102, 126, 234, 0.5);
  }
  50% {
    box-shadow: 0 15px 40px rgba(102, 126, 234, 0.7);
  }
  100% {
    box-shadow: 0 10px 30px rgba(102, 126, 234, 0.5);
  }
}

.login-header h1 {
  margin: 20px 0 10px;
  font-size: 24px;
  color: #304156;
  font-weight: 600;
}

.login-header p {
  color: #bfcbd9;
  font-size: 14px;
  letter-spacing: 0.5px;
}

.login-form {
  margin-top: 30px;
  animation: slideIn 1s ease 0.3s both;
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

:deep(.el-input) {
  border-radius: 8px;
  transition: all 0.3s ease;
}

:deep(.el-input__wrapper) {
  border-radius: 8px;
  transition: all 0.3s ease;
  background-color: #ffffff !important;
  border: 1px solid #e8e8e8 !important;
}

:deep(.el-input__wrapper:hover) {
  border-color: #409EFF !important;
}

:deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 2px rgba(64, 158, 255, 0.2);
  border-color: #409EFF !important;
  background-color: #ffffff !important;
}

.login-button {
  width: 100%;
  border-radius: 8px;
  padding: 12px 0;
  font-size: 16px;
  font-weight: 600;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: 1px solid #667eea;
  color: #ffffff;
  transition: all 0.3s ease;
}

.login-button:hover {
  background: linear-gradient(135deg, #5a6edb 0%, #6a4290 100%);
  border-color: #5a6edb;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
  transform: translateY(-1px);
}

.login-button:active {
  transform: translateY(0);
}

.login-button:disabled {
  background: #c0c4cc;
  border-color: #c0c4cc;
  cursor: not-allowed;
}

.login-tip {
  margin-top: 20px;
  text-align: center;
  color: #bfcbd9;
  font-size: 13px;
  animation: fadeIn 1s ease 0.6s both;
}

.login-tip p {
  margin: 0;
  letter-spacing: 0.5px;
}

/* 响应式设计 */
@media (max-width: 480px) {
  .login-box {
    width: 90%;
    padding: 30px;
  }
  
  .login-header h1 {
    font-size: 20px;
  }
}
</style>
