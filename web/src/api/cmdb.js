
import service from '@/utils/request'

export const getAssetsList = (params) => {
  return service({
    url: '/cmdb/getAssetsList',
    method: 'get',
    params: params
  })
}

export const getApplyList = (params) => {
  return service({
    url: '/cmdb/getApplyList',
    method: 'get',
    params: params
  })
}

export const createApply = (data) => {
  return service({
    url: '/cmdb/createApply',
    method: 'post',
    data: data,
  })
}

export const deleteApply = (data) => {
  return service({
    url: '/cmdb/deleteApply',
    method: 'delete',
    data: data,
  })
}

export const getTransferList = (params) => {
  return service({
    url: '/cmdb/getTransferList',
    method: 'get',
    params: params
  })
}

export const createTransfer = (data) => {
  return service({
    url: '/cmdb/createTransfer',
    method: 'post',
    data: data,
  })
}

export const deleteTransfer = (data) => {
  return service({
    url: '/cmdb/deleteTransfer',
    method: 'delete',
    data: data,
  })
}

export const getRenewList = (params) => {
  return service({
    url: '/cmdb/getRenewList',
    method: 'get',
    params: params
  })
}

export const createRenew = (data) => {
  return service({
    url: '/cmdb/createRenew',
    method: 'post',
    data: data,
  })
}

export const deleteRenew = (data) => {
  return service({
    url: '/cmdb/deleteRenew',
    method: 'delete',
    data: data,
  })
}
