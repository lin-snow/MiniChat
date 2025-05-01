import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { fetchGetEchosByPage } from '@/service/api'

export const useEchoStore = defineStore('echoStore', () => {
  /**
   * state
   */
  const echoList = ref<App.Api.Ech0.Echo[]>([])
  const isLoading = ref<boolean>(true)
  const total = ref<number>(0)
  const pageSize = ref<number>(5)
  const page = ref<number>(0)
  const current = ref<number>(1)
  const hasMore = computed(() => {
    return total.value > echoList.value.length
  })

  /**
   * actions
   */
  async function getEchosByPage() {
    if (current.value <= page.value)
      return

    isLoading.value = true

    await fetchGetEchosByPage({
      page: current.value,
      pageSize: pageSize.value,
    })
      .then((res) => {
        if (res.code === 1) {
          total.value = res.data.total
          echoList.value = [...echoList.value, ...res.data.items]
          page.value += 1
        }
      })
      .finally(() => {
        isLoading.value = false
      })
  }

  const refreshEchos = () => {
    current.value = 1
    page.value = 0
    echoList.value = []
    getEchosByPage()
  }

  return { echoList, isLoading, total, pageSize, page, current, hasMore, getEchosByPage, refreshEchos }
})
