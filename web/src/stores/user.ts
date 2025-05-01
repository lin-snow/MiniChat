import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { fetchLogin, fetchSignup, fetchGetCurrentUser } from '@/service/api'
import { saveAuthToken } from '@/service/request/shared'
import { localStg } from '@/utils/storage'
import { theToast } from '@/utils/toast'
import { useRouter } from 'vue-router'

export const useUserStore = defineStore('userStore', () => {
  /**
   * state
   */
  const user = ref<App.Api.User.User | null>(null)
  const isLogin = computed(() => !!user.value)
  const router = useRouter()

  /**
   * actions
   */
  async function login(userInfo: App.Api.Auth.LoginParams) {
    await fetchLogin(userInfo).then((res) => {
      const token = String(res.data)

      if (token && token.length > 0) {
        // 保存token到localStorage
        saveAuthToken(token)

        // 获取用户信息
        fetchGetCurrentUser().then((res) => {
          user.value = res.data
        })

        // 登录成功
        theToast.success('登录成功')
        router.push({ name: 'home' })
      }
    })
  }

  async function signup(userInfo: App.Api.Auth.SignupParams) {
    return await fetchSignup(userInfo).then((res) => {
      // 注册成功，前往登录
      if (res.code === 1) {
        theToast.success('注册成功,请登录！')
        return true
      }

      // 注册失败
      return false
    })
  }

  async function logout() {
    // 清除token
    localStg.removeItem('token')
    user.value = null

    // 跳转到首页
    router.push({ name: 'home' })
  }

  async function autoLogin() {
    // 检查localStorage中是否有token
    const token = String(localStg.getItem('token'))
    if (token && token.length > 0 && token !== 'undefined' && token !== 'null') {
      // 如果有token，则获取用户信息
      await fetchGetCurrentUser().then((res) => {
        user.value = res.data
      })
    }
  }

  return { user, isLogin, login, signup, logout, autoLogin }
})
