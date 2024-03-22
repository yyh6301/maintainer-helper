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
