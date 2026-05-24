<script setup lang="ts">
import { cn } from '@/brand/lib/cn'

defineProps<{ class?: string }>()
</script>

<template>
  <div :class="cn('animate-pulse rounded-md bg-zinc-100 dark:bg-zinc-800', $props.class)" />
</template>
