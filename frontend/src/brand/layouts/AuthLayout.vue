<script setup lang="ts">
import { computed } from 'vue'
import { RouterLink } from 'vue-router'
import { useAppStore } from '@/stores'

const appStore = useAppStore()
const siteName = computed(() => appStore.siteName || 'API Gateway')
const siteLogo = computed(() => appStore.siteLogo || '/logo.png')
</script>

<template>
  <div class="min-h-screen bg-zinc-950 text-white">
    <div class="grid min-h-screen lg:grid-cols-[1.1fr_0.9fr]">
      <section class="hidden bg-[radial-gradient(circle_at_20%_20%,rgba(99,102,241,0.32),transparent_28%),radial-gradient(circle_at_80%_30%,rgba(20,184,166,0.24),transparent_24%),linear-gradient(135deg,#09090b,#18181b)] p-10 lg:flex lg:flex-col lg:justify-between">
        <RouterLink to="/" class="flex items-center gap-3">
          <img :src="siteLogo" alt="" class="h-9 w-9 rounded-md bg-white/10 object-contain">
          <span class="font-semibold">{{ siteName }}</span>
        </RouterLink>
        <div class="max-w-xl">
          <p class="text-sm uppercase tracking-[0.3em] text-zinc-400">Commercial API access</p>
          <h1 class="mt-4 text-5xl font-semibold tracking-tight">一个账户，统一管理 AI API 售卖链路。</h1>
          <p class="mt-5 text-lg text-zinc-300">注册、购买、创建 Key、查看用量，全流程面向终端客户重新设计。</p>
        </div>
      </section>
      <section class="flex items-center justify-center bg-white px-4 py-10 text-zinc-950 dark:bg-zinc-950 dark:text-zinc-50">
        <div class="w-full max-w-md">
          <RouterLink to="/" class="mb-8 flex items-center gap-3 lg:hidden">
            <img :src="siteLogo" alt="" class="h-8 w-8 rounded-md object-contain">
            <span class="font-semibold">{{ siteName }}</span>
          </RouterLink>
          <slot />
        </div>
      </section>
    </div>
  </div>
</template>
