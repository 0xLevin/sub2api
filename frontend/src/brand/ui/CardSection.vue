<script setup lang="ts">
import { computed } from 'vue'
import { cn } from '@/brand/lib/cn'

const props = withDefaults(
  defineProps<{
    role?: 'header' | 'content' | 'footer'
    class?: string
  }>(),
  { role: 'content', class: '' }
)

const classes = computed(() =>
  cn(
    {
      'flex flex-col space-y-1.5 p-6': props.role === 'header',
      'p-6 pt-0': props.role === 'content',
      'flex items-center p-6 pt-0': props.role === 'footer'
    },
    props.class
  )
)
</script>

<template>
  <div :class="classes">
    <slot />
  </div>
</template>
