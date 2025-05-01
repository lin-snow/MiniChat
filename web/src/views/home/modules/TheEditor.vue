<template>
  <div
    class="sm:max-w-sm bg-white rounded-lg ring-1 ring-gray-200 ring-inset mx-auto shadow-sm hover:shadow-md"
  >
    <div class="mx-auto w-full px-3 py-4">
      <!-- Title && Nav -->
      <div class="flex justify-between items-center py-1 px-2">
        <div class="flex flex-row items-center gap-2 justify-between">
          <div class="text-xl">👾</div>
          <h2 class="text-slate-600 font-bold italic">{{ appName }}</h2>
        </div>
        <div class="flex flex-row items-center gap-2">
          <!-- RSS -->
          <div>
            <a href="/rss" title="RSS">
              <!-- icon -->
              <Rss class="w-6 h-6 text-gray-400" />
            </a>
          </div>
          <!-- Github -->
          <div>
            <a href="https://github.com/lin-snow/Ech0" target="_blank" title="Github">
              <!-- icon -->
              <Github class="w-6 h-6 text-gray-400" />
            </a>
          </div>
          <!-- PanelPage -->
          <div class="relative">
            <RouterLink to="/panel" title="面板">
              <!-- icon -->
              <Panel class="w-6 h-6 text-gray-400" />
            </RouterLink>
            <!-- <span class="absolute -top-1 -right-1 block w-2 h-2 bg-green-500 rounded-full ring-1 ring-white"></span> -->
          </div>
        </div>
      </div>

      <!-- Editor -->
      <div class="rounded-lg p-2 mb-1">
        <TheMdEditor v-model="echoToAdd.content" class="rounded-lg" />
      </div>

      <!-- Buttons -->
      <div class="flex flex-row items-center justify-between px-2">
        <div class="flex flex-row items-center gap-2">
          <!-- Photo Upload -->
          <div>
            <input
              id="file-input"
              class="hidden"
              type="file"
              accept="image/*"
              ref="fileInput"
              @change="handleUploadImage"
            />
            <BaseButton
              :icon="ImageUpload"
              @click="handTriggerUpload"
              class="w-8 h-8 rounded-md"
              title="上传图片"
            />
          </div>
          <!-- Privacy Set -->
          <div>
            <BaseButton
              :icon="echoToAdd.private ? Private : Public"
              @click="handlePrivate"
              class="w-8 h-8 rounded-md"
              title="是否私密"
            />
          </div>
        </div>

        <div class="flex flex-row items-center gap-2">
          <!-- Clear -->
          <div>
            <BaseButton
              :icon="Clear"
              @click="handleClear"
              class="w-8 h-8 rounded-md"
              title="清空输入和图片"
            />
          </div>
          <!-- Publish -->
          <div>
            <BaseButton
              :icon="Publish"
              @click="handleAddEcho"
              class="w-8 h-8 rounded-md"
              title="发布"
            />
          </div>
        </div>
      </div>

      <!-- Preview Image -->
      <div
        v-if="echoToAdd.image_url"
        class="rounded-lg overflow-hidden shadow-lg w-5/6 mx-auto my-3"
      >
        <a :href="`${apiUrl}${echoToAdd.image_url}`" data-fancybox>
          <img
            :src="`${apiUrl}${echoToAdd.image_url}`"
            alt="Image"
            class="max-w-full object-cover"
            loading="lazy"
          />
        </a>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import Github from '@/components/icons/github.vue'
import Rss from '@/components/icons/rss.vue'
import Panel from '@/components/icons/panel.vue'
import ImageUpload from '@/components/icons/image.vue'
import Public from '@/components/icons/public.vue'
import Private from '@/components/icons/private.vue'
import Clear from '@/components/icons/clear.vue'
import Publish from '@/components/icons/publish.vue'
import BaseButton from '@/components/common/BaseButton.vue'
import TheMdEditor from '@/components/advanced/TheMdEditor.vue'
import { theToast } from '@/utils/toast'
import { Fancybox } from '@fancyapps/ui'
import { onMounted, ref } from 'vue'
import { fetchUploadImage, fetchAddEcho, fetchGetSettings } from '@/service/api'
import { getApiUrl } from '@/service/request/shared'
import { useEchoStore } from '@/stores/echo'
import '@fancyapps/ui/dist/fancybox/fancybox.css'

const appName = ref<string>(import.meta.env.VITE_APP_NAME)
const apiUrl = getApiUrl()
const echoStore = useEchoStore()

const echoToAdd = ref<App.Api.Ech0.EchoToAdd>({
  content: '',
  image_url: null,
  private: false,
})

const fileInput = ref<HTMLInputElement | null>(null)
const handTriggerUpload = () => {
  if (fileInput.value) {
    fileInput.value.click()
  }
}
const handleUploadImage = async (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files && target.files.length > 0) {
    const file = target.files[0]
    fetchUploadImage(file).then((res) => {
      if (res.code === 1) {
        echoToAdd.value.image_url = res.data
        theToast.success('图片上传成功！')
      }
    })
  }
}

const handlePrivate = () => {
  echoToAdd.value.private = !echoToAdd.value.private
}

const handleClear = () => {
  echoToAdd.value.content = ''
  echoToAdd.value.image_url = null
  echoToAdd.value.private = false
}

const handleAddEcho = () => {
  if (!echoToAdd.value.content && !echoToAdd.value.image_url) {
    theToast.error('内容不能为空！')
    return
  }

  fetchAddEcho(echoToAdd.value).then((res) => {
    if (res.code === 1) {
      theToast.success('发布成功！')
      handleClear()
      echoStore.refreshEchos()
    }
  })
}

onMounted(() => {
  Fancybox.bind('[data-fancybox]', {})
  fetchGetSettings().then((res) => {
    if (res.code === 1) {
      appName.value = res.data.server_name
    }
  })
})
</script>
