import service from '@/utils/request'

export const getWorkflowTemplateList = (params) => {
  return service({
    url: '/workflow/getTemplateList',
    method: 'get',
    params: params
  })
}

export const createWorkflowTemplate = (data) => {
  return service({
    url: '/workflow/createTemplate',
    method: 'post',
    data: data,
  })
}

export const deleteWorkflowTemplate = (data) => {
  return service({
    url: '/workflow/deleteTemplate',
    method: 'delete',
    data: data,
  })
}

export const updateWorkflowTemplate = (data) => {
  return service({
    url: '/workflow/updateTemplate',
    method: 'post',
    data: data,
  })
}

export const getWorkflowTemplateById = (data) => {
  return service({
    url: '/workflow/getTemplateById',
    method: 'post',
    data: data,
  })
}

export const getTemplateStatusList = (params) => {
  return service({
    url: '/workflow/getStatusList',
    method: 'get',
    params: params
  })
}

export const createTemplateStatus = (data) => {
  return service({
    url: '/workflow/createStatus',
    method: 'post',
    data: data,
  })
}

export const deleteTemplateStatus = (data) => {
  return service({
    url: '/workflow/deleteStatus',
    method: 'delete',
    data: data,
  })
}

export const updateTemplateStatus = (data) => {
  return service({
    url: '/workflow/updateStatus',
    method: 'post',
    data: data,
  })
}

export const getTemplateStatusById = (data) => {
  return service({
    url: '/workflow/getStatusById',
    method: 'post',
    data: data,
  })
}

export const getCircleList = (params) => {
  return service({
    url: '/workflow/getCircleList',
    method: 'get',
    params: params
  })
}

export const createCircle = (data) => {
  return service({
    url: '/workflow/createCircle',
    method: 'post',
    data: data,
  })
}

export const deleteCircle = (data) => {
  return service({
    url: '/workflow/deleteCircle',
    method: 'delete',
    data: data,
  })
}

export const updateCircle = (data) => {
  return service({
    url: '/workflow/updateCircle',
    method: 'post',
    data: data,
  })
}

export const getCircleById = (data) => {
  return service({
    url: '/workflow/getCircleById',
    method: 'post',
    data: data,
  })
}

export const getOrderList = (params) => {
  return service({
    url: '/workflow/getOrderList',
    method: 'get',
    params: params
  })
}

export const getOrderDetail = (data) => {
  return service({
    url: '/workflow/getOrderDetailById',
    method: 'post',
    data: data
  })
}

export const createWorkflowOrder = (data) => {
  return service({
    url: '/workflow/createOrder',
    method: 'post',
    data: data,
  })
}
