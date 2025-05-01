// 封装ofetch

import { ofetch } from 'ofetch'
import { getAuthToken } from './shared'
import { theToast } from '@/utils/toast'

interface RequestOptions {
  url: string
  method: 'GET' | 'POST' | 'PUT' | 'DELETE'
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  data?: any
}

const ofetchInstance = ofetch.create({
  baseURL: import.meta.env.VITE_SERVICE_BASE_URL,
  timeout: 10000,

  // 请求拦截器
  onRequest({ options }) {
    const token = getAuthToken()
    if (token && token.length > 0) {
      options.headers.append('Authorization', token)
    }
  },
  // 响应拦截器
  onResponse({ response }) {
    // 处理响应数据
    if (response.status !== 200) {
      throw new Error(`Request failed with status ${response.status}`)
    }
  },
})

export const request = async <T>(requestOptions: RequestOptions): Promise<App.Api.Response<T>> => {
  // 检查是否使用正向代理
  if (import.meta.env.VITE_PROXY === 'YES') {
    const proxyUrl = import.meta.env.VITE_PROXY_URL
    if (!proxyUrl) {
      throw new Error('Proxy URL is not defined')
    }
    requestOptions.url = `${proxyUrl}${requestOptions.url}`
  }

  // 处理响应数据（code1表示成功，code0表示失败）
  return ofetchInstance<App.Api.Response<T>>(requestOptions.url, {
    method: requestOptions.method,
    body: requestOptions.data,
  }).then((res) => {
    if (res.code !== 1) {
      theToast.error(res.msg || '请求失败')
    }

    return res
  })
}
