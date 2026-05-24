<script setup lang="ts">
import { DialogClose, DialogContent, DialogOverlay, DialogPortal, DialogRoot, DialogTitle } from 'reka-ui'
import { X } from 'lucide-vue-next'
import { cn } from '@/brand/lib/cn'

defineProps<{
  open: boolean
  title?: string
  class?: string
}>()

defineEmits<{ (e: 'update:open', value: boolean): void }>()
</script>

<template>
  <DialogRoot :open="open" @update:open="$emit('update:open', $event)">
    <DialogPortal>
      <DialogOverlay class="fixed inset-0 z-50 bg-black/60" />
      <DialogContent :class="cn('fixed left-1/2 top-1/2 z-50 grid w-[calc(100vw-2rem)] max-w-lg -translate-x-1/2 -translate-y-1/2 gap-4 rounded-lg border border-zinc-200 bg-white p-6 shadow-lg dark:border-zinc-800 dark:bg-zinc-950', $props.class)">
        <DialogTitle v-if="title" class="text-lg font-semibold text-zinc-950 dark:text-zinc-50">
          {{ title }}
        </DialogTitle>
        <slot />
        <DialogClose class="absolute right-4 top-4 rounded-sm opacity-70 transition-opacity hover:opacity-100">
          <X class="h-4 w-4" />
        </DialogClose>
      </DialogContent>
    </DialogPortal>
  </DialogRoot>
</template>
