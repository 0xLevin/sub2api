<script setup lang="ts">
import {
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuPortal,
  DropdownMenuRoot,
  DropdownMenuTrigger
} from 'reka-ui'

defineProps<{ items: Array<{ label: string; action: () => void; danger?: boolean }> }>()
</script>

<template>
  <DropdownMenuRoot>
    <DropdownMenuTrigger as-child>
      <slot />
    </DropdownMenuTrigger>
    <DropdownMenuPortal>
      <DropdownMenuContent class="z-50 min-w-40 rounded-md border border-zinc-200 bg-white p-1 shadow-md dark:border-zinc-800 dark:bg-zinc-950">
        <DropdownMenuItem
          v-for="item in items"
          :key="item.label"
          :class="[
            'relative flex cursor-pointer select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none hover:bg-zinc-100 dark:hover:bg-zinc-800',
            item.danger ? 'text-red-600 dark:text-red-400' : 'text-zinc-900 dark:text-zinc-50'
          ]"
          @select="item.action"
        >
          {{ item.label }}
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenuPortal>
  </DropdownMenuRoot>
</template>
