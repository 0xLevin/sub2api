<script setup lang="ts">
import { X } from 'lucide-vue-next'
import { useBrandToast } from '@/brand/composables/useBrandToast'

const { toasts, dismiss } = useBrandToast()
</script>

<template>
  <div class="fixed right-4 top-4 z-[100] flex w-[calc(100vw-2rem)] max-w-sm flex-col gap-2">
    <div
      v-for="toast in toasts"
      :key="toast.id"
      :class="[
        'rounded-lg border p-4 shadow-lg',
        toast.variant === 'destructive'
          ? 'border-red-200 bg-red-50 text-red-950 dark:border-red-900 dark:bg-red-950 dark:text-red-50'
          : 'border-zinc-200 bg-white text-zinc-950 dark:border-zinc-800 dark:bg-zinc-950 dark:text-zinc-50'
      ]"
    >
      <div class="flex items-start justify-between gap-3">
        <div>
          <div class="text-sm font-semibold">{{ toast.title }}</div>
          <div v-if="toast.description" class="mt-1 text-sm opacity-80">
            {{ toast.description }}
          </div>
        </div>
        <button class="rounded-md p-1 opacity-60 hover:opacity-100" @click="dismiss(toast.id)">
          <X class="h-4 w-4" />
        </button>
      </div>
    </div>
  </div>
</template>
