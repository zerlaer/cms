<template>
  <div class="inventory-stats">
    <el-row :gutter="20">
      <!-- 统计卡片 -->
      <el-col :span="6">
        <el-card class="stat-card" shadow="hover" @click.native="handleStatClick('all')" style="cursor: pointer; padding: 0;">
          <div class="stat-icon" style="background-color: #409EFF; color: white; border-radius: 12px; margin: 15px;">
            <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"></path>
              <polyline points="3.27 6.96 12 12.01 20.73 6.96"></polyline>
              <line x1="12" y1="22.08" x2="12" y2="12"></line>
            </svg>
          </div>
        </el-card>
      </el-col>

      <el-col :span="6">
        <el-card class="stat-card" shadow="hover" @click.native="handleStatClick('hasStock')" style="cursor: pointer; padding: 0;">
          <div class="stat-icon" style="background-color: #67c23a; color: white; border-radius: 12px; margin: 15px;">
            <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
              <polyline points="22 4 12 14.01 9 11.01"></polyline>
            </svg>
          </div>
        </el-card>
      </el-col>

      <el-col :span="6">
        <el-card class="stat-card" shadow="hover" @click.native="handleStatClick('lowStock')" style="cursor: pointer; padding: 0;">
          <div class="stat-icon" style="background-color: #e6a23c; color: white; border-radius: 12px; margin: 15px;">
            <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"></path>
              <line x1="12" y1="9" x2="12" y2="13"></line>
              <line x1="12" y1="17" x2="12.01" y2="17"></line>
            </svg>
          </div>
        </el-card>
      </el-col>

      <el-col :span="6">
        <el-card class="stat-card" shadow="hover" @click.native="handleStatClick('noStock')" style="cursor: pointer; padding: 0;">
          <div class="stat-icon" style="background-color: #f56c6c; color: white; border-radius: 12px; margin: 15px;">
            <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="12" cy="12" r="10"></circle>
              <line x1="15" y1="9" x2="9" y2="15"></line>
              <line x1="9" y1="9" x2="15" y2="15"></line>
            </svg>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 库存分类统计 -->
    <el-row :gutter="20" style="margin-top: 20px">
      <el-col :span="24">
        <el-card>
          <template #header>
            <span>库存分类统计</span>
          </template>
          <div class="chart-container" style="height: 400px">
            <el-chart :option="chartOption" />
          </div>
        </el-card>
      </el-col>


    </el-row>

    <!-- 库存明细表格 -->
    <el-card style="margin-top: 20px">
      <template #header>
        <div class="card-header">
          <span>库存明细</span>
          <el-input
            v-model="searchQuery"
            placeholder="搜索商品名称、编号"
            clearable
            style="width: 300px"
            @input="handleSearch"
          >
            <template #prefix>
              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="#606266" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="11" cy="11" r="8"></circle>
                <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
              </svg>
            </template>
          </el-input>
        </div>
      </template>

      <div class="table-container">
        <el-table :data="tableData" style="width: 100%" v-loading="loading" stripe border :cell-style="{ padding: '5px 0' }" :header-cell-style="{ padding: '8px 0' }">
        <el-table-column type="index" label="序号" width="60" align="center" />
        <el-table-column prop="productCode" label="编号" min-width="100" align="center" />
        <el-table-column prop="source" label="渠道" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="row.source === 'LCSC' ? 'success' : 'primary'" size="small">
              {{ row.source === 'LCSC' ? '立创商城' : '淘宝' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="brand" label="品牌" min-width="80" align="center" />
        <el-table-column prop="model" label="型号" min-width="120" align="center" />
        <el-table-column prop="package" label="封装" min-width="80" align="center" />
        <el-table-column prop="name" label="商品" show-overflow-tooltip min-width="200" align="center" />
        <el-table-column prop="quantity" label="数量" width="80" align="center" />
        <el-table-column prop="currentStock" label="库存" width="130" align="center">
          <template #default="{ row }">
            <div class="stock-highlight">
              <el-tag :type="row.currentStock > 50 ? 'success' : row.currentStock > 0 ? 'warning' : 'danger'" size="large">
                {{ row.currentStock }}
              </el-tag>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="price" label="总价" width="110" align="center">
          <template #default="{ row }">
            ¥{{ row.price.toFixed(2) }}
          </template>
        </el-table-column>
        <el-table-column label="单价" width="110" align="center">
          <template #default="{ row }">
            ¥{{ (row.quantity > 0 ? row.price / row.quantity : 0).toFixed(2) }}
          </template>
        </el-table-column>
        </el-table>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { getComponents } from '@/api/component'
import { useRouter } from 'vue-router'
import { PieChart } from 'echarts'

const router = useRouter()
const loading = ref(false)
const tableData = ref([])
const searchQuery = ref('')

// 统计数据
const totalComponents = ref(0)
const totalValue = ref(0)
const inStockCount = ref(0)
const lowStockCount = ref(0)
const noStockCount = ref(0)

// 图表数据
const chartOption = ref({
  title: {
    text: '库存分类统计',
    left: 'center'
  },
  tooltip: {
    trigger: 'item',
    formatter: '{a} <br/>{b}: {c} ({d}%)'
  },
  legend: {
    orient: 'vertical',
    left: 'left',
    data: []
  },
  series: [
    {
      name: '库存分类',
      type: 'pie',
      radius: '50%',
      center: ['50%', '60%'],
      data: [],
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      }
    }
  ]
})

// 分类统计数据
const categoryStats = ref([
  { label: '电容/电阻/电感', value: 0, percentage: 0, color: '#67c23a' },
  { label: '微控制器/逻辑器件', value: 0, percentage: 0, color: '#409eff' },
  { label: '二极管/晶体管', value: 0, percentage: 0, color: '#e6a23c' },
  { label: '板级电路保护/电源管理', value: 0, percentage: 0, color: '#f56c6c' },
  { label: '接口芯片/时钟和计时', value: 0, percentage: 0, color: '#909399' },
  { label: '连接器/端子/开关', value: 0, percentage: 0, color: '#722ed1' },
  { label: '数据转换芯片/射频无线', value: 0, percentage: 0, color: '#13c2c2' },
  { label: '光电器件/显示模块', value: 0, percentage: 0, color: '#eb2f96' },
  { label: '滤波器/晶体/振荡器', value: 0, percentage: 0, color: '#faad14' },
  { label: '存储器/传感器/隔离器件', value: 0, percentage: 0, color: '#52c41a' },
  { label: '继电器/蜂鸣器/马达', value: 0, percentage: 0, color: '#1890ff' },
  { label: '功能模块/通信模块', value: 0, percentage: 0, color: '#fa541c' },
  { label: '放大器/开发板', value: 0, percentage: 0, color: '#722ed1' },
  { label: '工具/仪器仪表/耗材', value: 0, percentage: 0, color: '#13c2c2' }
])

// 最近入库记录
const recentStockIn = ref([
  { productCode: 'RES-001', name: '贴片电阻', quantity: 1000, date: '2024-01-15' },
  { productCode: 'CAP-001', name: '贴片电容', quantity: 5000, date: '2024-01-14' },
  { productCode: 'IC-001', name: '运算放大器', quantity: 500, date: '2024-01-13' }
])

// 处理统计卡片点击
const handleStatClick = (type) => {
  // 跳转到元件列表页面
  router.push({ 
    path: '/', 
    query: { 
      filter: type === 'all' ? '' : type 
    } 
  })
}

const loadData = async () => {
  loading.value = true
  try {
    // 获取所有元件数据
    const res = await getComponents({ page: 1, pageSize: 1000 })
    tableData.value = res.data || []
    
    // 计算统计数据
    calculateStats()
  } catch (error) {
  } finally {
    loading.value = false
  }
}

const calculateStats = () => {
  const data = tableData.value
  
  // 元件总数
  totalComponents.value = data.length
  
  // 库存总价值
  totalValue.value = data.reduce((sum, item) => sum + (item.price * item.currentStock), 0).toFixed(2)
  
  // 有库存元件数量
  inStockCount.value = data.filter(item => item.currentStock > 0).length
  
  // 库存不足（少于 50）
  lowStockCount.value = data.filter(item => item.currentStock > 0 && item.currentStock < 50).length
  
  // 无库存
  noStockCount.value = data.filter(item => item.currentStock === 0).length
  
  // 更新分类统计
  updateCategoryStats()
  
  // 更新最近入库记录（这里用最后添加的元件模拟）
  updateRecentStockIn()
}

const updateCategoryStats = () => {
  const data = tableData.value
  const total = data.length
  
  if (total === 0) {
    categoryStats.value.forEach(stat => {
      stat.value = 0
      stat.percentage = 0
    })
    return
  }
  
  // 新的分类体系
  const categories = {
    '电容/电阻/电感': 0,
    '微控制器/逻辑器件': 0,
    '二极管/晶体管': 0,
    '板级电路保护/电源管理': 0,
    '接口芯片/时钟和计时': 0,
    '连接器/端子/开关': 0,
    '数据转换芯片/射频无线': 0,
    '光电器件/显示模块': 0,
    '滤波器/晶体/振荡器': 0,
    '存储器/传感器/隔离器件': 0,
    '继电器/蜂鸣器/马达': 0,
    '功能模块/通信模块': 0,
    '放大器/开发板': 0,
    '工具/仪器仪表/耗材': 0
  }
  
  data.forEach(item => {
    const name = item.name.toLowerCase()
    const brand = item.brand.toLowerCase()
    const model = item.model.toLowerCase()
    
    // 电容/电阻/电感
    if (name.includes('电阻') || name.includes('电容') || name.includes('电感') || 
        name.includes('resistor') || name.includes('capacitor') || name.includes('inductor') ||
        brand.includes('res') || brand.includes('cap') || brand.includes('ind')) {
      categories['电容/电阻/电感']++
    }
    // 微控制器/逻辑器件
    else if (name.includes('微控制器') || name.includes('单片机') || name.includes('mcu') || 
             name.includes('controller') || name.includes('logic') || name.includes('芯片') ||
             brand.includes('mcu') || brand.includes('stm') || brand.includes('arduino')) {
      categories['微控制器/逻辑器件']++
    }
    // 二极管/晶体管
    else if (name.includes('二极管') || name.includes('晶体管') || name.includes('diode') || 
             name.includes('transistor') || brand.includes('diode') || brand.includes('transistor')) {
      categories['二极管/晶体管']++
    }
    // 板级电路保护/电源管理
    else if (name.includes('保护') || name.includes('电源') || name.includes('power') || 
             name.includes('protection') || brand.includes('power') || brand.includes('protect')) {
      categories['板级电路保护/电源管理']++
    }
    // 接口芯片/时钟和计时
    else if (name.includes('接口') || name.includes('时钟') || name.includes('接口') || 
             name.includes('clock') || name.includes('interface') || brand.includes('clock')) {
      categories['接口芯片/时钟和计时']++
    }
    // 连接器/端子/开关
    else if (name.includes('连接器') || name.includes('端子') || name.includes('开关') || 
             name.includes('connector') || name.includes('switch') || name.includes('terminal')) {
      categories['连接器/端子/开关']++
    }
    // 数据转换芯片/射频无线
    else if (name.includes('转换') || name.includes('射频') || name.includes('无线') || 
             name.includes('convert') || name.includes('rf') || name.includes('wireless')) {
      categories['数据转换芯片/射频无线']++
    }
    // 光电器件/显示模块
    else if (name.includes('光电') || name.includes('显示') || name.includes('led') || 
             name.includes('display') || name.includes('optical') || brand.includes('led')) {
      categories['光电器件/显示模块']++
    }
    // 滤波器/晶体/振荡器
    else if (name.includes('滤波') || name.includes('晶体') || name.includes('振荡器') || 
             name.includes('filter') || name.includes('crystal') || name.includes('oscillator')) {
      categories['滤波器/晶体/振荡器']++
    }
    // 存储器/传感器/隔离器件
    else if (name.includes('存储') || name.includes('传感器') || name.includes('隔离') || 
             name.includes('memory') || name.includes('sensor') || name.includes('isolation')) {
      categories['存储器/传感器/隔离器件']++
    }
    // 继电器/蜂鸣器/马达
    else if (name.includes('继电器') || name.includes('蜂鸣器') || name.includes('马达') || 
             name.includes('relay') || name.includes('buzzer') || name.includes('motor')) {
      categories['继电器/蜂鸣器/马达']++
    }
    // 功能模块/通信模块
    else if (name.includes('模块') || name.includes('通信') || name.includes('module') || 
             name.includes('communication') || brand.includes('module')) {
      categories['功能模块/通信模块']++
    }
    // 放大器/开发板
    else if (name.includes('放大器') || name.includes('开发板') || name.includes('amplifier') || 
             name.includes('board') || brand.includes('amp') || brand.includes('board')) {
      categories['放大器/开发板']++
    }
    // 工具/仪器仪表/耗材
    else if (name.includes('工具') || name.includes('仪器') || name.includes('仪表') || 
             name.includes('耗材') || name.includes('tool') || name.includes('meter')) {
      categories['工具/仪器仪表/耗材']++
    }
    // 默认为第一个分类
    else {
      categories['电容/电阻/电感']++
    }
  })
  
  // 更新分类统计数据
  categoryStats.value.forEach((stat, index) => {
    const label = stat.label
    stat.value = categories[label]
    stat.percentage = parseFloat(((categories[label] / total) * 100).toFixed(1))
  })
}

const updateRecentStockIn = () => {
  // 取最新的 5 个元件作为最近入库记录（实际项目中应该有专门的入库记录表）
  const recent = tableData.value.slice(-5).reverse()
  recentStockIn.value = recent.map(item => ({
    productCode: item.productCode,
    name: item.name,
    quantity: item.currentStock,
    date: item.createdAt ? new Date(item.createdAt).toLocaleDateString() : new Date().toLocaleDateString()
  }))
}

const handleSearch = () => {
  loadData()
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.inventory-stats {
  padding: 15px;
  width: 100%;
}

.stat-card {
  border-radius: 8px;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.stat-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.stat-icon {
  width: 80px;
  height: 80px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.stat-info {
  flex: 1;
}

.stat-value {
  font-size: 28px;
  font-weight: bold;
  color: #303133;
  margin-bottom: 5px;
  line-height: 1.2;
}

.stat-label {
  font-size: 14px;
  color: #909399;
}

.chart-container {
  padding: 10px;
}

.stat-item {
  margin-bottom: 20px;
  transition: all 0.3s ease;
}

.stat-item:last-child {
  margin-bottom: 0;
}

.stat-item-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
}

.stat-item-label {
  font-size: 14px;
  color: #606266;
  font-weight: 500;
}

.stat-item-value {
  font-size: 14px;
  font-weight: bold;
  color: #303133;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 10px;
}

:deep(.el-table) {
  font-size: 14px;
  border-radius: 8px;
  overflow: hidden;
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

.stock-highlight {
  display: flex;
  justify-content: center;
  align-items: center;
}

:deep(.stock-highlight .el-tag) {
  font-size: 16px;
  font-weight: bold;
  padding: 6px 12px;
  min-width: 60px;
  text-align: center;
  border-radius: 20px;
}

:deep(.el-card) {
  width: 100%;
  border-radius: 8px;
  overflow: hidden;
}

:deep(.el-card__body) {
  padding: 15px;
}

:deep(.el-progress) {
  margin-top: 5px;
}

:deep(.el-progress__text) {
  font-size: 12px;
  color: #909399;
}

/* 移动端响应式设计 */
@media (max-width: 768px) {
  .inventory-stats {
    padding: 10px;
  }
  
  .stat-icon {
    width: 60px;
    height: 60px;
    border-radius: 12px;
  }
  
  .stat-icon svg {
    width: 24px;
    height: 24px;
    stroke: currentColor;
  }
  
  .stat-item {
    margin-bottom: 15px;
  }
  
  .stat-item-label,
  .stat-item-value {
    font-size: 12px;
  }
  
  :deep(.el-table) {
    font-size: 11px;
  }
  
  :deep(.el-table th) {
    font-size: 11px;
    padding: 6px 0 !important;
  }
  
  :deep(.el-table td) {
    padding: 4px 0 !important;
  }
  
  :deep(.stock-highlight .el-tag) {
    font-size: 12px;
    padding: 4px 8px;
    min-width: 40px;
  }
  
  .card-header {
    flex-direction: column;
    align-items: flex-start;
  }
  
  .card-header .el-input {
    width: 100% !important;
  }
  
  :deep(.el-card__body) {
    padding: 10px;
  }
  
  :deep(.el-progress__text) {
    font-size: 10px;
  }
  
  /* 移动端表格滚动 */
  .table-container {
    overflow-x: auto;
    -webkit-overflow-scrolling: touch;
    margin: 0 -10px;
  }
  
  :deep(.el-table) {
    min-width: 800px;
  }
  
  /* 移动端隐藏最近入库记录 */
  .recent-stock-section {
    display: none;
  }
}
</style>
