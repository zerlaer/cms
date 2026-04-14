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

    <!-- 库存状态统计和入库出库记录 -->
    <el-row :gutter="20" style="margin-top: 20px">
      <el-col :span="12">
        <el-card>
          <template #header>
            <span style="display: block; text-align: center; font-weight: bold;">库存状态统计</span>
          </template>
          <div ref="chartRef" class="chart-container" style="height: 350px"></div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card>
          <template #header>
            <span style="display: block; text-align: center; font-weight: bold;">入库出库记录</span>
          </template>
          <div ref="stockChartRef" class="chart-container" style="height: 350px"></div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 更多数据可视化图表 -->
    <el-row :gutter="20" style="margin-top: 20px">
      <el-col :span="12">
        <el-card>
          <template #header>
            <span style="display: block; text-align: center; font-weight: bold;">品牌分布</span>
          </template>
          <div ref="brandChartRef" class="chart-container" style="height: 350px"></div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card>
          <template #header>
            <span style="display: block; text-align: center; font-weight: bold;">封装类型分布</span>
          </template>
          <div ref="packageChartRef" class="chart-container" style="height: 350px"></div>
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
            @keyup.enter="handleSearch"
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
import { ref, computed, onMounted, watch, nextTick } from 'vue'
import { getComponents } from '@/api/component'
import { useRouter } from 'vue-router'
import * as echarts from 'echarts'

const router = useRouter()
const loading = ref(false)
const tableData = ref([])
const searchQuery = ref('')
const chartRef = ref(null)
let chart = null
const stockChartRef = ref(null)
let stockChart = null
const brandChartRef = ref(null)
let brandChart = null
const packageChartRef = ref(null)
let packageChart = null

// 统计数据
const totalComponents = ref(0)
const totalValue = ref(0)
const inStockCount = ref(0)
const lowStockCount = ref(0)
const noStockCount = ref(0)

// 图表数据
const chartOption = ref({
  title: {
    show: false
  },
  tooltip: {
    trigger: 'item',
    formatter: '{a} <br/>{b}: {c} ({d}%)'
  },
  legend: {
    type: 'scroll',
    orient: 'horizontal',
    bottom: '0%',
    data: []
  },
  series: [
    {
      name: '库存状态',
      type: 'pie',
      radius: ['40%', '70%'],
      center: ['50%', '50%'],
      data: [],
      padAngle: 5,
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        },
        label: {
          show: true,
          fontSize: 18,
          fontWeight: 'bold'
        }
      },
      label: {
        show: true,
        formatter: '{b}: {c} ({d}%)'
      },
      labelLine: {
        show: true
      },
      itemStyle: {
        borderColor: '#fff',
        borderWidth: 2,
        borderRadius: 5,
        borderType: 'solid'
      }
    }
  ]
})

// 入库出库记录图表数据
const stockChartOption = ref({
  title: {
    show: false
  },
  tooltip: {
    trigger: 'axis'
  },
  legend: {
    data: ['入库', '出库'],
    bottom: '0%',
    textStyle: {
      fontSize: 12
    }
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '15%',
    containLabel: true
  },
  xAxis: {
    type: 'category',
    boundaryGap: false,
    data: [],
    axisLabel: {
      fontSize: 12,
      rotate: 45
    }
  },
  yAxis: {
    type: 'value',
    axisLabel: {
      fontSize: 12
    }
  },
  series: [
    {
      name: '入库',
      type: 'line',
      data: [],
      smooth: true,
      itemStyle: {
        color: '#67c23a'
      },
      areaStyle: {
        color: {
          type: 'linear',
          x: 0,
          y: 0,
          x2: 0,
          y2: 1,
          colorStops: [{
            offset: 0, color: 'rgba(103, 194, 58, 0.5)'
          }, {
            offset: 1, color: 'rgba(103, 194, 58, 0.1)'
          }]
        }
      }
    },
    {
      name: '出库',
      type: 'line',
      data: [],
      smooth: true,
      itemStyle: {
        color: '#f56c6c'
      },
      areaStyle: {
        color: {
          type: 'linear',
          x: 0,
          y: 0,
          x2: 0,
          y2: 1,
          colorStops: [{
            offset: 0, color: 'rgba(245, 108, 108, 0.5)'
          }, {
            offset: 1, color: 'rgba(245, 108, 108, 0.1)'
          }]
        }
      }
    }
  ]
})

// 品牌分布图表数据
const brandChartOption = ref({
  title: {
    show: false
  },
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'shadow'
    }
  },
  legend: {
    show: false
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
  xAxis: {
    type: 'value',
    axisLine: {
      show: false
    },
    axisTick: {
      show: false
    },
    splitLine: {
      show: false
    },
    axisLabel: {
      show: true,
      fontSize: 12
    }
  },
  yAxis: {
    type: 'category',
    data: [],
    axisLine: {
      show: false
    },
    axisTick: {
      show: false
    },
    axisLabel: {
      show: true,
      fontSize: 12,
      rotate: 30
    }
  },
  series: [
    {
      name: '元件数量',
      type: 'bar',
      data: [],
      itemStyle: {
        color: new echarts.graphic.LinearGradient(0, 0, 1, 0, [
          { offset: 0, color: '#83bff6' },
          { offset: 0.5, color: '#188df0' },
          { offset: 1, color: '#188df0' }
        ])
      },
      label: {
        show: true,
        position: 'right',
        formatter: '{c}',
        fontSize: 12
      }
    }
  ]
})

// 封装类型分布图表数据
const packageChartOption = ref({
  title: {
    show: false
  },
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'shadow'
    }
  },
  legend: {
    show: false
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
  xAxis: {
    type: 'value',
    axisLine: {
      show: false
    },
    axisTick: {
      show: false
    },
    splitLine: {
      show: false
    },
    axisLabel: {
      show: true,
      fontSize: 12
    }
  },
  yAxis: {
    type: 'category',
    data: [],
    axisLine: {
      show: false
    },
    axisTick: {
      show: false
    },
    axisLabel: {
      show: true,
      fontSize: 12,
      rotate: 30
    }
  },
  series: [
    {
      name: '元件数量',
      type: 'bar',
      data: [],
      itemStyle: {
        color: new echarts.graphic.LinearGradient(0, 0, 1, 0, [
          { offset: 0, color: '#fccb05' },
          { offset: 0.5, color: '#f56c6c' },
          { offset: 1, color: '#f56c6c' }
        ])
      },
      label: {
        show: true,
        position: 'right',
        formatter: '{c}',
        fontSize: 12
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
  // 检查登录状态
  const isLoggedIn = localStorage.getItem('isLoggedIn')
  if (!isLoggedIn) {
    // 未登录，跳转到登录页面
    router.push('/login')
    return
  }
  
  loading.value = true
  try {
    // 获取所有元件数据
    const res = await getComponents({ page: 1, pageSize: 1000 })
    let data = res.data || []
    
    // 根据搜索关键词过滤数据
    if (searchQuery.value) {
      const query = searchQuery.value.toLowerCase()
      data = data.filter(item => 
        item.name.toLowerCase().includes(query) || 
        item.productCode.toLowerCase().includes(query) ||
        item.brand.toLowerCase().includes(query) ||
        item.model.toLowerCase().includes(query)
      )
    }
    
    tableData.value = data
    
    // 计算统计数据
    calculateStats()
  } catch (error) {
    // 错误已经在request拦截器中处理
    console.error('Load data error:', error)
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
  
  // 库存不足（少于 10）
  lowStockCount.value = data.filter(item => item.currentStock > 0 && item.currentStock < 10).length
  
  // 无库存
  noStockCount.value = data.filter(item => item.currentStock === 0).length
  
  // 更新分类统计
  updateCategoryStats()
  
  // 更新图表
  nextTick(() => {
    updateChart()
    updateStockChart()
    updateBrandChart()
    updatePackageChart()
  })
  
  // 更新最近入库记录（这里用最后添加的元件模拟）
  updateRecentStockIn()
}

// 初始化图表
const initChart = () => {
  if (chartRef.value) {
    chart = echarts.init(chartRef.value)
    updateChart()
  }
  if (stockChartRef.value) {
    stockChart = echarts.init(stockChartRef.value)
    updateStockChart()
  }
  if (brandChartRef.value) {
    brandChart = echarts.init(brandChartRef.value)
    updateBrandChart()
  }
  if (packageChartRef.value) {
    packageChart = echarts.init(packageChartRef.value)
    updatePackageChart()
  }
}

// 更新图表
const updateChart = () => {
  if (chart) {
    chart.setOption(chartOption.value)
  }
}

// 更新入库出库记录图表
const updateStockChart = () => {
  if (stockChart) {
    stockChart.setOption(stockChartOption.value)
  }
}

// 更新品牌分布图表
const updateBrandChart = () => {
  const data = tableData.value
  
  // 统计每个品牌的元件数量
  const brandCountMap = {}
  data.forEach(item => {
    const brand = item.brand || '其他'
    if (!brandCountMap[brand]) {
      brandCountMap[brand] = 0
    }
    brandCountMap[brand]++
  })
  
  // 转换为图表数据并排序（取前10个）
  const brandData = Object.entries(brandCountMap)
    .map(([brand, count]) => ({ brand, count }))
    .sort((a, b) => {
      // 首韩品牌始终排在第一位
      if (a.brand === '首韩') return -1
      if (b.brand === '首韩') return 1
      // 其他品牌按数量从多到少排序
      return b.count - a.count
    })
    .slice(0, 10)
    // 反转数据顺序，实现旋转180度效果
    .reverse()
  
  // 更新图表数据
  brandChartOption.value.yAxis.data = brandData.map(item => item.brand)
  brandChartOption.value.series[0].data = brandData.map(item => item.count)
  
  if (brandChart) {
    brandChart.setOption(brandChartOption.value)
  }
}

// 更新封装类型分布图表
const updatePackageChart = () => {
  const data = tableData.value
  
  // 统计每个封装类型的元件数量
  const packageCountMap = {}
  data.forEach(item => {
    const packageType = item.package || '其他'
    if (!packageCountMap[packageType]) {
      packageCountMap[packageType] = 0
    }
    packageCountMap[packageType]++
  })
  
  // 转换为图表数据并排序（取前10个）
  const packageData = Object.entries(packageCountMap)
    .map(([packageType, count]) => ({ packageType, count }))
    .sort((a, b) => b.count - a.count)
    .slice(0, 10)
    // 反转数据顺序，实现旋转180度效果
    .reverse()
  
  // 更新图表数据
  packageChartOption.value.yAxis.data = packageData.map(item => item.packageType)
  packageChartOption.value.series[0].data = packageData.map(item => item.count)
  
  if (packageChart) {
    packageChart.setOption(packageChartOption.value)
  }
}



const updateCategoryStats = () => {
  const data = tableData.value
  const total = data.length
  
  if (total === 0) {
    categoryStats.value.forEach(stat => {
      stat.value = 0
      stat.percentage = 0
    })
    // 更新库存状态统计
    inStockCount.value = 0
    lowStockCount.value = 0
    noStockCount.value = 0
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
  
  // 库存状态统计
  let inStock = 0
  let lowStock = 0
  let noStock = 0
  
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
    
    // 库存状态统计
    if (item.currentStock > 0 && item.currentStock >= 10) {
      inStock++
    } else if (item.currentStock > 0 && item.currentStock < 10) {
      lowStock++
    } else {
      noStock++
    }
  })
  
  // 更新库存状态统计
  inStockCount.value = inStock
  lowStockCount.value = lowStock
  noStockCount.value = noStock
  
  // 更新分类统计数据
  categoryStats.value.forEach((stat, index) => {
    const label = stat.label
    stat.value = categories[label]
    stat.percentage = parseFloat(((categories[label] / total) * 100).toFixed(1))
  })
  
  // 对分类数据进行排序（从高到低）
  categoryStats.value.sort((a, b) => b.value - a.value)
  
  // 更新库存状态图表数据
  const chartData = [
    { value: inStock, name: '有库存', itemStyle: { color: '#67c23a' } },
    { value: lowStock, name: '库存不足', itemStyle: { color: '#e6a23c' } },
    { value: noStock, name: '无库存', itemStyle: { color: '#f56c6c' } }
  ]
  
  const legendData = ['有库存', '库存不足', '无库存']
  
  chartOption.value.series[0].data = chartData
  chartOption.value.legend.data = legendData
  
  // 模拟入库出库记录数据（最近7天）
  const dates = []
  const inStockData = []
  const outStockData = []
  
  for (let i = 6; i >= 0; i--) {
    const date = new Date()
    date.setDate(date.getDate() - i)
    dates.push(`${date.getMonth() + 1}/${date.getDate()}`)
    inStockData.push(Math.floor(Math.random() * 10) + 1)
    outStockData.push(Math.floor(Math.random() * 8) + 1)
  }
  
  stockChartOption.value.xAxis.data = dates
  stockChartOption.value.series[0].data = inStockData
  stockChartOption.value.series[1].data = outStockData
  
  // 更新入库出库记录图表
  updateStockChart()
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

const handleSearch = (event) => {
  // 忽略事件对象，只调用loadData函数
  loadData()
}

onMounted(() => {
  loadData()
  nextTick(() => {
    initChart()
  })
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
  height: 100%;
}

.status-stats {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.status-item {
  display: flex;
  align-items: center;
  gap: 15px;
  padding: 10px;
  background-color: #f5f7fa;
  border-radius: 8px;
}

.status-icon {
  width: 48px;
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.status-info {
  flex: 1;
}

.status-label {
  font-size: 14px;
  color: #606266;
  margin-bottom: 5px;
}

.status-value {
  font-size: 20px;
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
