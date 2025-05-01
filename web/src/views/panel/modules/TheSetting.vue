<template>
  <div class="w-full px-2">
    <div class="rounded-md shadow-sm ring-1 ring-gray-200 ring-inset bg-white p-4">
      <!-- 系统设置 -->
      <div>
        <div class="flex flex-row items-center justify-between mb-3">
          <h1 class="text-gray-600 font-bold text-lg">系统设置</h1>
          <div class="flex flex-row items-center justify-end gap-2 w-14">
            <button v-if="editMode" @click="handleUpdateSystemSetting" title="编辑">
              <Saveupdate class="w-5 h-5 text-gray-400 hover:w-6 hover:h-6" />
            </button>
            <button @click="editMode = !editMode" title="编辑">
              <Edit class="w-5 h-5 text-gray-400 hover:w-6 hover:h-6" />
            </button>
          </div>
        </div>
        <div class="flex flex-row items-center justify-start text-gray-500 gap-2 h-14">
          <h2 class="font-semibold w-24">服务名称:</h2>
          <span v-if="!editMode">{{ sysSetting?.server_name }}</span>
          <BaseInput
            v-else
            v-model="sysSetting.server_name"
            type="text"
            placeholder="请输入服务名称"
            class="w-32 !py-1"
          />
        </div>
        <div class="flex flex-row items-center justify-start text-gray-500 gap-2 h-14">
          <h2 class="font-semibold w-24">允许注册:</h2>
          <BaseSwitch
            v-model="sysSetting.allow_register"
            :disabled="!editMode"
            class="w-14"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import BaseInput from '@/components/common/BaseInput.vue'
import Edit from '@/components/icons/edit.vue'
import Saveupdate from '@/components/icons/saveupdate.vue'
import { ref, onMounted } from 'vue'
import { fetchGetSettings, fetchUpdateSettings } from '@/service/api'
import { theToast } from '@/utils/toast'
import BaseSwitch from '@/components/common/BaseSwitch.vue'

const sysSetting = ref<App.Api.Setting.SystemSetting>({
  server_name: '',
  allow_register: true,
})
const editMode = ref<boolean>(false)

const handleUpdateSystemSetting = async () => {
  await fetchUpdateSettings(sysSetting.value)
    .then((res) => {
      if (res.code === 1) {
        theToast.success(res.msg)
      }
    })
    .finally(() => {
      editMode.value = false
      // 重新获取设置
      fetchGetSettings().then((res) => {
        sysSetting.value = res.data
      })
    })
}

onMounted(() => {
  fetchGetSettings().then((res) => {
    sysSetting.value = res.data
  })
})
</script>
