<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { RouterLink } from 'vue-router'
import { ArrowRight, BarChart3, CreditCard, KeyRound, Loader2 } from 'lucide-vue-next'
import { keysAPI, usageAPI } from '@/api'
import subscriptionsAPI, { type SubscriptionSummary } from '@/api/subscriptions'
import type { ApiKey } from '@/types'
import type { UserDashboardStats } from '@/api/usage'
import AppShellLayout from '@/brand/layouts/AppShellLayout.vue'
import { formatCurrency, formatNumber, formatTokens } from '@/brand/lib/format'
import { Button, Card, CardSection, Badge } from '@/brand/ui'

const loading = ref(true)
const stats = ref<UserDashboardStats | null>(null)
const keys = ref<ApiKey[]>([])
const subscriptionSummary = ref<SubscriptionSummary | null>(null)

onMounted(async () => {
  try {
    const [statsResult, keyResult, subscriptionResult] = await Promise.allSettled([
      usageAPI.getDashboardStats(),
      keysAPI.list(1, 5),
      subscriptionsAPI.getSubscriptionSummary()
    ])
    if (statsResult.status === 'fulfilled') stats.value = statsResult.value
    if (keyResult.status === 'fulfilled') keys.value = keyResult.value.items
    if (subscriptionResult.status === 'fulfilled') subscriptionSummary.value = subscriptionResult.value
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <AppShellLayout>
    <div class="mb-8 flex flex-col justify-between gap-4 sm:flex-row sm:items-end">
      <div>
        <Badge variant="outline">Dashboard</Badge>
        <h1 class="mt-3 text-3xl font-semibold tracking-tight">控制台概览</h1>
        <p class="mt-2 text-zinc-500">查看账户产出、Key 状态和套餐使用情况。</p>
      </div>
      <RouterLink to="/app/billing">
        <Button>购买套餐 <ArrowRight class="h-4 w-4" /></Button>
      </RouterLink>
    </div>

    <div v-if="loading" class="flex h-48 items-center justify-center">
      <Loader2 class="h-8 w-8 animate-spin text-zinc-500" />
    </div>
    <div v-else class="space-y-6">
      <div class="grid gap-4 md:grid-cols-2 xl:grid-cols-4">
        <Card>
          <CardSection role="header" class="pb-3"><div class="text-sm text-zinc-500">今日请求</div></CardSection>
          <CardSection><div class="text-3xl font-semibold">{{ formatNumber(stats?.today_requests) }}</div></CardSection>
        </Card>
        <Card>
          <CardSection role="header" class="pb-3"><div class="text-sm text-zinc-500">今日 Token</div></CardSection>
          <CardSection><div class="text-3xl font-semibold">{{ formatTokens(stats?.today_tokens) }}</div></CardSection>
        </Card>
        <Card>
          <CardSection role="header" class="pb-3"><div class="text-sm text-zinc-500">今日成本</div></CardSection>
          <CardSection><div class="text-3xl font-semibold">{{ formatCurrency(stats?.today_actual_cost) }}</div></CardSection>
        </Card>
        <Card>
          <CardSection role="header" class="pb-3"><div class="text-sm text-zinc-500">有效套餐</div></CardSection>
          <CardSection><div class="text-3xl font-semibold">{{ subscriptionSummary?.active_count || 0 }}</div></CardSection>
        </Card>
      </div>

      <div class="grid gap-6 lg:grid-cols-[1.1fr_0.9fr]">
        <Card>
          <CardSection role="header">
            <div class="flex items-center justify-between">
              <div>
                <div class="text-lg font-semibold">API Keys</div>
                <div class="text-sm text-zinc-500">最近创建的访问密钥</div>
              </div>
              <RouterLink to="/app/keys"><Button variant="outline" size="sm">管理</Button></RouterLink>
            </div>
          </CardSection>
          <CardSection>
            <div class="space-y-3">
              <div v-for="key in keys" :key="key.id" class="flex items-center justify-between rounded-lg border border-zinc-200 p-3 dark:border-zinc-800">
                <div class="flex items-center gap-3">
                  <KeyRound class="h-4 w-4 text-zinc-500" />
                  <div>
                    <div class="font-medium">{{ key.name }}</div>
                    <div class="font-mono text-xs text-zinc-500">{{ key.key }}</div>
                  </div>
                </div>
                <Badge :variant="key.status === 'active' ? 'success' : 'secondary'">{{ key.status }}</Badge>
              </div>
              <div v-if="!keys.length" class="rounded-lg border border-dashed border-zinc-300 p-6 text-center text-sm text-zinc-500 dark:border-zinc-700">
                还没有 API Key。
              </div>
            </div>
          </CardSection>
        </Card>

        <Card>
          <CardSection role="header">
            <div class="text-lg font-semibold">下一步</div>
            <div class="text-sm text-zinc-500">完成售卖链路里的关键动作</div>
          </CardSection>
          <CardSection>
            <div class="grid gap-3">
              <RouterLink to="/app/keys" class="flex items-center gap-3 rounded-lg border border-zinc-200 p-4 hover:bg-zinc-50 dark:border-zinc-800 dark:hover:bg-zinc-900">
                <KeyRound class="h-5 w-5" />
                <span class="font-medium">创建或复制 API Key</span>
              </RouterLink>
              <RouterLink to="/app/usage" class="flex items-center gap-3 rounded-lg border border-zinc-200 p-4 hover:bg-zinc-50 dark:border-zinc-800 dark:hover:bg-zinc-900">
                <BarChart3 class="h-5 w-5" />
                <span class="font-medium">查看用量明细</span>
              </RouterLink>
              <RouterLink to="/app/billing" class="flex items-center gap-3 rounded-lg border border-zinc-200 p-4 hover:bg-zinc-50 dark:border-zinc-800 dark:hover:bg-zinc-900">
                <CreditCard class="h-5 w-5" />
                <span class="font-medium">购买套餐或查看订单</span>
              </RouterLink>
            </div>
          </CardSection>
        </Card>
      </div>
    </div>
  </AppShellLayout>
</template>
