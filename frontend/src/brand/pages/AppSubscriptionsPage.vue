<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { RouterLink } from 'vue-router'
import { CreditCard, Infinity, Loader2, RefreshCw } from 'lucide-vue-next'
import subscriptionsAPI from '@/api/subscriptions'
import type { UserSubscription } from '@/types'
import AppShellLayout from '@/brand/layouts/AppShellLayout.vue'
import { useBrandToast } from '@/brand/composables/useBrandToast'
import { formatCurrency, formatDate } from '@/brand/lib/format'
import { Badge, Button, Card, CardSection } from '@/brand/ui'

const toast = useBrandToast()
const loading = ref(true)
const subscriptions = ref<UserSubscription[]>([])

const activeCount = computed(() => subscriptions.value.filter((item) => item.status === 'active').length)

function statusVariant(status: string) {
  if (status === 'active') return 'success'
  if (status === 'expired') return 'secondary'
  return 'warning'
}

function formatPercent(used?: number, limit?: number | null) {
  if (!limit) return '0%'
  return `${Math.min(((used || 0) / limit) * 100, 100).toFixed(1)}%`
}

function progressWidth(used?: number, limit?: number | null) {
  if (!limit) return '0%'
  return `${Math.min(((used || 0) / limit) * 100, 100)}%`
}

function usageColor(used?: number, limit?: number | null) {
  if (!limit) return 'bg-zinc-300 dark:bg-zinc-700'
  const ratio = (used || 0) / limit
  if (ratio >= 0.9) return 'bg-red-500'
  if (ratio >= 0.7) return 'bg-amber-500'
  return 'bg-emerald-500'
}

function formatLimit(value?: number | null) {
  return value ? formatCurrency(value) : '不限'
}

function daysRemaining(expiresAt?: string | null) {
  if (!expiresAt) return '长期有效'
  const diff = new Date(expiresAt).getTime() - Date.now()
  const days = Math.ceil(diff / 86400000)
  if (days <= 0) return '已过期'
  if (days === 1) return '剩余 1 天'
  return `剩余 ${days} 天`
}

async function load() {
  loading.value = true
  try {
    subscriptions.value = await subscriptionsAPI.getMySubscriptions()
  } catch (err) {
    toast.show({ title: '加载订阅失败', description: err instanceof Error ? err.message : '请稍后重试', variant: 'destructive' })
  } finally {
    loading.value = false
  }
}

onMounted(load)
</script>

<template>
  <AppShellLayout>
    <div class="mb-8 flex flex-col justify-between gap-4 sm:flex-row sm:items-end">
      <div>
        <Badge variant="outline">Subscriptions</Badge>
        <h1 class="mt-3 text-3xl font-semibold tracking-tight">我的订阅</h1>
        <p class="mt-2 text-zinc-500">查看套餐有效期、分组额度与当前用量。</p>
      </div>
      <div class="flex gap-2">
        <Button variant="outline" @click="load"><RefreshCw class="h-4 w-4" />刷新</Button>
        <RouterLink to="/app/billing"><Button><CreditCard class="h-4 w-4" />购买套餐</Button></RouterLink>
      </div>
    </div>

    <div class="mb-6 grid gap-4 md:grid-cols-3">
      <Card>
        <CardSection role="header" class="pb-3"><div class="text-sm text-zinc-500">有效订阅</div></CardSection>
        <CardSection><div class="text-3xl font-semibold">{{ activeCount }}</div></CardSection>
      </Card>
      <Card>
        <CardSection role="header" class="pb-3"><div class="text-sm text-zinc-500">总订阅</div></CardSection>
        <CardSection><div class="text-3xl font-semibold">{{ subscriptions.length }}</div></CardSection>
      </Card>
      <Card>
        <CardSection role="header" class="pb-3"><div class="text-sm text-zinc-500">可用模式</div></CardSection>
        <CardSection><div class="flex items-center gap-2 text-3xl font-semibold"><Infinity class="h-6 w-6 text-zinc-500" />额度制</div></CardSection>
      </Card>
    </div>

    <div v-if="loading" class="flex h-48 items-center justify-center"><Loader2 class="h-8 w-8 animate-spin text-zinc-500" /></div>
    <div v-else-if="!subscriptions.length" class="rounded-lg border border-dashed border-zinc-300 p-12 text-center dark:border-zinc-700">
      <div class="text-lg font-semibold">暂无订阅</div>
      <p class="mt-2 text-sm text-zinc-500">购买套餐后会在这里看到额度和有效期。</p>
      <RouterLink to="/app/billing"><Button class="mt-6">去购买</Button></RouterLink>
    </div>
    <div v-else class="grid gap-6 lg:grid-cols-2">
      <Card v-for="subscription in subscriptions" :key="subscription.id">
        <CardSection role="header">
          <div class="flex items-start justify-between gap-4">
            <div>
              <div class="text-xl font-semibold">{{ subscription.group?.name || `Group #${subscription.group_id}` }}</div>
              <p v-if="subscription.group?.description" class="mt-1 text-sm text-zinc-500">{{ subscription.group.description }}</p>
            </div>
            <Badge :variant="statusVariant(subscription.status)">{{ subscription.status }}</Badge>
          </div>
          <div class="mt-4 grid gap-3 text-sm sm:grid-cols-2">
            <div class="rounded-md bg-zinc-50 p-3 dark:bg-zinc-900">
              <div class="text-zinc-500">有效期</div>
              <div class="mt-1 font-medium">{{ daysRemaining(subscription.expires_at) }}</div>
              <div v-if="subscription.expires_at" class="text-xs text-zinc-500">{{ formatDate(subscription.expires_at) }}</div>
            </div>
            <div class="rounded-md bg-zinc-50 p-3 dark:bg-zinc-900">
              <div class="text-zinc-500">平台</div>
              <div class="mt-1 font-medium">{{ subscription.group?.platform || '-' }}</div>
            </div>
          </div>
        </CardSection>
        <CardSection>
          <div class="space-y-4">
            <div v-for="quota in [
              { label: '日额度', used: subscription.daily_usage_usd, limit: subscription.group?.daily_limit_usd },
              { label: '周额度', used: subscription.weekly_usage_usd, limit: subscription.group?.weekly_limit_usd },
              { label: '月额度', used: subscription.monthly_usage_usd, limit: subscription.group?.monthly_limit_usd }
            ]" :key="quota.label" class="space-y-2">
              <div class="flex items-center justify-between text-sm">
                <span class="font-medium">{{ quota.label }}</span>
                <span class="text-zinc-500">{{ formatCurrency(quota.used || 0) }} / {{ formatLimit(quota.limit) }}</span>
              </div>
              <div class="h-2 overflow-hidden rounded-full bg-zinc-100 dark:bg-zinc-800">
                <div :class="['h-full rounded-full', usageColor(quota.used, quota.limit)]" :style="{ width: progressWidth(quota.used, quota.limit) }" />
              </div>
              <div class="text-xs text-zinc-500">已用 {{ formatPercent(quota.used, quota.limit) }}</div>
            </div>
            <div v-if="!subscription.group?.daily_limit_usd && !subscription.group?.weekly_limit_usd && !subscription.group?.monthly_limit_usd" class="rounded-lg border border-emerald-200 bg-emerald-50 p-4 text-sm text-emerald-800 dark:border-emerald-900 dark:bg-emerald-950 dark:text-emerald-300">
              当前套餐未配置金额额度上限。
            </div>
          </div>
        </CardSection>
        <CardSection role="footer">
          <RouterLink :to="{ path: '/app/billing', query: { group: String(subscription.group_id) } }">
            <Button variant="outline" class="w-full">续费或购买</Button>
          </RouterLink>
        </CardSection>
      </Card>
    </div>
  </AppShellLayout>
</template>
