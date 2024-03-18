import service from '@/utils/request'

export const getWorkflowTemplateList = (params) => {
  return service({
    url: '/workflow/getTemplateList',
    method: 'get',
    params: params
  })
}

