<template>
  <el-card class="content-card">
    <!-- 工具栏 -->
    <div class="toolbar">
      <div class="filter-box">
        <el-input
          v-model="localSearchQuery"
          @input="handleSearchInput"
          placeholder="搜索商品名称、编号、品牌"
          clearable
          class="search-input"
        >
          <template #prefix>
            <el-icon><SearchFilled /></el-icon>
          </template>
        </el-input>
        <el-select v-model="filterType" placeholder="库存筛选" clearable style="width: 150px" @change="handleFilterChange">
          <el-option label="全部" value="" />
          <el-option label="有库存" value="hasStock" />
          <el-option label="库存不足" value="lowStock" />
          <el-option label="无库存" value="noStock" />
        </el-select>
      </div>
      <div class="action-box">
        <el-button type="primary" @click="$emit('add')" class="add-btn">
          <el-icon><PlusFilled /></el-icon>
          新增元件
        </el-button>
      </div>
    </div>

    <!-- 数据表格 -->
    <el-table 
      :data="tableData" 
      style="width: 100%" 
      v-loading="loading" 
      stripe
      border
      :cell-style="{ padding: '5px 0' }"
      :header-cell-style="{ padding: '8px 0' }"
      :fit="true"
      @filter-change="handleTableFilterChange"
    >
      <el-table-column type="index" label="序号" width="60" align="center" :index="indexMethod" />
      <el-table-column prop="productCode" column-key="productCode" label="编号" width="120" :resizable="true" :filters="[{ text: 'C开头', value: 'C' }]" :filter-multiple="false" align="center" />
      <el-table-column prop="source" column-key="source" label="渠道" width="100" :resizable="true" :filters="[{ text: '立创商城', value: 'LCSC' }, { text: '淘宝', value: 'TB' }]" :filter-multiple="true" align="center">
        <template #default="{ row }">
          <el-tag :type="row.source === 'LCSC' ? 'success' : 'primary'" size="small">
            {{ row.source === 'LCSC' ? '立创商城' : '淘宝' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="brand" column-key="brand" label="品牌" width="80" :resizable="true" :filters="brandOptions" :filter-multiple="true" align="center" />
      <el-table-column prop="model" column-key="model" label="型号" width="120" :resizable="true" filterable align="center" />
      <el-table-column prop="package" column-key="package" label="封装" width="80" :resizable="true" :filters="packageOptions" :filter-multiple="true" align="center" />
      <el-table-column prop="name" label="商品" show-overflow-tooltip width="350" min-width="100" :resizable="true" align="center" />
      <el-table-column prop="quantity" label="数量" width="80" :resizable="true" align="center" />
      <el-table-column prop="currentStock" label="库存" width="70" align="center">
        <template #default="{ row }">
          <el-tag :type="row.currentStock > 50 ? 'success' : row.currentStock > 0 ? 'warning' : 'danger'" size="small">
            {{ row.currentStock }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="price" label="总价" width="80" :resizable="true" align="center">
        <template #default="{ row }">
          <div class="price-container">
            <el-icon size="14"><Currency /></el-icon>
            <span class="price">{{ row.price.toFixed(2) }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="单价" width="80" :resizable="true" align="center">
        <template #default="{ row }">
          <div class="price-container">
            <el-icon size="14"><Currency /></el-icon>
            <span class="price">{{ (row.quantity > 0 ? row.price / row.quantity : 0).toFixed(2) }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180" fixed="right" align="center">
        <template #default="{ row }">
          <el-button type="primary" size="small" link @click="$emit('edit', row)">
            编辑
          </el-button>
          <el-button type="success" size="small" link @click="$emit('stock-in', row)">
            入库
          </el-button>
          <el-button type="warning" size="small" link @click="$emit('stock-out', row)">
            出库
          </el-button>
          <el-button type="danger" size="small" link @click="$emit('delete', row)">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页 -->
    <div class="pagination-container">
      <div class="pagination-info">
        <span>共 {{ pagination.total }} 条数据</span>
      </div>
      <el-pagination
        :current-page="pagination.page"
        :page-size="pagination.pageSize"
        :page-sizes="[10, 20, 50, 100]"
        :total="pagination.total"
        layout="sizes, prev, pager, next"
        @size-change="handleSizeChange"
        @current-change="handlePageChange"
      />
    </div>
  </el-card>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { getComponents } from '@/api/component'
import { useRoute } from 'vue-router'

const route = useRoute()

const props = defineProps({
  tableData: {
    type: Array,
    default: () => []
  },
  loading: {
    type: Boolean,
    default: false
  },
  searchQuery: {
    type: String,
    default: ''
  },
  pagination: {
    type: Object,
    default: () => ({
      page: 1,
      pageSize: 10,
      total: 0
    })
  }
})

const emit = defineEmits(['search', 'page-change', 'size-change', 'add', 'edit', 'delete', 'stock-in', 'stock-out', 'export', 'import', 'table-filter'])

const localSearchQuery = ref('')
const filterType = ref('')
const packageOptions = ref([])
const brandOptions = ref([])

const loadData = async () => {
  // 由父组件处理数据加载
}

// 提取选项
const extractPackageOptions = (data) => {
  const packages = new Set(data.map(item => item.package).filter(Boolean))
  packageOptions.value = Array.from(packages).sort().map(pkg => ({ text: pkg, value: pkg }))
  
  const brands = new Set(data.map(item => item.brand).filter(Boolean))
  brandOptions.value = Array.from(brands).sort().map(brand => ({ text: brand, value: brand }))
}

const handleSearchInput = (value) => {
  localSearchQuery.value = value
  emit('search', filterType.value, '')
}

const handleFilterChange = () => {
  emit('search', filterType.value)
}

// 处理分页变化
const handlePageChange = (page) => {
  emit('page-change', page)
}

const handleSizeChange = (size) => {
  emit('size-change', size)
}

// 计算序号
const indexMethod = (index) => {
  return (props.pagination.page - 1) * props.pagination.pageSize + index + 1
}

// 处理表格列筛选
const handleTableFilterChange = (filters) => {
  emit('table-filter', filters)
}

// 监听路由参数变化
watch(() => route.query.filter, (newFilter) => {
  if (newFilter) {
    filterType.value = newFilter
    handleFilterChange()
  }
})

// 监听 tableData 变化，自动提取选项
watch(() => props.tableData, (newData) => {
  if (newData && newData.length > 0) {
    extractPackageOptions(newData)
  }
}, { deep: true })

// 暴露方法给父组件
defineExpose({
  loadData,
  extractPackageOptions,
  filterType
})

onMounted(() => {
  loadData()
  // 初始化时如果有数据，提取选项
  if (props.tableData && props.tableData.length > 0) {
    extractPackageOptions(props.tableData)
  }
})

// 监听 tableData 变化，自动提取选项
watch(() => props.tableData, (newData) => {
  if (newData && newData.length > 0) {
    extractPackageOptions(newData)
  }
}, { immediate: true })
</script>

<style scoped>
.content-card {
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  padding: 12px 0;
}

.filter-box {
  display: flex;
  gap: 10px;
  flex: 1;
}

.action-box {
  display: flex;
  gap: 10px;
  margin-left: 10px;
}

.add-btn {
  background-color: #409eff;
  border-color: #409eff;
  color: white;
  font-weight: 600;
  padding: 8px 16px;
  border-radius: 4px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.add-btn:hover {
  background-color: #66b1ff;
  border-color: #66b1ff;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
}

.search-input {
  width: 320px;
}

:deep(.el-input__wrapper) {
  border-radius: 4px;
  box-shadow: 0 0 0 1px #dcdfe6 inset;
}

:deep(.el-input__wrapper:hover) {
  box-shadow: 0 0 0 1px #c0c4cc inset;
}

:deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 1px #409eff inset;
}

:deep(.el-table) {
  font-size: 12px;
  border-radius: 4px;
  overflow: hidden;
  width: 100% !important;
  min-width: 100% !important;
}

:deep(.el-table__header-wrapper),
:deep(.el-table__body-wrapper),
:deep(.el-table__footer-wrapper) {
  width: 100% !important;
}

:deep(.el-table__header),
:deep(.el-table__body),
:deep(.el-table__footer) {
  width: 100% !important;
  min-width: 100% !important;
}

:deep(.el-table__column) {
  min-width: 1px;
}

:deep(.el-table__column-resize-proxy) {
  position: absolute;
  top: 0;
  bottom: 0;
  width: 1px;
  background-color: #409eff;
  cursor: col-resize;
  z-index: 10;
}

:deep(.el-table th) {
  background-color: #f5f7fa !important;
  color: #606266;
  font-weight: 600;
  font-size: 12px;
  padding: 8px 0 !important;
  text-align: center !important;
}

:deep(.el-table--striped .el-table__row--striped td) {
  background: #fafafa;
}

:deep(.el-table td) {
  color: #303133 !important;
  padding: 5px 0 !important;
  text-align: center !important;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

:deep(.el-table th) {
  text-align: center !important;
}

:deep(.el-table__row) {
  height: 45px;
}

:deep(.el-table__row:hover) {
  background-color: #f5f7fa !important;
}

.price-container {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  justify-content: center;
  vertical-align: middle;
}

:deep(.el-table .el-table__cell) {
  text-align: center !important;
  vertical-align: middle !important;
}

.price {
  font-weight: normal;
  color: #303133;
}

/* 列宽调整手柄样式 */
:deep(.el-table__column-resize-proxy) {
  position: absolute;
  top: 0;
  bottom: 0;
  width: 1px;
  background-color: #409EFF;
  z-index: 10;
}

:deep(.el-table__header th) {
  position: relative;
}

.price {
  color: #f56c6c;
  font-weight: 600;
  font-size: 12px;
}

.pagination-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 12px;
  padding: 10px 0 0;
  border-top: 1px solid #ebeef5;
}

.pagination-info {
  color: #909399;
  font-size: 12px;
}

:deep(.el-pagination) {
  padding: 0;
}

:deep(.el-pagination .el-select .el-input) {
  width: 100px;
}

:deep(.el-pagination .el-pager li) {
  min-width: 30px;
  height: 28px;
  line-height: 28px;
  font-size: 12px;
}

:deep(.el-button--small) {
  padding: 4px 6px;
  font-size: 12px;
}

:deep(.el-button--small.is-link) {
  padding: 4px 2px;
  font-size: 12px;
}
</style>
