<script setup lang="ts">
import { computed, onBeforeUnmount, onMounted, ref } from 'vue'
import { Activity, Gauge, Loader2, RefreshCw, Signal } from 'lucide-vue-next'
import { list as listChannelMonitors, status as getChannelMonitorStatus, type UserMonitorDetail, type UserMonitorView } from '@/api/channelMonitor'
import AppShellLayout from '@/brand/layouts/AppShellLayout.vue'
import { useBrandToast } from '@/brand/composables/useBrandToast'
import { Badge, Button, Card, CardSection, Dialog, Table } from '@/brand/ui'

const toast = useBrandToast()
const loading = ref(true)
const refreshing = ref(false)
const items = ref<UserMonitorView[]>([])
const selected = ref<UserMonitorView | null>(null)
const detail = ref<UserMonitorDetail | null>(null)
const detailLoading = ref(false)
const detailOpen = ref(false)
let abortController: AbortController | null = null

const healthyCount = computed(() => items.value.filter((item) => item.primary_status === 'operational').length)
const degradedCount = computed(() => Math.max(items.value.length - healthyCount.value, 0))
const avgAvailability = computed(() => {
  if (!items.value.length) return 0
  const total = items.value.reduce((sum, item) => sum + (item.availability_7d || 0), 0)
  return total / items.value.length
})

function statusVariant(status: string) {
  return status === 'operational' ? 'success' : status === 'unknown' || !status ? 'secondary' : 'warning'
}

function formatPercent(value?: number | null) {
  if (value == null || Number.isNaN(value)) return '-'
  return `${value.toFixed(2)}%`
}

function formatLatency(value?: number | null) {
  if (value == null) return '-'
  return `${Math.round(value)} ms`
}

async function load(silent = false) {
  if (abortController) abortController.abort()
  const ctrl = new AbortController()
  abortController = ctrl
  loading.value = !silent
  refreshing.value = silent
  try {
    const result = await listChannelMonitors({ signal: ctrl.signal })
    if (ctrl.signal.aborted) return
    items.value = result.items || []
  } catch (err) {
    const code = typeof err === 'object' && err && 'code' in err ? String(err.code) : ''
    if (code !== 'ERR_CANCELED') {
      toast.show({ title: '加载渠道状态失败', description: err instanceof Error ? err.message : '请稍后重试', variant: 'destructive' })
    }
  } finally {
    if (abortController === ctrl) {
      loading.value = false
      refreshing.value = false
      abortController = null
    }
  }
}

async function openDetail(item: UserMonitorView) {
  selected.value = item
  detail.value = null
  detailOpen.value = true
  detailLoading.value = true
  try {
    detail.value = await getChannelMonitorStatus(item.id)
  } catch (err) {
    toast.show({ title: '加载详情失败', description: err instanceof Error ? err.message : '请稍后重试', variant: 'destructive' })
  } finally {
    detailLoading.value = false
  }
}

onMounted(() => load())
onBeforeUnmount(() => {
  if (abortController) abortController.abort()
})
</script>

<template>
  <AppShellLayout>
    <div class="mb-8 flex flex-col justify-between gap-4 sm:flex-row sm:items-end">
      <div>
        <Badge variant="outline">Status</Badge>
        <h1 class="mt-3 text-3xl font-semibold tracking-tight">渠道状态</h1>
        <p class="mt-2 text-zinc-500">查看可用模型的健康状态、延迟与近 7 天可用率。</p>
      </div>
      <Button variant="outline" :disabled="refreshing" @click="load(true)">
        <RefreshCw :class="['h-4 w-4', refreshing ? 'animate-spin' : '']" />
        刷新
      </Button>
    </div>

    <div class="grid gap-4 md:grid-cols-3">
      <Card>
        <CardSection role="header" class="pb-3"><div class="text-sm text-zinc-500">正常渠道</div></CardSection>
        <CardSection><div class="flex items-center gap-3 text-3xl font-semibold"><Signal class="h-6 w-6 text-emerald-600" />{{ healthyCount }}</div></CardSection>
      </Card>
      <Card>
        <CardSection role="header" class="pb-3"><div class="text-sm text-zinc-500">异常/降级</div></CardSection>
        <CardSection><div class="flex items-center gap-3 text-3xl font-semibold"><Activity class="h-6 w-6 text-amber-600" />{{ degradedCount }}</div></CardSection>
      </Card>
      <Card>
        <CardSection role="header" class="pb-3"><div class="text-sm text-zinc-500">平均 7D 可用率</div></CardSection>
        <CardSection><div class="flex items-center gap-3 text-3xl font-semibold"><Gauge class="h-6 w-6 text-zinc-500" />{{ formatPercent(avgAvailability) }}</div></CardSection>
      </Card>
    </div>

    <Card class="mt-6">
      <CardSection>
        <div v-if="loading" class="flex h-48 items-center justify-center"><Loader2 class="h-8 w-8 animate-spin text-zinc-500" /></div>
        <Table v-else>
          <thead class="border-b border-zinc-200 text-left text-xs uppercase text-zinc-500 dark:border-zinc-800">
            <tr>
              <th class="px-3 py-3">渠道</th>
              <th class="px-3 py-3">分组</th>
              <th class="px-3 py-3">主模型</th>
              <th class="px-3 py-3">状态</th>
              <th class="px-3 py-3">延迟</th>
              <th class="px-3 py-3">7D 可用率</th>
              <th class="px-3 py-3"></th>
            </tr>
          </thead>
          <tbody class="divide-y divide-zinc-200 dark:divide-zinc-800">
            <tr v-for="item in items" :key="item.id">
              <td class="px-3 py-4 font-medium">{{ item.name }}</td>
              <td class="px-3 py-4 text-sm text-zinc-500">{{ item.group_name }}</td>
              <td class="px-3 py-4 font-mono text-xs">{{ item.primary_model }}</td>
              <td class="px-3 py-4"><Badge :variant="statusVariant(item.primary_status)">{{ item.primary_status }}</Badge></td>
              <td class="px-3 py-4 text-sm">{{ formatLatency(item.primary_latency_ms || item.primary_ping_latency_ms) }}</td>
              <td class="px-3 py-4 text-sm">{{ formatPercent(item.availability_7d) }}</td>
              <td class="px-3 py-4 text-right"><Button variant="outline" size="sm" @click="openDetail(item)">详情</Button></td>
            </tr>
          </tbody>
        </Table>
        <div v-if="!loading && !items.length" class="py-12 text-center text-sm text-zinc-500">暂无可见渠道状态。</div>
      </CardSection>
    </Card>

    <Dialog :open="detailOpen" :title="selected?.name || '渠道详情'" @update:open="detailOpen = $event">
      <div v-if="detailLoading" class="flex h-32 items-center justify-center"><Loader2 class="h-6 w-6 animate-spin text-zinc-500" /></div>
      <div v-else class="space-y-4">
        <div class="text-sm text-zinc-500">{{ detail?.group_name || selected?.group_name }}</div>
        <div class="space-y-3">
          <div v-for="model in detail?.models || []" :key="model.model" class="rounded-lg border border-zinc-200 p-3 dark:border-zinc-800">
            <div class="flex items-center justify-between gap-4">
              <div class="font-mono text-sm">{{ model.model }}</div>
              <Badge :variant="statusVariant(model.latest_status)">{{ model.latest_status }}</Badge>
            </div>
            <div class="mt-3 grid grid-cols-2 gap-3 text-sm sm:grid-cols-4">
              <div><div class="text-zinc-500">延迟</div><div>{{ formatLatency(model.latest_latency_ms) }}</div></div>
              <div><div class="text-zinc-500">7D</div><div>{{ formatPercent(model.availability_7d) }}</div></div>
              <div><div class="text-zinc-500">15D</div><div>{{ formatPercent(model.availability_15d) }}</div></div>
              <div><div class="text-zinc-500">30D</div><div>{{ formatPercent(model.availability_30d) }}</div></div>
            </div>
          </div>
        </div>
      </div>
    </Dialog>
  </AppShellLayout>
</template>
