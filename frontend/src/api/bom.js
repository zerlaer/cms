import request from '@/utils/request'

// 上传BOM文件并匹配
export const uploadBOM = (formData) => {
  return request({
    url: '/bom/import',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}

// 批量出库
export const batchStockOut = (batchData) => {
  return request({
    url: '/bom/batch-stock-out',
    method: 'post',
    data: batchData
  })
}

// 单个出库
export const stockOut = (componentId, data) => {
  return request({
    url: `/components/${componentId}/stock-out`,
    method: 'post',
    data: data
  })
}

// 导出BOM匹配结果
export const exportBOM = (data) => {
  return request({
    url: '/bom/export',
    method: 'post',
    data: data,
    responseType: 'blob'
  })
}
