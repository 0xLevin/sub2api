<script setup lang="ts">
import { TabsContent, TabsList, TabsRoot, TabsTrigger } from 'reka-ui'

defineProps<{
  modelValue: string
  tabs: Array<{ value: string; label: string }>
}>()

defineEmits<{ (e: 'update:modelValue', value: string): void }>()
</script>

<template>
  <TabsRoot :model-value="modelValue" @update:model-value="$emit('update:modelValue', $event)">
    <TabsList class="inline-flex h-10 items-center justify-center rounded-md bg-zinc-100 p-1 text-zinc-500 dark:bg-zinc-800 dark:text-zinc-400">
      <TabsTrigger
        v-for="tab in tabs"
        :key="tab.value"
        :value="tab.value"
        class="inline-flex items-center justify-center whitespace-nowrap rounded-sm px-3 py-1.5 text-sm font-medium transition-all data-[state=active]:bg-white data-[state=active]:text-zinc-950 data-[state=active]:shadow-sm dark:data-[state=active]:bg-zinc-950 dark:data-[state=active]:text-zinc-50"
      >
        {{ tab.label }}
      </TabsTrigger>
    </TabsList>
    <TabsContent v-for="tab in tabs" :key="tab.value" :value="tab.value" class="mt-4">
      <slot :name="tab.value" />
    </TabsContent>
  </TabsRoot>
</template>
