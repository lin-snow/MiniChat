<script setup lang="ts">
import { RouterView } from 'vue-router'
import { onMounted } from 'vue'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()

onMounted(async () => {
  // 自动登录
  await userStore.autoLogin()
})
</script>

<template>
  <RouterView />
</template>

<style scoped></style>
