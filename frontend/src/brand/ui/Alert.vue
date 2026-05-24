<script setup lang="ts">
import { cn } from '@/brand/lib/cn'

defineProps<{ class?: string }>()
</script>

<template>
  <div :class="cn('rounded-lg border border-zinc-200 bg-white p-4 text-sm text-zinc-950 dark:border-zinc-800 dark:bg-zinc-950 dark:text-zinc-50', $props.class)" role="alert">
    <slot />
  </div>
</template>
