<script setup lang="ts">
import { computed, ref } from 'vue'
import { RouterLink, useRoute, useRouter } from 'vue-router'
import { BarChart3, CreditCard, Home, KeyRound, LogOut, Menu, User, X } from 'lucide-vue-next'
import { useAppStore, useAuthStore } from '@/stores'
import { Button } from '@/brand/ui'

const route = useRoute()
const router = useRouter()
const appStore = useAppStore()
const authStore = useAuthStore()
const mobileOpen = ref(false)

const siteName = computed(() => appStore.siteName || 'API Gateway')
const siteLogo = computed(() => appStore.siteLogo || '/logo.png')
const user = computed(() => authStore.user)

const navItems = [
  { to: '/app', label: '概览', icon: Home },
  { to: '/app/keys', label: 'API Keys', icon: KeyRound },
  { to: '/app/usage', label: '用量', icon: BarChart3 },
  { to: '/app/billing', label: '购买与订单', icon: CreditCard }
]

function isActive(to: string) {
  return to === '/app' ? route.path === '/app' : route.path.startsWith(to)
}

async function logout() {
  await authStore.logout()
  router.push('/login')
}
</script>

<template>
  <div class="min-h-screen bg-zinc-50 text-zinc-950 dark:bg-zinc-950 dark:text-zinc-50">
    <aside class="fixed inset-y-0 left-0 z-40 hidden w-64 border-r border-zinc-200 bg-white dark:border-zinc-800 dark:bg-zinc-950 lg:block">
      <div class="flex h-16 items-center gap-3 border-b border-zinc-200 px-5 dark:border-zinc-800">
        <img :src="siteLogo" alt="" class="h-8 w-8 rounded-md object-contain">
        <div class="min-w-0">
          <div class="truncate text-sm font-semibold">{{ siteName }}</div>
          <div class="text-xs text-zinc-500">Customer Console</div>
        </div>
      </div>
      <nav class="space-y-1 p-3">
        <RouterLink
          v-for="item in navItems"
          :key="item.to"
          :to="item.to"
          :class="[
            'flex items-center gap-3 rounded-md px-3 py-2 text-sm font-medium',
            isActive(item.to)
              ? 'bg-zinc-950 text-white dark:bg-zinc-50 dark:text-zinc-950'
              : 'text-zinc-600 hover:bg-zinc-100 hover:text-zinc-950 dark:text-zinc-300 dark:hover:bg-zinc-800 dark:hover:text-white'
          ]"
        >
          <component :is="item.icon" class="h-4 w-4" />
          {{ item.label }}
        </RouterLink>
      </nav>
    </aside>

    <div class="lg:pl-64">
      <header class="sticky top-0 z-30 flex h-16 items-center justify-between border-b border-zinc-200 bg-white/90 px-4 backdrop-blur dark:border-zinc-800 dark:bg-zinc-950/90 lg:px-8">
        <button class="rounded-md p-2 lg:hidden" @click="mobileOpen = true">
          <Menu class="h-5 w-5" />
        </button>
        <div class="hidden lg:block">
          <div class="text-sm text-zinc-500">当前账户</div>
          <div class="text-sm font-medium">{{ user?.email || '-' }}</div>
        </div>
        <div class="flex items-center gap-3">
          <div class="hidden rounded-md border border-zinc-200 px-3 py-1.5 text-sm dark:border-zinc-800 sm:block">
            余额 ${{ user?.balance?.toFixed(2) || '0.00' }}
          </div>
          <RouterLink to="/profile">
            <Button variant="outline" size="icon"><User class="h-4 w-4" /></Button>
          </RouterLink>
          <Button variant="ghost" size="icon" @click="logout"><LogOut class="h-4 w-4" /></Button>
        </div>
      </header>

      <main class="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
        <slot />
      </main>
    </div>

    <div v-if="mobileOpen" class="fixed inset-0 z-50 lg:hidden">
      <div class="absolute inset-0 bg-black/50" @click="mobileOpen = false" />
      <div class="absolute inset-y-0 left-0 w-72 bg-white p-4 dark:bg-zinc-950">
        <div class="mb-6 flex items-center justify-between">
          <div class="flex items-center gap-3">
            <img :src="siteLogo" alt="" class="h-8 w-8 rounded-md object-contain">
            <span class="text-sm font-semibold">{{ siteName }}</span>
          </div>
          <button class="rounded-md p-2" @click="mobileOpen = false">
            <X class="h-5 w-5" />
          </button>
        </div>
        <nav class="space-y-1">
          <RouterLink
            v-for="item in navItems"
            :key="item.to"
            :to="item.to"
            class="flex items-center gap-3 rounded-md px-3 py-2 text-sm font-medium"
            @click="mobileOpen = false"
          >
            <component :is="item.icon" class="h-4 w-4" />
            {{ item.label }}
          </RouterLink>
        </nav>
      </div>
    </div>
  </div>
</template>
