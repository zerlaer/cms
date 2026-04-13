<template>
  <div class="inventory-stats">
    <el-row :gutter="20">
      <!-- 统计卡片 -->
      <el-col :span="6">
        <el-card class="stat-card" shadow="hover" @click.native="handleStatClick('all')" style="cursor: pointer">
          <div class="stat-content">
            <div class="stat-icon" style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%)">
              <el-icon :size="32"><Box /></el-icon>
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
              <el-icon :size="32"><CircleCheckFilled /></el-icon>
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
              <el-icon :size="32"><Warning /></el-icon>
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
              <el-icon :size="32"><CircleCloseFilled /></el-icon>
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
              <el-icon><SearchFilled /></el-icon>
            </template>
          </el-input>
        </div>
      </template>

      <el-table :data="tableData" style="width: 100%" v-loading="loading" stripe>
        <el-table-column prop="productCode" label="商品编号" width="120" />
        <el-table-column prop="brand" label="品牌" width="100" />
        <el-table-column prop="name" label="商品名称" min-width="150" />
        <el-table-column prop="currentStock" label="当前库存" width="100">
          <template #default="{ row }">
            <el-tag :type="row.currentStock > 50 ? 'success' : row.currentStock > 0 ? 'warning' : 'danger'">
              {{ row.currentStock }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="quantity" label="订购数量" width="100" />
        <el-table-column prop="price" label="单价" width="100">
          <template #default="{ row }">
            ¥{{ row.price.toFixed(2) }}
          </template>
        </el-table-column>
        <el-table-column prop="totalValue" label="总价值" width="100">
          <template #default="{ row }">
            ¥{{ (row.price * row.currentStock).toFixed(2) }}
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
  padding: 20px;
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
</style>
