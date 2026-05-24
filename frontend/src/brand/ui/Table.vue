<script setup lang="ts">
import { cn } from '@/brand/lib/cn'

defineProps<{ class?: string }>()
</script>

<template>
  <div class="w-full overflow-auto">
    <table :class="cn('w-full caption-bottom text-sm', $props.class)">
      <slot />
    </table>
  </div>
</template>
