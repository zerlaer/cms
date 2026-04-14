<template>
  <div class="system-settings">
    <el-card>
      <template #header>
        <span>系统设置</span>
      </template>
      <el-tabs v-model="activeTab" type="border-card">
        <!-- 基本设置 -->
        <el-tab-pane label="基本设置" name="basic">
          <el-form :model="basicSettings" label-width="120px" style="max-width: 600px">
            <el-form-item label="系统名称">
              <el-input v-model="basicSettings.systemName" />
            </el-form-item>
            <el-form-item label="公司名称">
              <el-input v-model="basicSettings.companyName" />
            </el-form-item>
            <el-form-item label="联系人">
              <el-input v-model="basicSettings.contact" />
            </el-form-item>
            <el-form-item label="联系电话">
              <el-input v-model="basicSettings.phone" />
            </el-form-item>
            <el-form-item label="联系地址">
              <el-input v-model="basicSettings.address" type="textarea" :rows="3" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="saveBasicSettings">保存设置</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>

        <!-- 系统信息 -->
        <el-tab-pane label="系统信息" name="system">
          <el-descriptions :column="1" border>
            <el-descriptions-item label="系统版本">
              <el-tag type="success" size="large">v1.0.0</el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="系统名称">{{ basicSettings.systemName }}</el-descriptions-item>
            <el-descriptions-item label="公司名称">{{ basicSettings.companyName }}</el-descriptions-item>
            <el-descriptions-item label="操作系统">
              {{ osInfo }}
            </el-descriptions-item>
            <el-descriptions-item label="后端框架">
              <el-tag type="info">Go 1.21+</el-tag>
              <el-tag type="info" style="margin-left: 8px">Gin v1.9+</el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="前端框架">
              <el-tag type="warning">Vue 3</el-tag>
              <el-tag type="warning" style="margin-left: 8px">Element Plus</el-tag>
              <el-tag type="warning" style="margin-left: 8px">Vite</el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="数据库">
              <el-tag type="primary">MySQL 8.0+</el-tag>
              <el-tag type="primary" style="margin-left: 8px">GORM</el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="中间件">
              <el-tag type="danger">Viper</el-tag>
              <el-tag type="danger" style="margin-left: 8px">Zap</el-tag>
              <el-tag type="danger" style="margin-left: 8px">Excelize</el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="系统时间">{{ currentTime }}</el-descriptions-item>
            <el-descriptions-item label="运行时间">
              <span style="color: #67c23a">●</span> 正常运行中
            </el-descriptions-item>
            <el-descriptions-item label="浏览器信息">
              {{ browserInfo }}
            </el-descriptions-item>
            <el-descriptions-item label="屏幕分辨率">
              {{ screenResolution }}
            </el-descriptions-item>
          </el-descriptions>
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'

const activeTab = ref('basic')
const currentTime = ref('')
const browserInfo = ref('')
const screenResolution = ref('')
const osInfo = ref('')

const basicSettings = ref({
  systemName: '电子元件管理系统',
  companyName: '示例公司',
  contact: '张三',
  phone: '13800138000',
  address: '北京市朝阳区某某街道某某号'
})

const securitySettings = ref({
  currentPassword: '',
  newPassword: '',
  confirmPassword: ''
})

const profileSettings = ref({
  username: 'admin',
  nickname: '管理员',
  email: 'admin@example.com',
  phone: '13800138000'
})

// 获取浏览器信息
const getBrowserInfo = () => {
  const userAgent = navigator.userAgent
  let browser = 'Unknown Browser'
  
  if (userAgent.indexOf('Chrome') > -1) {
    browser = 'Chrome'
  } else if (userAgent.indexOf('Firefox') > -1) {
    browser = 'Firefox'
  } else if (userAgent.indexOf('Safari') > -1) {
    browser = 'Safari'
  } else if (userAgent.indexOf('Edge') > -1) {
    browser = 'Edge'
  }
  
  browserInfo.value = `${browser} (${navigator.platform})`
}

// 获取操作系统信息
const getOSInfo = () => {
  const platform = navigator.platform
  const userAgent = navigator.userAgent
  
  if (platform.indexOf('Win') > -1) {
    if (userAgent.indexOf('Windows NT 10.0') > -1) {
      osInfo.value = 'Windows 10/11'
    } else if (userAgent.indexOf('Windows NT 6.3') > -1) {
      osInfo.value = 'Windows 8.1'
    } else if (userAgent.indexOf('Windows NT 6.2') > -1) {
      osInfo.value = 'Windows 8'
    } else if (userAgent.indexOf('Windows NT 6.1') > -1) {
      osInfo.value = 'Windows 7'
    } else {
      osInfo.value = 'Windows'
    }
  } else if (platform.indexOf('Mac') > -1) {
    osInfo.value = 'macOS'
  } else if (platform.indexOf('Linux') > -1) {
    osInfo.value = 'Linux'
  } else if (platform.indexOf('iPhone') > -1 || platform.indexOf('iPad') > -1) {
    osInfo.value = 'iOS'
  } else if (platform.indexOf('Android') > -1) {
    osInfo.value = 'Android'
  } else {
    osInfo.value = 'Unknown OS'
  }
}

// 获取屏幕分辨率
const getScreenResolution = () => {
  screenResolution.value = `${window.screen.width} x ${window.screen.height}`
}

// 更新时间
const updateTime = () => {
  const now = new Date()
  currentTime.value = now.toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
    hour12: false
  })
}

const saveBasicSettings = () => {
  ElMessage.success('基本设置保存成功')
}

const saveSecuritySettings = () => {
  if (!securitySettings.value.currentPassword) {
    ElMessage.warning('请输入当前密码')
    return
  }
  if (!securitySettings.value.newPassword) {
    ElMessage.warning('请输入新密码')
    return
  }
  if (securitySettings.value.newPassword !== securitySettings.value.confirmPassword) {
    ElMessage.error('两次输入的密码不一致')
    return
  }
  if (securitySettings.value.newPassword.length < 6) {
    ElMessage.warning('密码长度至少 6 位')
    return
  }
  ElMessage.success('密码修改成功')
  securitySettings.value = {
    currentPassword: '',
    newPassword: '',
    confirmPassword: ''
  }
}

const saveProfileSettings = () => {
  ElMessage.success('个人信息保存成功')
}

onMounted(() => {
  const updateTime = () => {
    const now = new Date()
    currentTime.value = now.toLocaleString('zh-CN', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit'
    })
  }
  updateTime()
  setInterval(updateTime, 1000)
  
  // 获取浏览器和操作系统信息
  getBrowserInfo()
  getOSInfo()
  getScreenResolution()
})
</script>

<style scoped>
.system-settings {
  padding: 20px;
}

:deep(.el-tabs__content) {
  padding: 20px;
}

:deep(.el-form-item__label) {
  font-weight: 600;
}
</style>
