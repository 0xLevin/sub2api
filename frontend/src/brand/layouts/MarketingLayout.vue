<script setup lang="ts">
import { computed, ref } from 'vue'
import { RouterLink } from 'vue-router'
import { Menu, X } from 'lucide-vue-next'
import { useAppStore, useAuthStore } from '@/stores'
import { Button } from '@/brand/ui'

const appStore = useAppStore()
const authStore = useAuthStore()
const mobileOpen = ref(false)

const siteName = computed(() => appStore.siteName || 'API Gateway')
const siteLogo = computed(() => appStore.siteLogo || '/logo.png')
const primaryHref = computed(() => (authStore.isAuthenticated ? '/app' : '/register'))
</script>

<template>
  <div class="min-h-screen bg-white text-zinc-950 dark:bg-zinc-950 dark:text-zinc-50">
    <header class="sticky top-0 z-40 border-b border-zinc-200/70 bg-white/90 backdrop-blur dark:border-zinc-800 dark:bg-zinc-950/90">
      <div class="mx-auto flex h-16 max-w-7xl items-center justify-between px-4 sm:px-6 lg:px-8">
        <RouterLink to="/" class="flex items-center gap-3">
          <img :src="siteLogo" alt="" class="h-8 w-8 rounded-md object-contain">
          <span class="text-sm font-semibold tracking-tight">{{ siteName }}</span>
        </RouterLink>
        <nav class="hidden items-center gap-8 text-sm text-zinc-600 dark:text-zinc-300 md:flex">
          <RouterLink to="/pricing" class="hover:text-zinc-950 dark:hover:text-white">价格</RouterLink>
          <RouterLink to="/key-usage" class="hover:text-zinc-950 dark:hover:text-white">用量查询</RouterLink>
          <a v-if="appStore.docUrl" :href="appStore.docUrl" target="_blank" rel="noopener noreferrer" class="hover:text-zinc-950 dark:hover:text-white">文档</a>
        </nav>
        <div class="hidden items-center gap-2 md:flex">
          <RouterLink v-if="!authStore.isAuthenticated" to="/login">
            <Button variant="ghost">登录</Button>
          </RouterLink>
          <RouterLink :to="primaryHref">
            <Button>{{ authStore.isAuthenticated ? '进入控制台' : '开始使用' }}</Button>
          </RouterLink>
        </div>
        <button class="rounded-md p-2 md:hidden" @click="mobileOpen = !mobileOpen">
          <Menu v-if="!mobileOpen" class="h-5 w-5" />
          <X v-else class="h-5 w-5" />
        </button>
      </div>
      <div v-if="mobileOpen" class="border-t border-zinc-200 px-4 py-4 dark:border-zinc-800 md:hidden">
        <div class="flex flex-col gap-3 text-sm">
          <RouterLink to="/pricing" @click="mobileOpen = false">价格</RouterLink>
          <RouterLink to="/key-usage" @click="mobileOpen = false">用量查询</RouterLink>
          <RouterLink to="/login" @click="mobileOpen = false">登录</RouterLink>
          <RouterLink :to="primaryHref" @click="mobileOpen = false">开始使用</RouterLink>
        </div>
      </div>
    </header>
    <main>
      <slot />
    </main>
  </div>
</template>
