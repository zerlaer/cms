<template>
  <div class="layout">
    <!-- 顶部导航栏 -->
    <el-header class="top-header">
      <div class="logo" @click="goToInventory" style="cursor: pointer">
        <div class="logo-icon">
          <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
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
        <span>电子元件仓库管理系统</span>
      </div>
      <div class="user-info">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 8px;">
          <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
          <circle cx="12" cy="7" r="4"></circle>
        </svg>
        <span>{{ username }}</span>
        <el-button type="danger" size="small" @click="handleLogout">
          退出
        </el-button>
      </div>
    </el-header>

    <div class="main-container">
      <!-- 收起按钮 -->
      <div class="toggle-btn-outer" :style="{ left: (isCollapse ? '56px' : '208px') }" @click="toggleSidebar">
        <svg v-if="!isCollapse" xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <line x1="17" y1="10" x2="3" y2="10"></line>
          <line x1="21" y1="6" x2="3" y2="6"></line>
          <line x1="21" y1="14" x2="3" y2="14"></line>
          <line x1="17" y1="18" x2="3" y2="18"></line>
        </svg>
        <svg v-else xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <line x1="21" y1="10" x2="7" y2="10"></line>
          <line x1="21" y1="6" x2="3" y2="6"></line>
          <line x1="21" y1="14" x2="3" y2="14"></line>
          <line x1="21" y1="18" x2="7" y2="18"></line>
        </svg>
      </div>
      <!-- 左侧边栏 -->
      <el-aside class="sidebar" :width="isCollapse ? '64px' : '220px'">
        <el-menu
          :default-active="activeMenu"
          :collapse="isCollapse"
          background-color="#304156"
          text-color="#bfcbd9"
          active-text-color="#409EFF"
          @select="handleMenuSelect"
        >
          <el-sub-menu index="1">
            <template #title>
              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 8px;">
                <path d="M21.21 15.89A10 10 0 1 1 8 2.83"></path>
                <path d="M22 12A10 10 0 0 0 12 2v10z"></path>
              </svg>
              <span>统计报表</span>
            </template>
            <el-menu-item index="1-1">库存统计</el-menu-item>
          </el-sub-menu>
          
          <el-sub-menu index="2">
            <template #title>
              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 8px;">
                <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
                <line x1="3" y1="9" x2="21" y2="9"></line>
                <line x1="9" y1="21" x2="9" y2="9"></line>
              </svg>
              <span>元件管理</span>
            </template>
            <el-menu-item index="2-1">全部元件</el-menu-item>
            <el-menu-item index="2-2">有库存</el-menu-item>
            <el-menu-item index="2-3">库存不足</el-menu-item>
            <el-menu-item index="2-4">无库存</el-menu-item>
          </el-sub-menu>
          
          <el-sub-menu index="3">
            <template #title>
              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 8px;">
                <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
                <polyline points="7 10 12 15 17 10"></polyline>
                <line x1="12" y1="15" x2="12" y2="3"></line>
              </svg>
              <span>导入导出</span>
            </template>
            <el-menu-item index="3-1" @click="importDialogVisible = true">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 6px;">
                <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
                <polyline points="17 8 12 3 7 8"></polyline>
                <line x1="12" y1="3" x2="12" y2="15"></line>
              </svg> 导入 Excel
            </el-menu-item>
            <el-menu-item index="3-2" @click="handleExport">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 6px;">
                <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
                <polyline points="7 10 12 15 17 10"></polyline>
                <line x1="12" y1="15" x2="12" y2="3"></line>
              </svg> 导出 Excel
            </el-menu-item>
          </el-sub-menu>
          
          <el-sub-menu index="4">
            <template #title>
              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 8px;">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                <circle cx="12" cy="7" r="4"></circle>
              </svg>
              <span>用户管理</span>
            </template>
            <el-menu-item index="4-1">用户列表</el-menu-item>
          </el-sub-menu>
          
          <el-menu-item index="5">
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 8px;">
              <circle cx="12" cy="12" r="3"></circle>
              <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"></path>
            </svg>
            <span>系统设置</span>
          </el-menu-item>
        </el-menu>
      </el-aside>

      <!-- 主内容区 -->
      <el-main class="main-content">
        <!-- 元件列表 -->
        <component-list 
          v-if="currentView === 'components'"
          ref="componentListRef"
          :table-data="tableData"
          :loading="loading"
          :search-query="searchQuery"
          :pagination="pagination"
          @search="handleSearch"
          @page-change="handlePageChange"
          @size-change="handleSizeChange"
          @add="handleAdd"
          @edit="handleEdit"
          @delete="handleDelete"
          @stock-in="handleStockIn"
          @stock-out="handleStockOut"
          @export="handleExport"
          @import="importDialogVisible = true"
          @table-filter="handleTableFilter"
        />
        
        <!-- 用户管理 -->
        <user-management v-else-if="currentView === 'users'" />
        
        <!-- 库存统计 -->
        <inventory-stats v-else-if="currentView === 'inventory'" />
        
        <!-- 系统设置 -->
        <system-settings v-else-if="currentView === 'settings'" />
      </el-main>
    </div>

    <!-- 新增/编辑对话框 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
        <el-form-item label="商品编号" prop="productCode">
          <el-input v-model="form.productCode" @blur="autoDetectSource" />
        </el-form-item>
        <el-form-item label="购买渠道" prop="source">
          <el-select v-model="form.source" placeholder="请选择购买渠道" style="width: 100%">
            <el-option label="立创商城" value="LCSC" />
            <el-option label="淘宝" value="TB" />
          </el-select>
        </el-form-item>
        <el-form-item label="品牌" prop="brand">
          <el-input v-model="form.brand" />
        </el-form-item>
        <el-form-item label="厂家型号" prop="model">
          <el-input v-model="form.model" />
        </el-form-item>
        <el-form-item label="封装" prop="package">
          <el-input v-model="form.package" />
        </el-form-item>
        <el-form-item label="商品名称" prop="name">
          <el-input v-model="form.name" />
        </el-form-item>
        <el-form-item label="订购数量" prop="quantity">
          <el-input-number v-model="form.quantity" :min="0" />
        </el-form-item>
        <el-form-item label="商品金额" prop="price">
          <el-input-number v-model="form.price" :min="0" :precision="2" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 入库对话框 -->
    <el-dialog v-model="stockInDialogVisible" title="入库操作" width="500px">
      <el-form :model="stockForm" ref="stockFormRef" label-width="100px">
        <el-form-item label="入库数量">
          <el-input-number v-model="stockForm.quantity" :min="1" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="stockForm.remark" type="textarea" :rows="3" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="stockInDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitStockIn">确定</el-button>
      </template>
    </el-dialog>

    <!-- 出库对话框 -->
    <el-dialog v-model="stockOutDialogVisible" title="出库操作" width="500px">
      <el-form :model="stockForm" ref="stockFormRef" label-width="100px">
        <el-form-item label="出库数量">
          <el-input-number v-model="stockForm.quantity" :min="1" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="stockForm.remark" type="textarea" :rows="3" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="stockOutDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitStockOut">确定</el-button>
      </template>
    </el-dialog>

    <!-- 导入对话框 -->
    <el-dialog v-model="importDialogVisible" title="导入 Excel" width="500px">
      <el-upload
        ref="uploadRef"
        drag
        :auto-upload="false"
        :limit="1"
        :on-change="handleFileChange"
        accept=".xlsx,.xls"
      >
        <el-icon class="el-icon--upload"><upload-filled /></el-icon>
        <div class="el-upload__text">
          将文件拖到此处，或<em>点击上传</em>
        </div>
      </el-upload>
      <template #footer>
        <el-button @click="importDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitImport">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, defineAsyncComponent, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getComponents,
  createComponent,
  updateComponent,
  deleteComponent,
  stockIn,
  stockOut,
  exportExcel,
  importExcel
} from '@/api/component'

// 异步加载组件
const ComponentList = defineAsyncComponent(() => import('./components/ComponentList.vue'))
const UserManagement = defineAsyncComponent(() => import('./UserManagement.vue'))
const InventoryStats = defineAsyncComponent(() => import('./InventoryStats.vue'))
const SystemSettings = defineAsyncComponent(() => import('./Settings.vue'))

const router = useRouter()
const route = useRoute()
const username = ref(localStorage.getItem('username') || 'Admin')
const componentListRef = ref(null)

const activeMenu = ref('1-1')
const currentView = ref('inventory')
const isCollapse = ref(false)

const toggleSidebar = () => {
  isCollapse.value = !isCollapse.value
}

const dialogVisible = ref(false)
const dialogTitle = ref('')
const formRef = ref(null)
const form = ref({
  id: null,
  productCode: '',
  source: '',
  brand: '',
  model: '',
  package: '',
  name: '',
  quantity: 0,
  price: 0
})

const rules = {
  productCode: [{ required: true, message: '请输入商品编号', trigger: 'blur' }],
  source: [{ required: true, message: '请选择购买渠道', trigger: 'change' }],
  brand: [{ required: true, message: '请输入品牌', trigger: 'blur' }],
  model: [{ required: true, message: '请输入厂家型号', trigger: 'blur' }],
  name: [{ required: true, message: '请输入商品名称', trigger: 'blur' }]
}

// 自动检测购买渠道
const autoDetectSource = () => {
  if (form.value.productCode && form.value.productCode.toUpperCase().startsWith('C')) {
    form.value.source = 'LCSC'
  } else if (form.value.productCode) {
    form.value.source = 'TB'
  }
}

const stockInDialogVisible = ref(false)
const stockOutDialogVisible = ref(false)
const stockForm = ref({
  quantity: 1,
  remark: ''
})
const currentComponentId = ref(null)

const importDialogVisible = ref(false)
const uploadRef = ref(null)
const selectedFile = ref(null)

// 元件列表数据
const tableData = ref([])
const loading = ref(false)
const pagination = ref({
  page: 1,
  pageSize: 20, // 默认每页20个
  total: 0
})
const searchQuery = ref('')
const filterType = ref('')
const packageFilter = ref('')
const tableFilters = ref({}) // 表格列筛选条件

// 加载数据
const loadData = async () => {
  loading.value = true
  try {
    // 当有筛选条件时，请求所有数据以便在前端进行筛选
    const res = await getComponents({
      page: 1,
      pageSize: 10000, // 请求足够多的数据
      search: searchQuery.value
    })
    // 后端返回格式：{ data: components, total: total, page: pageSize, pageSize: pageSize }
    let allData = res.data || []
    
    // 清理数据中的空格
    allData = allData.map(item => ({
      ...item,
      source: item.source ? item.source.trim() : '',
      brand: item.brand ? item.brand.trim() : '',
      package: item.package ? item.package.trim() : '',
      model: item.model ? item.model.trim() : ''
    }))
    
    // 根据过滤条件过滤数据
    if (filterType.value) {
      allData = filterData(allData, filterType.value)
    }
    
    // 根据封装筛选
    if (packageFilter.value) {
      allData = allData.filter(item => item.package === packageFilter.value)
    }
    
    // 根据表格列筛选
    if (Object.keys(tableFilters.value).length > 0) {
      Object.entries(tableFilters.value).forEach(([key, values]) => {
        if (Array.isArray(values) && values.length > 0) {
          switch (key) {
            case 'productCode':
              // 商品编号筛选
              allData = allData.filter(item => item.productCode.toUpperCase().startsWith(values[0]))
              break
            case 'source':
            case 'brand':
            case 'package':
              // 下拉菜单筛选
              allData = allData.filter(item => values.includes(item[key]))
              break
            case 'model':
              // 文本筛选
              allData = allData.filter(item => {
                const value = item[key]
                return values.some(val => value && value.toLowerCase().includes(val.toLowerCase()))
              })
              break
          }
        }
      })
    }
    
    // 计算分页数据
    const startIndex = (pagination.value.page - 1) * pagination.value.pageSize
    const endIndex = startIndex + pagination.value.pageSize
    tableData.value = allData.slice(startIndex, endIndex)
    pagination.value.total = allData.length
    
    // 提取选项
    if (componentListRef.value && componentListRef.value.extractPackageOptions) {
      componentListRef.value.extractPackageOptions(allData)
    }
  } catch (error) {
    ElMessage.error('加载数据失败')
  } finally {
    loading.value = false
  }
}

// 数据过滤
const filterData = (data, type) => {
  switch (type) {
    case 'inStock':
      // 有库存的元件
      return data.filter(item => item.currentStock > 0)
    case 'hasStock':
      // 有库存元件（同 inStock）
      return data.filter(item => item.currentStock > 0)
    case 'lowStock':
      // 库存不足（少于 50）
      return data.filter(item => item.currentStock > 0 && item.currentStock < 50)
    case 'noStock':
      // 无库存
      return data.filter(item => item.currentStock === 0)
    default:
      return data
  }
}

// 处理搜索和筛选
const handleSearch = (filter, pkgFilter, search) => {
  if (filter !== undefined) {
    filterType.value = filter
  }
  if (pkgFilter !== undefined) {
    packageFilter.value = pkgFilter
  }
  if (search !== undefined) {
    searchQuery.value = search
  }
  pagination.value.page = 1
  loadData()
}

// 监听路由参数变化
watch(() => route.query.filter, (newFilter) => {
  // 支持所有筛选条件，包括空字符串（显示全部）
  filterType.value = newFilter !== undefined ? newFilter : ''
  // 切换到元件列表视图
  currentView.value = 'components'
  activeMenu.value = '1-1'
  loadData()
})

const handlePageChange = (page) => {
  pagination.value.page = page
  loadData()
}

const handleSizeChange = (size) => {
  pagination.value.pageSize = size
  pagination.value.page = 1
  loadData()
}

// 处理表格列筛选
const handleTableFilter = (filters) => {
  tableFilters.value = filters
  pagination.value.page = 1
  loadData()
}

const handleFileChange = (file) => {
  selectedFile.value = file.raw
}

const handleMenuSelect = (index) => {
  const viewMap = {
    '1-1': 'inventory',
    '2-1': 'components',
    '2-2': 'components',
    '2-3': 'components',
    '2-4': 'components',
    '3-1': 'components',
    '3-2': 'components',
    '4-1': 'users',
    '5': 'settings'
  }
  
  // 菜单对应的筛选条件
  const filterMap = {
    '2-1': '',  // 全部元件
    '2-2': 'hasStock',  // 有库存
    '2-3': 'lowStock',  // 库存不足
    '2-4': 'noStock'  // 无库存
  }
  
  currentView.value = viewMap[index] || 'components'
  activeMenu.value = index
  
  // 设置筛选条件
  if (filterMap[index] !== undefined) {
    filterType.value = filterMap[index]
    loadData()
  } else {
    filterType.value = ''
  }
  
  // 如果是新增元件菜单
  if (index === '1-2') {
    setTimeout(() => handleAdd(), 100)
  }
}

// 跳转到库存统计
const goToInventory = () => {
  currentView.value = 'inventory'
  activeMenu.value = '1-1'
  filterType.value = ''
}

const handleAdd = () => {
  currentView.value = 'components'
  activeMenu.value = '2-1'
  dialogTitle.value = '新增元件'
  form.value = {
    id: null,
    productCode: '',
    source: '',
    brand: '',
    model: '',
    package: '',
    name: '',
    quantity: 0,
    price: 0
  }
  dialogVisible.value = true
}

const handleEdit = (row) => {
  currentView.value = 'components'
  dialogTitle.value = '编辑元件'
  form.value = { ...row }
  dialogVisible.value = true
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (form.value.id) {
          await updateComponent(form.value.id, form.value)
          ElMessage.success('更新成功')
        } else {
          await createComponent(form.value)
          ElMessage.success('创建成功')
        }
        dialogVisible.value = false
        if (componentListRef.value) {
          componentListRef.value.loadData()
        }
      } catch (error) {
      }
    }
  })
}

const handleDelete = (row) => {
  ElMessageBox.confirm('确定要删除该元件吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await deleteComponent(row.id)
      ElMessage.success('删除成功')
      if (componentListRef.value) {
        componentListRef.value.loadData()
      }
    } catch (error) {
    }
  })
}

const handleStockIn = (row) => {
  currentComponentId.value = row.id
  stockForm.value = { quantity: 1, remark: '' }
  stockInDialogVisible.value = true
}

const submitStockIn = async () => {
  try {
    await stockIn(currentComponentId.value, stockForm.value)
    ElMessage.success('入库成功')
    stockInDialogVisible.value = false
    if (componentListRef.value) {
      componentListRef.value.loadData()
    }
  } catch (error) {
    }
}

const handleStockOut = (row) => {
  currentComponentId.value = row.id
  stockForm.value = { quantity: 1, remark: '' }
  stockOutDialogVisible.value = true
}

const submitStockOut = async () => {
  try {
    await stockOut(currentComponentId.value, stockForm.value)
    ElMessage.success('出库成功')
    stockOutDialogVisible.value = false
    if (componentListRef.value) {
      componentListRef.value.loadData()
    }
  } catch (error) {
    }
}

const handleExport = async () => {
  try {
    const res = await exportExcel()
    const blob = new Blob([res], {
      type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
    })
    const url = window.URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    link.download = 'components.xlsx'
    link.click()
    window.URL.revokeObjectURL(url)
    ElMessage.success('导出成功')
  } catch (error) {
    }
}

const submitImport = async () => {
  if (!selectedFile.value) {
    ElMessage.warning('请选择文件')
    return
  }
  
  try {
    await importExcel(selectedFile.value)
    ElMessage.success('导入成功')
    importDialogVisible.value = false
    selectedFile.value = null
    if (uploadRef.value) {
      uploadRef.value.clearFiles()
    }
    if (componentListRef.value) {
      componentListRef.value.loadData()
    }
  } catch (error) {
    ElMessage.error('导入失败：' + (error.response?.data?.error || error.message || '未知错误'))
  }
}

const handleLogout = () => {
  ElMessageBox.confirm('确定要退出登录吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    localStorage.removeItem('isLoggedIn')
    localStorage.removeItem('username')
    ElMessage.success('已退出登录')
    setTimeout(() => {
      router.push('/login')
    }, 500)
  })
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.layout {
  min-height: 100vh;
  background-color: #f0f2f5;
}

.top-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  height: 60px;
}

.logo {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 20px;
  font-weight: bold;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 15px;
}

.user-info span {
  font-size: 14px;
}

.main-container {
  display: flex;
  height: calc(100vh - 60px);
  position: relative;
}

.sidebar {
  background-color: #304156;
  overflow-y: auto;
  box-shadow: 2px 0 8px rgba(0, 0, 0, 0.1);
  transition: width 0.3s ease;
}

.toggle-btn-outer {
  position: absolute;
  top: 6px;
  z-index: 100;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: #304156;
  background-color: transparent;
  border: none;
  transition: all 0.3s ease;
  width: 50px;
  height: 46px;
  line-height: 46px;
}

.toggle-btn-outer:hover {
  color: #409EFF;
}

.sidebar .el-menu {
  border-right: none;
}

.sidebar .el-menu-item {
  height: 50px;
  line-height: 50px;
  color: #bfcbd9 !important;
}

.sidebar .el-menu-item:hover {
  background-color: #263445 !important;
  color: #fff !important;
}

.sidebar .el-menu-item.is-active {
  background-color: #409EFF !important;
  color: #fff !important;
}

.sidebar .el-sub-menu .el-menu-item {
  background-color: #1f2d3d !important;
}

.sidebar .el-sub-menu .el-menu-item:hover {
  background-color: #263445 !important;
  color: #fff !important;
}

.sidebar .el-sub-menu .el-menu-item.is-active {
  background-color: #409EFF !important;
  color: #fff !important;
}

.main-content {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
}

:deep(.el-upload-dragger) {
  padding: 40px 20px;
}

@media (max-width: 768px) {
  .top-header {
    padding: 0 10px;
  }
  
  .logo span {
    display: none;
  }
  
  .user-info span {
    display: none;
  }
  
  .main-container {
    flex-direction: column;
  }
  
  .sidebar {
    width: 100% !important;
  }
}
</style>
