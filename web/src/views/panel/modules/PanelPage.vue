<template>
  <div class="px-4 pb-4 py-2 mt-4 mb-10 mx-auto flex justify-center items-center h-screen">
    <div class="h-3/5 sm:h-1/2 sm:max-w-sm px-2 my-4">
      <h1 class="text-6xl italic font-bold text-center text-gray-300 mb-5">Ech0s Panel</h1>
      <!-- 返回首页 / 登录 / 注册 -->
      <div class="flex justify-between items-center mb-5">
        <!-- 返回首页 -->
        <div>
          <BaseButton
            @click="$router.push('/')"
            class="text-gray-600 rounded-md !shadow-none !border-none !ring-0 !bg-transparent group"
            title="返回首页"
          >
            <Arrow
              class="w-7 h-7 rotate-180 transition-transform duration-200 group-hover:-translate-x-1"
            />
          </BaseButton>
        </div>
        <!-- 操作按钮 -->
        <div class="flex flex-row items-center gap-2">
          <!-- 状态 / 设置 -->
          <BaseButton
            :icon="ShowStatus ? Setting : Status"
            @click="ShowStatus = !ShowStatus"
            class="text-gray-600 rounded-md w-8 h-8"
            title="状态 / 设置"
          />

          <!-- 退出登录 -->
          <BaseButton
            :icon="Logout"
            @click="handleLogout"
            class="text-gray-600 rounded-md w-8 h-8"
            title="退出登录"
          />

          <!-- 登录 / 注册 -->
          <BaseButton
            :icon="Auth"
            @click="$router.push('/auth')"
            class="text-gray-600 rounded-md w-8 h-8"
            title="登录 / 注册"
          />
        </div>
      </div>
      <!-- TheStatus -->
      <TheStatus v-if="ShowStatus" />
      <!-- TheSetting -->
      <TheSetting v-else />
    </div>
  </div>
</template>

<script setup lang="ts">
import TheStatus from './TheStatus.vue'
import BaseButton from '@/components/common/BaseButton.vue'
import Arrow from '@/components/icons/arrow.vue'
import Auth from '@/components/icons/auth.vue'
import Status from '@/components/icons/status.vue'
import Setting from '@/components/icons/setting.vue'
import TheSetting from './TheSetting.vue'
import Logout from '@/components/icons/logout.vue'
import { ref } from 'vue'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const ShowStatus = ref<boolean>(true)

const handleLogout = () => {
  // 弹出浏览器确认框
  if (confirm('确定要退出登录吗？')) {
    // 清除用户信息
    userStore.logout()
  }
}
</script>
