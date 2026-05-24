<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { BarChart3, Loader2 } from 'lucide-vue-next'
import { usageAPI } from '@/api'
import type { UsageLog, PaginatedResponse } from '@/types'
import type { UserDashboardStats } from '@/api/usage'
import AppShellLayout from '@/brand/layouts/AppShellLayout.vue'
import { formatCurrency, formatDate, formatNumber, formatTokens } from '@/brand/lib/format'
import { Badge, Card, CardSection, Table } from '@/brand/ui'

const loading = ref(true)
const stats = ref<UserDashboardStats | null>(null)
const logs = ref<UsageLog[]>([])

onMounted(async () => {
  try {
    const [statsResult, logsResult] = await Promise.allSettled([
      usageAPI.getDashboardStats(),
      usageAPI.query({ page: 1, page_size: 20, sort_by: 'created_at', sort_order: 'desc' }) as Promise<PaginatedResponse<UsageLog>>
    ])
    if (statsResult.status === 'fulfilled') stats.value = statsResult.value
    if (logsResult.status === 'fulfilled') logs.value = logsResult.value.items
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <AppShellLayout>
    <div class="mb-8">
      <Badge variant="outline">Usage</Badge>
      <h1 class="mt-3 text-3xl font-semibold tracking-tight">用量明细</h1>
      <p class="mt-2 text-zinc-500">查看请求、Token 和成本变化。</p>
    </div>

    <div v-if="loading" class="flex h-48 items-center justify-center">
      <Loader2 class="h-8 w-8 animate-spin text-zinc-500" />
    </div>
    <div v-else class="space-y-6">
      <div class="grid gap-4 md:grid-cols-3">
        <Card>
          <CardSection role="header" class="pb-3"><div class="text-sm text-zinc-500">累计请求</div></CardSection>
          <CardSection><div class="text-3xl font-semibold">{{ formatNumber(stats?.total_requests) }}</div></CardSection>
        </Card>
        <Card>
          <CardSection role="header" class="pb-3"><div class="text-sm text-zinc-500">累计 Token</div></CardSection>
          <CardSection><div class="text-3xl font-semibold">{{ formatTokens(stats?.total_tokens) }}</div></CardSection>
        </Card>
        <Card>
          <CardSection role="header" class="pb-3"><div class="text-sm text-zinc-500">累计成本</div></CardSection>
          <CardSection><div class="text-3xl font-semibold">{{ formatCurrency(stats?.total_actual_cost) }}</div></CardSection>
        </Card>
      </div>

      <Card>
        <CardSection role="header">
          <div class="flex items-center gap-2">
            <BarChart3 class="h-5 w-5" />
            <div>
              <div class="text-lg font-semibold">最近请求</div>
              <div class="text-sm text-zinc-500">最近 20 条用量记录</div>
            </div>
          </div>
        </CardSection>
        <CardSection>
          <Table>
            <thead class="border-b border-zinc-200 text-left text-xs uppercase text-zinc-500 dark:border-zinc-800">
              <tr>
                <th class="px-3 py-3">时间</th>
                <th class="px-3 py-3">模型</th>
                <th class="px-3 py-3">Token</th>
                <th class="px-3 py-3">成本</th>
                <th class="px-3 py-3">耗时</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-zinc-200 dark:divide-zinc-800">
              <tr v-for="log in logs" :key="log.id">
                <td class="px-3 py-4 text-sm text-zinc-500">{{ formatDate(log.created_at) }}</td>
                <td class="px-3 py-4 font-medium">{{ log.model }}</td>
                <td class="px-3 py-4 text-sm">{{ formatTokens(log.input_tokens + log.output_tokens + log.cache_creation_tokens + log.cache_read_tokens) }}</td>
                <td class="px-3 py-4 text-sm">{{ formatCurrency(log.actual_cost) }}</td>
                <td class="px-3 py-4 text-sm text-zinc-500">{{ log.duration_ms }}ms</td>
              </tr>
            </tbody>
          </Table>
          <div v-if="!logs.length" class="py-12 text-center text-sm text-zinc-500">
            暂无用量记录。
          </div>
        </CardSection>
      </Card>
    </div>
  </AppShellLayout>
</template>
