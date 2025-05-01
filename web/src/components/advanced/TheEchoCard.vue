<template>
  <div>
    <!-- 日期时间 && 操作按钮 -->
    <div class="flex justify-between items-center">
      <!-- 日期时间 -->
      <div class="flex justify-start items-center h-auto">
        <!-- 小点 -->
        <div class="w-2 h-2 rounded-full bg-orange-600 mr-2"></div>
        <!-- 具体日期时间 -->
        <div class="flex justify-start text-sm text-orange-500">
          {{ formatDate(props.echo.created_at) }}
        </div>
      </div>
      <!-- 操作按钮 -->
      <div class="flex justify-end items-center gap-1 h-auto">
        <!-- 是否隐私 -->
        <span v-if="props.echo.private" title="私密状态"><Lock /></span>
        <!-- 删除 -->
        <button v-if="userStore.isLogin" @click="handleDeleteEcho(props.echo.id)" title="删除">
          <Roll />
        </button>
      </div>
    </div>

    <!-- 图片 && 内容 -->
    <div class="border-l-2 border-gray-300 p-6 ml-1">
      <!-- 图片 -->
      <div
        v-if="props.echo.image_url"
        class="rounded-lg overflow-hidden shadow-lg w-5/6 mx-auto mb-2"
      >
        <a :href="`${API_URL}${props.echo.image_url}`" data-fancybox>
          <img
            :src="`${API_URL}${props.echo.image_url}`"
            alt="Image"
            class="max-w-full object-cover"
            loading="lazy"
          />
        </a>
      </div>
      <!-- 内容 -->
      <div>
        <MdPreview
          :id="previewOptions.proviewId"
          :modelValue="props.echo.content"
          :theme="previewOptions.theme"
          :show-code-row-number="previewOptions.showCodeRowNumber"
          :preview-theme="previewOptions.previewTheme"
          :code-theme="previewOptions.codeTheme"
          :code-style-reverse="previewOptions.codeStyleReverse"
          :no-img-zoom-in="previewOptions.noImgZoomIn"
          :code-foldable="previewOptions.codeFoldable"
          :auto-fold-threshold="previewOptions.autoFoldThreshold"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Fancybox } from '@fancyapps/ui'
import { MdPreview } from 'md-editor-v3'
import { getApiUrl } from '@/service/request/shared'
import { onMounted } from 'vue'
import { fetchDeleteEcho } from '@/service/api'
import { theToast } from '@/utils/toast'
import { useUserStore } from '@/stores/user'
import '@fancyapps/ui/dist/fancybox/fancybox.css'
import 'md-editor-v3/lib/preview.css'
import Roll from '../icons/roll.vue'
import Lock from '../icons/lock.vue'

const emit = defineEmits(['refresh'])

type Echo = App.Api.Ech0.Echo

const props = defineProps<{ echo: Echo }>()
const API_URL = getApiUrl()
const userStore = useUserStore()
const previewOptions = {
  proviewId: 'preview-only',
  theme: 'light' as 'light' | 'dark',
  showCodeRowNumber: false,
  previewTheme: 'github',
  codeTheme: 'atom',
  codeStyleReverse: true,
  noImgZoomIn: false,
  codeFoldable: true,
  autoFoldThreshold: 15,
}

const handleDeleteEcho = (echoId: number) => {
  // 浏览器alert弹窗确认删除
  if (confirm('确定要删除吗？')) {
    fetchDeleteEcho(echoId).then(() => {
      theToast.success('删除成功！')
      // 触发父组件的刷新事件emit
      emit('refresh')
    })
  }
}

const formatDate = (dateString: string) => {
  // 当天则显示（时：分）
  // 非当天但是三内天则显示几天前
  // 超过三天则显示（时：分 年月日）
  const date = new Date(dateString)
  const now = new Date()
  const diff = now.getTime() - date.getTime()
  const diffInDays = Math.floor(diff / (1000 * 60 * 60 * 24))
  const diffInHours = Math.floor(diff / (1000 * 60 * 60))
  const diffInMinutes = Math.floor(diff / (1000 * 60))

  const diffInSeconds = Math.floor(diff / 1000)
  if (diffInSeconds < 60) {
    return '刚刚'
  } else if (diffInMinutes < 60) {
    return `${diffInMinutes}分钟前`
  } else if (diffInHours < 24) {
    return `${diffInHours}小时前`
  } else if (diffInDays < 3) {
    return `${diffInDays}天前`
  } else {
    return date.toLocaleString() // 返回完整的日期和时间
  }
}

onMounted(() => {
  Fancybox.bind('[data-fancybox]', {
    // Your custom options
  })
})
</script>

<style scoped lang="css">
#preview-only {
  background-color: inherit;
}

.md-editor {
  font-family: 'LXGW WenKai Screen';
}

:deep(ul li) {
  list-style-type: disc;
}
:deep(ul li li) {
  list-style-type: circle;
}
:deep(ul li li li) {
  list-style-type: square;
}
:deep(ol li) {
  list-style-type: decimal;
}
</style>
