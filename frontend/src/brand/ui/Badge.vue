<script setup lang="ts">
import { computed } from 'vue'
import { cn } from '@/brand/lib/cn'

const props = withDefaults(
  defineProps<{ variant?: 'default' | 'secondary' | 'outline' | 'success' | 'warning'; class?: string }>(),
  { variant: 'default', class: '' }
)

const classes = computed(() =>
  cn(
    'inline-flex items-center rounded-md px-2.5 py-0.5 text-xs font-semibold transition-colors',
    {
      'bg-zinc-950 text-zinc-50 dark:bg-zinc-50 dark:text-zinc-950': props.variant === 'default',
      'bg-zinc-100 text-zinc-900 dark:bg-zinc-800 dark:text-zinc-50': props.variant === 'secondary',
      'border border-zinc-200 text-zinc-950 dark:border-zinc-800 dark:text-zinc-50': props.variant === 'outline',
      'bg-emerald-100 text-emerald-800 dark:bg-emerald-950 dark:text-emerald-300': props.variant === 'success',
      'bg-amber-100 text-amber-800 dark:bg-amber-950 dark:text-amber-300': props.variant === 'warning'
    },
    props.class
  )
)
</script>

<template>
  <span :class="classes">
    <slot />
  </span>
</template>
