<script setup lang="ts">
import { cn } from '@/brand/lib/cn'

defineProps<{
  modelValue?: string | number
  type?: string
  placeholder?: string
  disabled?: boolean
  class?: string
}>()

defineEmits<{ (e: 'update:modelValue', value: string): void }>()
</script>

<template>
  <input
    :value="modelValue"
    :type="type || 'text'"
    :placeholder="placeholder"
    :disabled="disabled"
    :class="
      cn(
        'flex h-10 w-full rounded-md border border-zinc-200 bg-white px-3 py-2 text-sm text-zinc-950 ring-offset-white',
        'placeholder:text-zinc-500 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-zinc-950 focus-visible:ring-offset-2',
        'disabled:cursor-not-allowed disabled:opacity-50 dark:border-zinc-800 dark:bg-zinc-950 dark:text-zinc-50 dark:ring-offset-zinc-950 dark:placeholder:text-zinc-400 dark:focus-visible:ring-zinc-300',
        $props.class
      )
    "
    @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
  >
</template>
