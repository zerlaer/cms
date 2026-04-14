<template>
  <el-card class="content-card">
    <template #header>
      <div class="card-header">
        <h2>BOM清单智能匹配</h2>
        <p class="subtitle">上传BOM文件，自动匹配库存并支持批量出库</p>
      </div>
    </template>

    <!-- 文件上传区域 -->
    <div class="upload-section">
      <el-upload
        class="upload-demo"
        drag
        :auto-upload="false"
        :on-change="handleFileChange"
        :show-file-list="false"
        accept=".xlsx,.xls"
        :disabled="loading"
      >
        <div class="upload-content">
          <div class="upload-arrow">
            <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M12 19V5"></path>
              <polyline points="5 12 12 5 19 12"></polyline>
            </svg>
          </div>
          <div class="upload-text">
            点击或拖拽文件到此处上传
          </div>
          <div class="upload-tip">
            支持 XLSX、XLS 格式的 BOM 文件，需包含 Comment、Footprint、Quantity 列
          </div>
        </div>
      </el-upload>

      <div v-if="selectedFile" class="file-info">
        <el-tag type="info" effect="dark">{{ selectedFile.name }}</el-tag>
        <el-button type="primary" @click="importBOM" :loading="loading" class="import-btn">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 6px;">
            <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
            <polyline points="7 10 12 15 17 10"></polyline>
            <line x1="12" y1="15" x2="12" y2="3"></line>
          </svg>
          开始匹配
        </el-button>
        <el-button @click="selectedFile = null" class="cancel-btn">
          取消
        </el-button>
      </div>
    </div>

    <!-- 匹配结果 -->
    <div v-if="bomItems.length > 0" class="result-section">
      <div class="result-header">
        <h3>匹配结果</h3>
        <div class="result-stats">
          <el-tag 
            :type="matchedCount > 0 ? 'success' : 'info'" 
            :effect="activeFilter === 'matched' ? 'dark' : 'plain'"
            @click="filterByMatchStatus('matched')"
            class="filter-tag"
          >
            匹配: {{ matchedCount }} 项
          </el-tag>
          <el-tag 
            :type="unmatchedCount > 0 ? 'warning' : 'info'" 
            :effect="activeFilter === 'unmatched' ? 'dark' : 'plain'"
            @click="filterByMatchStatus('unmatched')"
            class="filter-tag"
          >
            未匹配: {{ unmatchedCount }} 项
          </el-tag>
          <el-tag 
            type="info" 
            :effect="activeFilter === 'all' ? 'dark' : 'plain'"
            @click="filterByMatchStatus('all')"
            class="filter-tag"
          >
            全部
          </el-tag>
          <el-button 
            type="info" 
            @click="exportExcel"
            :loading="exportLoading"
            class="export-btn"
          >
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 6px;">
              <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
              <polyline points="7 10 12 15 17 10"></polyline>
              <line x1="12" y1="3" x2="12" y2="15"></line>
            </svg>
            导出Excel
          </el-button>
        </div>
        <el-button 
          v-if="matchedCount > 0" 
          type="primary" 
          @click="handleBatchStockOut" 
          :loading="stockOutLoading"
          class="stock-out-btn"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" style="margin-right: 6px;">
            <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
            <polyline points="7 10 12 15 17 10"></polyline>
            <line x1="12" y1="3" x2="12" y2="15"></line>
          </svg>
          批量出库
        </el-button>
      </div>

      <div class="table-container">
        <el-table 
          :data="filteredBomItems" 
          style="width: 100%" 
          v-loading="loading"
          stripe
          border
          :cell-style="{ padding: '8px 0' }"
          :header-cell-style="{ padding: '10px 0' }"
        >
          <el-table-column type="index" label="序号" width="60" align="center" />
          <el-table-column prop="comment" label="Comment" min-width="150" align="center">
            <template #default="{ row }">
              <span>{{ row.comment }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="footprint" label="Footprint" width="120" align="center" />
          <el-table-column prop="model" label="型号" width="120" align="center" />
          <el-table-column prop="parameters" label="参数/规格" width="150" align="center" />
          <el-table-column prop="quantity" label="数量" width="80" align="center" />
          <el-table-column label="匹配状态" width="100" align="center">
            <template #default="{ row }">
              <el-tag :type="row.matched ? 'success' : 'danger'">
                {{ row.matched ? '已匹配' : '未匹配' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="库存信息" min-width="200" align="center">
            <template #default="{ row }">
              <div v-if="row.matched && row.component">
                <div class="component-info">
                  <div class="component-name">{{ row.component.name }}</div>
                  <div class="component-stock">
                    库存: <el-tag :type="row.component.currentStock > 0 ? 'success' : 'danger'">
                      {{ row.component.currentStock }}
                    </el-tag>
                  </div>
                  <div class="component-model">{{ row.component.model }}</div>
                  <div class="component-package">{{ row.component.package }}</div>
                </div>
              </div>
              <div v-else>
                <el-tag type="info">无匹配库存</el-tag>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="100" align="center">
            <template #default="{ row }">
              <el-button 
                v-if="row.matched && row.component && row.component.currentStock >= row.quantity" 
                type="success" 
                size="small" 
                @click="handleSingleStockOut(row)"
              >
                出库
              </el-button>
              <el-button 
                v-else-if="row.matched && row.component && row.component.currentStock < row.quantity" 
                type="warning" 
                size="small" 
                disabled
              >
库存不足
              </el-button>

            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>

    <!-- 批量出库对话框 -->
    <el-dialog v-model="stockOutDialogVisible" title="批量出库" width="500px">
      <el-form :model="stockOutForm" label-width="80px">
        <el-form-item label="出库备注">
          <el-input v-model="stockOutForm.remark" type="textarea" rows="3" placeholder="请输入出库备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="stockOutDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="confirmBatchStockOut" :loading="stockOutLoading">
            确认出库
          </el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 单个出库对话框 -->
    <el-dialog v-model="singleStockOutDialogVisible" title="出库" width="400px">
      <el-form :model="singleStockOutForm" label-width="80px">
        <el-form-item label="元件名称">
          <el-input v-model="singleStockOutForm.componentName" disabled />
        </el-form-item>
        <el-form-item label="出库数量">
          <el-input-number 
            v-model="singleStockOutForm.quantity" 
            :min="1" 
            :max="singleStockOutForm.maxQuantity"
            :step="1"
          />
        </el-form-item>
        <el-form-item label="出库备注">
          <el-input v-model="singleStockOutForm.remark" type="textarea" rows="2" placeholder="请输入出库备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="singleStockOutDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="confirmSingleStockOut" :loading="stockOutLoading">
            确认出库
          </el-button>
        </span>
      </template>
    </el-dialog>
  </el-card>
</template>

<script setup>
import { ref, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { uploadBOM, batchStockOut, stockOut } from '@/api/bom'
import * as XLSX from 'xlsx'

const selectedFile = ref(null)
const loading = ref(false)
const stockOutLoading = ref(false)
const bomItems = ref([])
const stockOutDialogVisible = ref(false)
const singleStockOutDialogVisible = ref(false)

const stockOutForm = ref({
  remark: 'BOM批量出库'
})

const singleStockOutForm = ref({
  componentId: null,
  componentName: '',
  quantity: 1,
  maxQuantity: 0,
  remark: 'BOM出库'
})

const activeFilter = ref('all')
const exportLoading = ref(false)
const filteredBomItems = computed(() => {
  if (activeFilter.value === 'all') {
    return bomItems.value
  } else if (activeFilter.value === 'matched') {
    return bomItems.value.filter(item => item.matched)
  } else if (activeFilter.value === 'unmatched') {
    return bomItems.value.filter(item => !item.matched)
  }
  return bomItems.value
})

const matchedCount = computed(() => {
  return bomItems.value.filter(item => item.matched).length
})

const unmatchedCount = computed(() => {
  return bomItems.value.filter(item => !item.matched).length
})

const handleFileChange = (file) => {
  selectedFile.value = file.raw
}

const filterByMatchStatus = (status) => {
  activeFilter.value = status
}

const exportExcel = () => {
  if (filteredBomItems.value.length === 0) {
    ElMessage.warning('没有可导出的数据')
    return
  }
  
  exportLoading.value = true
  try {
    // 准备导出数据
    const exportData = filteredBomItems.value.map((item, index) => {
      return {
        序号: index + 1,
        Comment: item.comment || '',
        Footprint: item.footprint || '',
        型号: item.model || '',
        参数规格: item.parameters || '',
        数量: item.quantity || 0,
        匹配状态: item.matched ? '已匹配' : '未匹配',
        元件名称: item.matched && item.component ? item.component.name : '',
        库存数量: item.matched && item.component ? item.component.currentStock : '',
        元件型号: item.matched && item.component ? item.component.model : '',
        元件封装: item.matched && item.component ? item.component.package : ''
      }
    })
    
    // 创建工作簿和工作表
    const wb = XLSX.utils.book_new()
    const ws = XLSX.utils.json_to_sheet(exportData)
    
    // 设置列宽
    const columnWidths = [
      { wch: 8 },  // 序号
      { wch: 20 }, // Comment
      { wch: 12 }, // Footprint
      { wch: 12 }, // 型号
      { wch: 15 }, // 参数规格
      { wch: 8 },  // 数量
      { wch: 10 }, // 匹配状态
      { wch: 20 }, // 元件名称
      { wch: 10 }, // 库存数量
      { wch: 12 }, // 元件型号
      { wch: 12 }  // 元件封装
    ]
    ws['!cols'] = columnWidths
    
    // 将工作表添加到工作簿
    XLSX.utils.book_append_sheet(wb, ws, 'BOM匹配结果')
    
    // 生成当前时间字符串，格式为：YYYY-MM-DD HH-mm
    const now = new Date()
    const year = now.getFullYear()
    const month = String(now.getMonth() + 1).padStart(2, '0')
    const day = String(now.getDate()).padStart(2, '0')
    const hours = String(now.getHours()).padStart(2, '0')
    const minutes = String(now.getMinutes()).padStart(2, '0')
    const timeString = `${year}${month}${day}${hours}${minutes}`
    
    // 生成Excel文件并下载
    XLSX.writeFile(wb, `BOM匹配结果-${timeString}.xlsx`)
    
    ElMessage.success('导出成功')
  } catch (error) {
    console.error('导出失败详情：', error)
    ElMessage.error('导出失败：' + (error.message || '未知错误'))
  } finally {
    exportLoading.value = false
  }
}

const importBOM = async () => {
  if (!selectedFile.value) {
    ElMessage.warning('请选择文件')
    return
  }

  loading.value = true
  try {
    const formData = new FormData()
    formData.append('file', selectedFile.value)
    const response = await uploadBOM(formData)
    bomItems.value = response.data
    ElMessage.success('BOM匹配完成')
  } catch (error) {
    ElMessage.error('BOM匹配失败：' + (error.response?.data?.error || error.message || '未知错误'))
  } finally {
    loading.value = false
  }
}

const handleBatchStockOut = () => {
  stockOutDialogVisible.value = true
}

const confirmBatchStockOut = async () => {
  // 准备批量出库数据
  const batchData = bomItems.value
    .filter(item => item.matched && item.component && item.component.currentStock >= item.quantity)
    .map(item => ({
      componentId: item.component.id,
      quantity: item.quantity,
      remark: stockOutForm.value.remark
    }))

  if (batchData.length === 0) {
    ElMessage.warning('没有可出库的项目')
    return
  }

  stockOutLoading.value = true
  try {
    await batchStockOut(batchData)
    ElMessage.success('批量出库成功')
    stockOutDialogVisible.value = false
    // 刷新匹配结果
    await importBOM()
  } catch (error) {
    ElMessage.error('批量出库失败：' + (error.response?.data?.error || error.message || '未知错误'))
  } finally {
    stockOutLoading.value = false
  }
}

const handleSingleStockOut = (row) => {
  singleStockOutForm.value = {
    componentId: row.component.id,
    componentName: row.component.name,
    quantity: row.quantity,
    maxQuantity: Math.min(row.quantity, row.component.currentStock),
    remark: `BOM出库 - ${row.comment}`
  }
  singleStockOutDialogVisible.value = true
}

const confirmSingleStockOut = async () => {
  stockOutLoading.value = true
  try {
    await stockOut(singleStockOutForm.value.componentId, {
      quantity: singleStockOutForm.value.quantity,
      remark: singleStockOutForm.value.remark
    })
    ElMessage.success('出库成功')
    singleStockOutDialogVisible.value = false
    // 刷新匹配结果
    await importBOM()
  } catch (error) {
    ElMessage.error('出库失败：' + (error.response?.data?.error || error.message || '未知错误'))
  } finally {
    stockOutLoading.value = false
  }
}
</script>

<style scoped>
.content-card {
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  margin: 0 20px;
  max-width: 75vw;
  margin-left: auto;
  margin-right: auto;
  min-height: 60vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 40px 20px;
  box-sizing: border-box;
}

.card-header {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 16px 0;
}

.card-header h2 {
  margin: 3px 0 2% 0;
  font-size: 30px;
  color: #409eff;
}

.subtitle {
  margin: 0;
  color: #909399;
  font-size: 18px;
}

.upload-section {
  padding: 0;
  text-align: center;
  border: 2px dashed #d9d9d9;
  border-radius: 8px;
  margin: 20px auto;
  flex: 1;
  display: flex;
  flex-direction: column;
  width: 100%;
  max-width: none;
  box-sizing: border-box;
  min-height: 400px;
  align-items: center;
  justify-content: center;
}

.upload-section .upload-demo {
  width: 100%;
  max-width: none;
  height: 100%;
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  box-sizing: border-box;
}

.upload-section .upload-demo .el-upload-dragger {
  border: none !important;
  background-color: transparent !important;
  width: 100% !important;
  height: 100% !important;
  margin: 0 !important;
  padding: 0 !important;
  box-shadow: none !important;
  outline: none !important;
  border-radius: 8px !important;
  display: flex;
  align-items: center;
  justify-content: center;
}

.upload-section .upload-demo .el-upload-dragger:hover {
  border: none !important;
  background-color: rgba(64, 158, 255, 0.05) !important;
  box-shadow: none !important;
  outline: none !important;
  border-radius: 8px !important;
}

.upload-section .upload-demo .el-upload-dragger.is-dragover {
  border: none !important;
  background-color: rgba(64, 158, 255, 0.05) !important;
  box-shadow: none !important;
  outline: none !important;
  border-radius: 8px !important;
}

/* 确保Element Plus的上传组件没有任何边框 */
.upload-section .upload-demo .el-upload {
  border: none !important;
  background-color: transparent !important;
  outline: none !important;
  width: 100% !important;
  height: 100% !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

.upload-section .upload-demo .el-upload:focus {
  outline: none !important;
  box-shadow: none !important;
  border: none !important;
}

.upload-section .upload-demo .el-upload--text {
  border: none !important;
  outline: none !important;
  box-shadow: none !important;
}

.upload-section .upload-demo .el-upload--text:focus {
  outline: none !important;
  box-shadow: none !important;
  border: none !important;
}

.upload-section .upload-demo .el-upload__input {
  outline: none !important;
  box-shadow: none !important;
  border: none !important;
}

.upload-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
  width: 100%;
  height: 100%;
  padding: 80px 60px;
  border: none;
  border-radius: 8px;
  transition: all 0.3s ease;
  box-sizing: border-box;
  margin: 0 auto;
}

.upload-content:hover {
  background-color: rgba(64, 158, 255, 0.05);
}

.upload-arrow {
  margin-bottom: 20px;
  color: #409eff;
}

.upload-text {
  font-size: 18px;
  font-weight: 500;
  color: #303133;
  margin-bottom: 12px;
}

.upload-tip {
  font-size: 14px;
  color: #909399;
  margin-top: 12px;
  text-align: center;
  padding: 0 20px;
}

.file-info {
  margin-top: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
}

.import-btn {
  background-color: #409eff;
  border-color: #409eff;
  min-width: 100px;
  height: 32px;
  font-size: 13px;
}

.cancel-btn {
  border-color: #d9d9d9;
  color: #606266;
  min-width: 80px;
  height: 32px;
  font-size: 13px;
}

.export-btn {
  min-width: 100px;
  height: 32px;
  font-size: 13px;
}

.stock-out-btn {
  background-color: #67c23a;
  border-color: #67c23a;
  min-width: 100px;
  height: 32px;
  font-size: 13px;
}

:deep(.el-button) {
  min-width: 80px;
  height: 32px;
  font-size: 13px;
}

:deep(.el-tag) {
  height: 32px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
}

.result-section {
  margin: 20px;
  flex: 1;
}

.result-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
  padding-bottom: 8px;
  border-bottom: 1px solid #ebeef5;
  flex-wrap: wrap;
  gap: 10px;
}

.result-header h3 {
  margin: 0;
  font-size: 18px;
  color: #303133;
}

.result-stats {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
  align-items: center;
}

.filter-tag {
  cursor: pointer;
  transition: all 0.3s ease;
  font-size: 13px;
  padding: 4px 10px;
  height: 32px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
}

.filter-tag:hover {
  transform: translateY(-1px);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.table-container {
  overflow-x: auto;
  -webkit-overflow-scrolling: touch;
  border-radius: 4px;
  border: 1px solid #ebeef5;
}

:deep(.el-table) {
  font-size: 14px;
  border-radius: 4px;
  overflow: hidden;
}

:deep(.el-table th) {
  background-color: #f5f7fa !important;
  color: #606266;
  font-weight: 600;
  font-size: 14px;
  padding: 12px 0 !important;
  text-align: center !important;
}

:deep(.el-table td) {
  color: #303133 !important;
  font-size: 14px;
  padding: 10px 0 !important;
  text-align: center !important;
}

.component-info {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}

.component-name {
  font-weight: 600;
  color: #303133;
  font-size: 14px;
}

.component-stock {
  margin: 4px 0;
  font-size: 14px;
}

.component-model,
.component-package {
  font-size: 12px;
  color: #909399;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .upload-section {
    padding: 20px 10px;
  }

  .result-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 10px;
  }

  .result-stats {
    width: 100%;
    justify-content: space-between;
  }

  .stock-out-btn {
    width: 100%;
  }

  :deep(.el-table) {
    font-size: 13px;
  }

  :deep(.el-table th) {
    font-size: 13px;
    padding: 10px 0 !important;
  }

  :deep(.el-table td) {
    font-size: 13px;
    padding: 8px 0 !important;
  }
}
</style>
