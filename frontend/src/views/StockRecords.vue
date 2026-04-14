<template>
  <div class="stock-records">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>入库出库</span>
          <el-button type="primary" @click="handleRefresh" size="small">
            <el-icon><Refresh /></el-icon> 刷新
          </el-button>
        </div>
      </template>
      
      <el-form :inline="true" :model="searchForm" class="mb-4">
        <el-form-item label="记录类型">
          <el-select v-model="searchForm.type" placeholder="请选择记录类型">
            <el-option label="全部" value=""></el-option>
            <el-option label="入库" value="in"></el-option>
            <el-option label="出库" value="out"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="元件名称">
          <el-input v-model="searchForm.componentName" placeholder="请输入元件名称" />
        </el-form-item>
        <el-form-item label="时间范围">
          <el-date-picker
            v-model="searchForm.dateRange"
            type="daterange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="YYYY-MM-DD"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">
            <el-icon><Search /></el-icon> 查询
          </el-button>
          <el-button @click="resetForm">
            <el-icon><Refresh /></el-icon> 重置
          </el-button>
        </el-form-item>
      </el-form>
      
      <el-table
        v-loading="loading"
        :data="records"
        style="width: 100%"
        border
      >
        <el-table-column prop="id" label="记录ID" width="80" align="center" />
        <el-table-column prop="componentName" label="元件名称" min-width="150" align="center" />
        <el-table-column prop="componentModel" label="元件型号" min-width="120" align="center" />
        <el-table-column prop="type" label="操作类型" width="80" align="center">
          <template #default="scope">
            <el-tag :type="scope.row.type === 'in' ? 'success' : 'warning'">
              {{ scope.row.type === 'in' ? '入库' : '出库' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="quantity" label="数量" width="80" align="center" />
        <el-table-column prop="remark" label="备注" min-width="150" align="center" />
        <el-table-column prop="createdAt" label="操作时间" min-width="180" align="center" />
      </el-table>
      
      <el-pagination
        v-if="total > 0"
        class="mt-4"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
        :page-size="pagination.pageSize"
        :current-page="pagination.page"
        @size-change="handleSizeChange"
        @current-change="handlePageChange"
      />
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { getStockRecords } from '@/api/component'
import { Refresh, Search } from '@element-plus/icons-vue'

const router = useRouter()

const loading = ref(false)
const records = ref([])
const total = ref(0)
const pagination = ref({
  page: 1,
  pageSize: 20
})

const searchForm = ref({
  type: '',
  componentName: '',
  dateRange: []
})

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
    const params = {
      page: pagination.value.page,
      pageSize: pagination.value.pageSize,
      type: searchForm.value.type,
      componentName: searchForm.value.componentName
    }
    
    if (searchForm.value.dateRange && searchForm.value.dateRange.length === 2) {
      params.startDate = searchForm.value.dateRange[0]
      params.endDate = searchForm.value.dateRange[1]
    }
    
    const res = await getStockRecords(params)
    records.value = res.data || []
    total.value = res.total || 0
  } catch (error) {
    // 错误已经在request拦截器中处理
    console.error('Load records error:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.value.page = 1
  loadData()
}

const resetForm = () => {
  searchForm.value = {
    type: '',
    componentName: '',
    dateRange: []
  }
  pagination.value.page = 1
  loadData()
}

const handleRefresh = () => {
  loadData()
}

const handleSizeChange = (size) => {
  pagination.value.pageSize = size
  pagination.value.page = 1
  loadData()
}

const handlePageChange = (page) => {
  pagination.value.page = page
  loadData()
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.stock-records {
  padding: 0 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.mb-4 {
  margin-bottom: 16px;
}

.mt-4 {
  margin-top: 16px;
  display: flex;
  justify-content: center;
}

/* 确保表单居中显示 */
.el-form {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 16px;
}
</style>