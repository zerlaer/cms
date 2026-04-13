<template>
  <div class="inventory-stats">
    <el-row :gutter="20">
      <!-- 统计卡片 -->
      <el-col :span="6">
        <el-card class="stat-card" shadow="hover" @click.native="handleStatClick('all')" style="cursor: pointer">
          <div class="stat-content">
            <div class="stat-icon" style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%)">
              <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"></path>
                <polyline points="3.27 6.96 12 12.01 20.73 6.96"></polyline>
                <line x1="12" y1="22.08" x2="12" y2="12"></line>
              </svg>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ totalComponents }}</div>
              <div class="stat-label">元件总数</div>
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :span="6">
        <el-card class="stat-card" shadow="hover" @click.native="handleStatClick('hasStock')" style="cursor: pointer">
          <div class="stat-content">
            <div class="stat-icon" style="background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%)">
              <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
                <polyline points="22 4 12 14.01 9 11.01"></polyline>
              </svg>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ inStockCount }}</div>
              <div class="stat-label">有库存元件</div>
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :span="6">
        <el-card class="stat-card" shadow="hover" @click.native="handleStatClick('lowStock')" style="cursor: pointer">
          <div class="stat-content">
            <div class="stat-icon" style="background: linear-gradient(135deg, #fa709a 0%, #fee140 100%)">
              <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"></path>
                <line x1="12" y1="9" x2="12" y2="13"></line>
                <line x1="12" y1="17" x2="12.01" y2="17"></line>
              </svg>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ lowStockCount }}</div>
              <div class="stat-label">库存不足</div>
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :span="6">
        <el-card class="stat-card" shadow="hover" @click.native="handleStatClick('noStock')" style="cursor: pointer">
          <div class="stat-content">
            <div class="stat-icon" style="background: linear-gradient(135deg, #ff9a9e 0%, #fecfef 100%)">
              <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="12" cy="12" r="10"></circle>
                <line x1="15" y1="9" x2="9" y2="15"></line>
                <line x1="9" y1="9" x2="15" y2="15"></line>
              </svg>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ noStockCount }}</div>
              <div class="stat-label">无库存</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 库存分类统计 -->
    <el-row :gutter="20" style="margin-top: 20px">
      <el-col :span="12">
        <el-card>
          <template #header>
            <span>库存分类统计</span>
          </template>
          <div class="chart-container">
            <div class="stat-item" v-for="item in categoryStats" :key="item.label">
              <div class="stat-item-header">
                <span class="stat-item-label">{{ item.label }}</span>
                <span class="stat-item-value">{{ item.value }}</span>
              </div>
              <el-progress :percentage="item.percentage" :color="item.color" />
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :span="12">
        <el-card>
          <template #header>
            <span>最近入库记录</span>
          </template>
          <el-table :data="recentStockIn" style="width: 100%" size="small">
            <el-table-column prop="productCode" label="商品编号" />
            <el-table-column prop="name" label="名称" />
            <el-table-column prop="quantity" label="数量" />
            <el-table-column prop="date" label="日期" />
          </el-table>
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
    </el-card>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { getComponents } from '@/api/component'
import { useRouter } from 'vue-router'

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

// 分类统计
const categoryStats = ref([
  { label: '电阻类', value: 0, percentage: 0, color: '#67c23a' },
  { label: '电容类', value: 0, percentage: 0, color: '#409eff' },
  { label: '集成电路', value: 0, percentage: 0, color: '#e6a23c' },
  { label: '其他', value: 0, percentage: 0, color: '#909399' }
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
  
  // 根据品牌或型号简单分类（实际项目中应该有专门的分类字段）
  const categories = {
    '电阻类': 0,
    '电容类': 0,
    '集成电路': 0,
    '其他': 0
  }
  
  data.forEach(item => {
    const name = item.name.toLowerCase()
    const brand = item.brand.toLowerCase()
    
    if (name.includes('电阻') || brand.includes('res')) {
      categories['电阻类']++
    } else if (name.includes('电容') || brand.includes('cap')) {
      categories['电容类']++
    } else if (name.includes('集成') || name.includes('IC') || name.includes('芯片')) {
      categories['集成电路']++
    } else {
      categories['其他']++
    }
  })
  
  categoryStats.value[0].value = categories['电阻类']
  categoryStats.value[0].percentage = parseFloat(((categories['电阻类'] / total) * 100).toFixed(1))
  
  categoryStats.value[1].value = categories['电容类']
  categoryStats.value[1].percentage = parseFloat(((categories['电容类'] / total) * 100).toFixed(1))
  
  categoryStats.value[2].value = categories['集成电路']
  categoryStats.value[2].percentage = parseFloat(((categories['集成电路'] / total) * 100).toFixed(1))
  
  categoryStats.value[3].value = categories['其他']
  categoryStats.value[3].percentage = parseFloat(((categories['其他'] / total) * 100).toFixed(1))
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
}

.stat-content {
  display: flex;
  align-items: center;
  gap: 20px;
}

.stat-icon {
  width: 80px;
  height: 80px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
}

.stat-info {
  flex: 1;
}

.stat-value {
  font-size: 28px;
  font-weight: bold;
  color: #303133;
  margin-bottom: 5px;
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
}

:deep(.el-card) {
  width: 100%;
}

:deep(.el-card__body) {
  padding: 15px;
}
</style>
