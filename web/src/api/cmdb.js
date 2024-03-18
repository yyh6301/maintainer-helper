
import service from '@/utils/request'

export const getAssetsList = (params) => {
  return service({
    url: '/cmdb/getAssetsList',
    method: 'get',
    params: params
  })
}
