<template>
  <div class="layout">
    <!-- 顶部导航栏 -->
    <el-header class="top-header">
      <div class="header-left">
        <!-- 菜单按钮 -->
        <el-button @click="menuDrawerVisible = true" class="menu-btn" style="background: transparent; border: none; padding: 8px; color: white;">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="4" y1="12" x2="20" y2="12"></line>
            <line x1="4" y1="6" x2="20" y2="6"></line>
            <line x1="4" y1="18" x2="20" y2="18"></line>
          </svg>
        </el-button>
        <div class="logo" @click="goToInventory" style="cursor: pointer">
          <span>电子元件管理系统</span>
        </div>
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

    <!-- 菜单抽屉 -->
    <el-drawer
      v-model="menuDrawerVisible"
      title=""
      direction="ltr"
      size="250px"
    >
      <template #title>
        <div style="display: flex; align-items: center; justify-content: center; width: 100%;">
          <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="color: white;">
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
      </template>
      <el-menu
        :default-active="activeMenu"
        background-color="#304156"
        text-color="#bfcbd9"
        active-text-color="#409EFF"
        @select="handleMenuSelect"
      >
          <el-sub-menu index="1">
            <template #title>
              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 8px;">
                <path d="M3 3v18h18"></path>
                <path d="M19 9H5"></path>
                <path d="M16 19H5"></path>
                <path d="M12 14H5"></path>
              </svg>
              <span>统计报表</span>
            </template>
            <el-menu-item index="1-1">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 6px;">
                <path d="M3 3v18h18"></path>
                <path d="M19 9H5"></path>
                <path d="M16 19H5"></path>
                <path d="M12 14H5"></path>
              </svg> 库存统计
            </el-menu-item>
          </el-sub-menu>
          
          <el-sub-menu index="2">
            <template #title>
              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 8px;">
                <rect x="2" y="7" width="20" height="14" rx="2" ry="2"></rect>
                <path d="M16 21V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16"></path>
              </svg>
              <span>元件管理</span>
            </template>
            <el-menu-item index="2-1">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 6px;">
                <rect x="2" y="7" width="20" height="14" rx="2" ry="2"></rect>
                <path d="M16 21V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16"></path>
              </svg> 全部元件
            </el-menu-item>
            <el-menu-item index="2-2">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 6px;">
                <polyline points="22 12 18 12 15 21 9 3 6 12 2 12"></polyline>
              </svg> 有库存
            </el-menu-item>
            <el-menu-item index="2-3">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 6px;">
                <polyline points="22 12 18 12 15 21 9 3 6 12 2 12"></polyline>
              </svg> 库存不足
            </el-menu-item>
            <el-menu-item index="2-4">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 6px;">
                <polyline points="22 12 18 12 15 21 9 3 6 12 2 12"></polyline>
              </svg> 无库存
            </el-menu-item>
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
          
          <el-menu-item index="6" @click="handleBOMMatch">
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 8px;">
              <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
              <polyline points="14 2 14 8 20 8"></polyline>
              <line x1="16" y1="13" x2="8" y2="13"></line>
              <line x1="16" y1="17" x2="8" y2="17"></line>
              <polyline points="10 9 9 9 8 9"></polyline>
            </svg> BOM匹配
          </el-menu-item>
          
          <el-menu-item index="7" @click="handleStockRecords">
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 8px;">
              <path d="M22 12h-4l-3 9L9 3l-3 9H2"></path>
            </svg> 入库出库
          </el-menu-item>
          
          <el-sub-menu index="4">
            <template #title>
              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 8px;">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                <circle cx="12" cy="7" r="4"></circle>
              </svg>
              <span>用户管理</span>
            </template>
            <el-menu-item index="4-1">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 6px;">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                <circle cx="12" cy="7" r="4"></circle>
              </svg> 用户列表
            </el-menu-item>
          </el-sub-menu>
          
          <el-menu-item index="5">
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 8px;">
              <circle cx="12" cy="12" r="3"></circle>
              <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"></path>
            </svg> 系统设置
          </el-menu-item>
        </el-menu>
    </el-drawer>

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
        
        <!-- BOM匹配 -->
        <BOMMatch v-else-if="currentView === 'bom'" />
        
        <!-- 用户管理 -->
        <user-management v-else-if="currentView === 'users'" />
        
        <!-- 库存统计 -->
        <inventory-stats v-else-if="currentView === 'inventory'" />
        
        <!-- 系统设置 -->
        <system-settings v-else-if="currentView === 'settings'" />
        
        <!-- 入库出库记录 -->
        <stock-records v-else-if="currentView === 'stockRecords'" />
      </el-main>

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
import { getComponents, createComponent, updateComponent, deleteComponent, stockIn, stockOut, importExcel } from '@/api/component'
import * as XLSX from 'xlsx'

// 异步加载组件
const ComponentList = defineAsyncComponent(() => import('./components/ComponentList.vue'))
const UserManagement = defineAsyncComponent(() => import('./UserManagement.vue'))
const InventoryStats = defineAsyncComponent(() => import('./InventoryStats.vue'))
const SystemSettings = defineAsyncComponent(() => import('./Settings.vue'))
const BOMMatch = defineAsyncComponent(() => import('./BOMMatch.vue'))
const StockRecords = defineAsyncComponent(() => import('./StockRecords.vue'))

const router = useRouter()
const route = useRoute()
const username = ref(localStorage.getItem('username') || 'Admin')
const componentListRef = ref(null)

const activeMenu = ref('1-1')
const currentView = ref(localStorage.getItem('currentView') || 'inventory')
const menuDrawerVisible = ref(false)



const dialogVisible = ref(false)
const dialogTitle = ref('')
const formRef = ref(null)
const form = ref({
  productCode: '',
  source: '',
  brand: '',
  category: '',
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
  category: [{ required: true, message: '请输入分类', trigger: 'blur' }],
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
  // 检查登录状态
  const isLoggedIn = localStorage.getItem('isLoggedIn')
  if (!isLoggedIn) {
    // 未登录，跳转到登录页面
    router.push('/login')
    return
  }
  
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
    // 错误已经在request拦截器中处理
    console.error('Load data error:', error)
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
  
  // 存储当前视图到localStorage，以便刷新时保持
  localStorage.setItem('currentView', currentView.value)
  
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
    // 加载所有数据用于导出
    loading.value = true
    const res = await getComponents({
      page: 1,
      pageSize: 10000, // 请求足够多的数据
      search: searchQuery.value
    })
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
    
    // 准备导出数据
    const exportData = allData.map((item, index) => {
      return {
        序号: index + 1,
        编号: item.productCode || '',
        渠道: item.source === 'LCSC' ? '立创商城' : item.source === 'TB' ? '淘宝' : '',
        品牌: item.brand || '',
        型号: item.model || '',
        封装: item.package || '',
        商品: item.name || '',
        数量: item.quantity || 0,
        库存: item.currentStock || 0,
        总价: item.price ? item.price.toFixed(2) : '0.00',
        单价: item.quantity > 0 ? (item.price / item.quantity).toFixed(2) : '0.00'
      }
    })
    
    // 创建工作簿和工作表
    const wb = XLSX.utils.book_new()
    const ws = XLSX.utils.json_to_sheet(exportData)
    
    // 设置列宽
    const columnWidths = [
      { wch: 8 },  // 序号
      { wch: 12 }, // 编号
      { wch: 10 }, // 渠道
      { wch: 8 },  // 品牌
      { wch: 12 }, // 型号
      { wch: 8 },  // 封装
      { wch: 30 }, // 商品
      { wch: 8 },  // 数量
      { wch: 7 },  // 库存
      { wch: 8 },  // 总价
      { wch: 8 }   // 单价
    ]
    ws['!cols'] = columnWidths
    
    // 将工作表添加到工作簿
    XLSX.utils.book_append_sheet(wb, ws, '元件清单')
    
    // 生成当前时间字符串，格式为：YYYY-MM-DD HH-mm
    const now = new Date()
    const year = now.getFullYear()
    const month = String(now.getMonth() + 1).padStart(2, '0')
    const day = String(now.getDate()).padStart(2, '0')
    const hours = String(now.getHours()).padStart(2, '0')
    const minutes = String(now.getMinutes()).padStart(2, '0')
    const timeString = `${year}${month}${day}${hours}${minutes}`
    
    // 生成Excel文件并下载
    XLSX.writeFile(wb, `元件清单-${timeString}.xlsx`)
    
    ElMessage.success('导出成功')
  } catch (error) {
    console.error('导出失败详情：', error)
    ElMessage.error('导出失败：' + (error.message || '未知错误'))
  } finally {
    loading.value = false
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
    // 错误已经在request拦截器中处理
    console.error('Import error:', error)
  }
}

// 跳转到BOM匹配页面
const handleBOMMatch = () => {
  currentView.value = 'bom'
  activeMenu.value = '6'
  // 存储当前视图到localStorage，以便刷新时保持
  localStorage.setItem('currentView', currentView.value)
}

// 跳转到入库出库记录页面
const handleStockRecords = () => {
  currentView.value = 'stockRecords'
  activeMenu.value = '7'
  // 存储当前视图到localStorage，以便刷新时保持
  localStorage.setItem('currentView', currentView.value)
}

const handleLogout = () => {
  ElMessageBox.confirm('确定要退出登录吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    localStorage.removeItem('isLoggedIn')
    localStorage.removeItem('username')
    localStorage.removeItem('currentView')
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

.header-left {
  display: flex;
  align-items: center;
  gap: 15px;
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

/* 菜单抽屉样式 */
:deep(.el-drawer__header) {
  background-color: #304156;
  color: white;
  margin: 0;
  padding: 20px;
  border-bottom: 1px solid #404e65;
}

:deep(.el-drawer__title) {
  color: white;
  font-size: 20px;
   text-align: center;
  width: 100%;
  font-weight: bold;
}

:deep(.el-drawer__body) {
  padding: 0;
  background-color: #304156;
}

:deep(.el-menu) {
  border-right: none;
  background-color: #304156;
}

:deep(.el-menu-item) {
  height: 50px;
  line-height: 50px;
  
  color: #bfcbd9 !important;
}

:deep(.el-menu-item svg) {
  margin-left: 45px;
}

:deep(.el-menu-item:hover) {
  background-color: #263445 !important;
  color: #fff !important;
}

:deep(.el-menu-item.is-active) {
  background-color: #409EFF !important;
  color: #fff !important;
}

:deep(.el-sub-menu .el-menu-item) {
  background-color: #1f2d3d !important;
}

:deep(.el-sub-menu .el-menu-item svg) {
  margin-left: 45px;
}

:deep(.el-sub-menu .el-menu-item:hover) {
  background-color: #263445 !important;
  color: #fff !important;
}

:deep(.el-sub-menu .el-menu-item.is-active) {
  background-color: #409EFF !important;
  color: #fff !important;
}

:deep(.el-sub-menu__title) {
  height: 50px;
  line-height: 50px;
}

:deep(.el-sub-menu__title svg) {
  margin-left: 45px;
}

:deep(.el-sub-menu__title:hover) {
  background-color: #263445 !important;
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
  
  .user-info .el-button {
    padding: 4px 8px;
    font-size: 12px;
  }
  
  .main-container {
    flex-direction: column;
  }
  
  .sidebar {
    width: 100% !important;
    height: auto;
    max-height: 200px;
  }
  
  .main-content {
    padding: 10px;
  }
  
  .toggle-btn-outer {
    top: 2px;
    width: 40px;
    height: 40px;
  }
  
  .el-dialog {
    width: 90% !important;
    max-width: 90% !important;
  }
  
  .el-form-item {
    margin-bottom: 12px;
  }
  
  .el-form-item__label {
    font-size: 12px;
    width: 80px;
  }
  
  .el-input,
  .el-select,
  .el-input-number {
    width: 100%;
  }
}
</style>
