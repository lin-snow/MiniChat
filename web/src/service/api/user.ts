import { request } from '../request'

// 获取当前登录用户信息
export function fetchGetCurrentUser() {
  return request<App.Api.User.User>({
    url: '/user',
    method: 'GET',
  })
}



