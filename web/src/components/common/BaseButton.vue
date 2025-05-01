<template>
  <button
    class="cursor-pointer p-1.5 bg-gray-50 shadow-sm ring-inset ring-1 ring-gray-300 text-gray-700"
    :class="customClass"
    :disabled="disabled"
    @click="onClick"
  >
    <span v-if="icon" class="flex items-center justify-center">
      <component :is="icon" class="w-full h-full" />
    </span>
    <span><slot /></span>
  </button>
</template>

<script setup lang="ts">
import type { Component } from 'vue'

const props = defineProps<{
  icon?: Component
  disabled?: boolean
  class?: string // 接收父组件传递的 class
}>()

const emit = defineEmits<{
  (e: 'click', event: MouseEvent): void
}>()

const customClass = props.class

function onClick(event: MouseEvent) {
  if (!props.disabled) emit('click', event)
}
</script>

<style scoped></style>
