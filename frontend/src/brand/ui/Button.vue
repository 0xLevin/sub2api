<script setup lang="ts">
import { computed } from 'vue'
import { cn } from '@/brand/lib/cn'

const props = withDefaults(
  defineProps<{
    variant?: 'default' | 'secondary' | 'outline' | 'ghost' | 'destructive'
    size?: 'sm' | 'md' | 'lg' | 'icon'
    type?: 'button' | 'submit' | 'reset'
    disabled?: boolean
    class?: string
  }>(),
  {
    variant: 'default',
    size: 'md',
    type: 'button',
    disabled: false,
    class: ''
  }
)

const classes = computed(() =>
  cn(
    'inline-flex items-center justify-center gap-2 rounded-md text-sm font-medium transition-colors',
    'focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-zinc-950 focus-visible:ring-offset-2',
    'disabled:pointer-events-none disabled:opacity-50 dark:focus-visible:ring-zinc-300',
    {
      'bg-zinc-950 text-white hover:bg-zinc-800 dark:bg-zinc-50 dark:text-zinc-950 dark:hover:bg-zinc-200':
        props.variant === 'default',
      'bg-zinc-100 text-zinc-900 hover:bg-zinc-200 dark:bg-zinc-800 dark:text-zinc-50 dark:hover:bg-zinc-700':
        props.variant === 'secondary',
      'border border-zinc-200 bg-white hover:bg-zinc-100 hover:text-zinc-900 dark:border-zinc-800 dark:bg-zinc-950 dark:hover:bg-zinc-800 dark:hover:text-zinc-50':
        props.variant === 'outline',
      'hover:bg-zinc-100 hover:text-zinc-900 dark:hover:bg-zinc-800 dark:hover:text-zinc-50':
        props.variant === 'ghost',
      'bg-red-600 text-white hover:bg-red-700 dark:bg-red-700 dark:hover:bg-red-800':
        props.variant === 'destructive',
      'h-8 px-3 text-xs': props.size === 'sm',
      'h-10 px-4 py-2': props.size === 'md',
      'h-11 px-8': props.size === 'lg',
      'h-10 w-10': props.size === 'icon'
    },
    props.class
  )
)
</script>

<template>
  <button :type="type" :disabled="disabled" :class="classes">
    <slot />
  </button>
</template>
