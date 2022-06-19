import request from '@/utils/request'
import qs from 'qs'

export function getSysConfigList(params) {
  return request({
    // url: '/vue-admin-template/table/list',
    url: '/api/v1/system/config/list/',
    method: 'get',
    params: params
  })
}
// 查询参数详细
export function getConfig(configId) {
  return request({
    url: '/api/v1/system/config/get/?id=' + configId,
    method: 'get'
  })
}

// 根据参数键名查询参数值
export function getConfigKey(configKey) {
  return request({
    url: '/api/v1/system/config/configKey/' + configKey,
    method: 'get'
  })
}

// 新增参数配置
export function addConfig(data) {
  return request({
    url: '/api/v1/system/config/add',
    method: 'post',
    // transformRequest: [function(data) {
    //   // 对 data 进行任意转换处理
    //   return qs.stringify(data)
    // }],
    data: data
    // data: JSON.stringify(data)
    // data: qs.stringify(data)
    // data
    // data: {
    //   'configName': 123,
    //   'configKey': 12334
    // }
  })
}

// 修改参数配置
export function updateConfig(data) {
  return request({
    url: '/api/v1/system/config/edit',
    method: 'put',
    data: data
  })
}
// 删除参数配置
export function delConfig(configId) {
  return request({
    url: '/api/v1/system/config/delete',
    method: 'delete',
    data: { ids: configId }
  })
}

// 导出参数
export function exportConfig(query) {
  return request({
    url: '/api/v1/system/config/export',
    method: 'get',
    params: query
  })
}
