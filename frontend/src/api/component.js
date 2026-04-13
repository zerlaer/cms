import request from '@/utils/request'

export const getComponents = (params) => {
  return request({
    url: '/components',
    method: 'get',
    params
  })
}

export const getComponentById = (id) => {
  return request({
    url: `/components/${id}`,
    method: 'get'
  })
}

export const createComponent = (data) => {
  return request({
    url: '/components',
    method: 'post',
    data
  })
}

export const updateComponent = (id, data) => {
  return request({
    url: `/components/${id}`,
    method: 'put',
    data
  })
}

export const deleteComponent = (id) => {
  return request({
    url: `/components/${id}`,
    method: 'delete'
  })
}

export const stockIn = (id, data) => {
  return request({
    url: `/components/${id}/stock-in`,
    method: 'post',
    data
  })
}

export const stockOut = (id, data) => {
  return request({
    url: `/components/${id}/stock-out`,
    method: 'post',
    data
  })
}

export const getStockRecords = (id) => {
  return request({
    url: `/components/${id}/records`,
    method: 'get'
  })
}

export const exportExcel = () => {
  return request({
    url: '/excel/export',
    method: 'get',
    responseType: 'blob'
  })
}

export const importExcel = (file) => {
  const formData = new FormData()
  formData.append('file', file)
  return request({
    url: '/excel/import',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}
