<template>
  <div class="sm:max-w-sm mx-auto px-2 my-4">
    <!-- Echos -->
    <div>
      <div v-for="echo in echoStore.echoList" :key="echo.id">
        <TheEchoCard :echo="echo" @refresh="handleRefresh" />
      </div>
    </div>
    <!-- 加载更多 -->
    <div v-if="echoStore.hasMore" class="mb-4 mt-1 -ml-2">
      <BaseButton @click="handleLoadMore" class="rounded-full !hover:bg-orange-100">
        <span class="text-gray-600 text-md text-center px-2 py-1">继续装填</span>
      </BaseButton>
    </div>
    <!-- 没有更多 -->
    <div v-else class="mx-auto my-5 text-center">
      <span class="text-xl text-gray-400">没有啦！🎉</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import TheEchoCard from '@/components/advanced/TheEchoCard.vue'
import { onMounted } from 'vue'
import { useEchoStore } from '@/stores/echo'
import BaseButton from '@/components/common/BaseButton.vue'

const echoStore = useEchoStore()

const handleLoadMore = async () => {
  echoStore.current = echoStore.current + 1
  await echoStore.getEchosByPage()
}

// 刷新数据
const handleRefresh = () => {
  echoStore.refreshEchos()
}

onMounted(async () => {
  // 获取数据
  await echoStore.getEchosByPage()
})
</script>
