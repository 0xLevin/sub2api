<script setup lang="ts">
import { cn } from '@/brand/lib/cn'

defineProps<{ for?: string; class?: string }>()
</script>

<template>
  <label :for="$props.for" :class="cn('text-sm font-medium leading-none text-zinc-900 peer-disabled:cursor-not-allowed peer-disabled:opacity-70 dark:text-zinc-50', $props.class)">
    <slot />
  </label>
</template>
