import axios from 'axios'
import { ElMessage } from 'element-plus'

const request = axios.create({
  baseURL: '/api',
  timeout: 10000
})

request.interceptors.request.use(
  config => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  error => {
    ElMessage.error('请求发送失败，请检查网络连接')
    return Promise.reject(error)
  }
)

request.interceptors.response.use(
  response => {
    return response.data
  },
  error => {
    if (error.response?.status === 401) {
      // 清除登录状态
      localStorage.removeItem('token')
      localStorage.removeItem('username')
      localStorage.removeItem('isLoggedIn')
      localStorage.removeItem('userInfo')
      
      // 跳转到登录页面
      window.location.href = '/login'
    } else if (error.response?.status === 400) {
      // 处理400错误，通常是请求参数或文件上传问题
      ElMessage.error(error.response?.data?.error || '请求参数错误')
    } else if (error.response?.status === 500) {
      // 处理500错误，通常是服务器内部问题
      ElMessage.error(error.response?.data?.error || '服务器内部错误')
    } else {
      // 其他错误
      ElMessage.error('请求失败，请稍后重试')
    }
    return Promise.reject(error)
  }
)

export default request
