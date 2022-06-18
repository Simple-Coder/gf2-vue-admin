import request from '@/utils/request'

export function getSysConfigList(params) {
  return request({
    // url: '/vue-admin-template/table/list',
    url: '/api/v1/system/config/list/',
    method: 'get',
    params: params
  })
}
