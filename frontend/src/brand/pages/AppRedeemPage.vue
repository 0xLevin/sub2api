<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { CheckCircle2, Gift, Loader2, WalletCards } from 'lucide-vue-next'
import { authAPI, redeemAPI, type RedeemHistoryItem } from '@/api'
import { useAuthStore } from '@/stores/auth'
import AppShellLayout from '@/brand/layouts/AppShellLayout.vue'
import { useBrandToast } from '@/brand/composables/useBrandToast'
import { formatCurrency, formatDate } from '@/brand/lib/format'
import { Alert, Badge, Button, Card, CardSection, Input, Label, Table } from '@/brand/ui'

const authStore = useAuthStore()
const toast = useBrandToast()
const user = computed(() => authStore.user)

const code = ref('')
const submitting = ref(false)
const historyLoading = ref(false)
const history = ref<RedeemHistoryItem[]>([])
const contactInfo = ref('')
const result = ref<{ message: string; type: string; value: number; new_balance?: number; new_concurrency?: number } | null>(null)
const error = ref('')

function formatHistoryValue(item: RedeemHistoryItem) {
  if (item.type === 'balance' || item.type === 'admin_balance') {
    return `${item.value >= 0 ? '+' : ''}${formatCurrency(item.value)}`
  }
  if (item.type === 'subscription') {
    return `${item.validity_days || Math.round(item.value)} 天${item.group?.name ? ` - ${item.group.name}` : ''}`
  }
  return `${item.value >= 0 ? '+' : ''}${item.value} 并发`
}

function titleFor(item: RedeemHistoryItem) {
  if (item.type === 'subscription') return '订阅兑换'
  if (item.type.includes('balance')) return item.type.startsWith('admin') ? '余额调整' : '余额兑换'
  if (item.type.includes('concurrency')) return item.type.startsWith('admin') ? '并发调整' : '并发兑换'
  return item.type
}

async function loadHistory() {
  historyLoading.value = true
  try {
    history.value = await redeemAPI.getHistory()
  } finally {
    historyLoading.value = false
  }
}

async function submit() {
  if (!code.value.trim()) {
    error.value = '请输入兑换码。'
    return
  }
  submitting.value = true
  error.value = ''
  result.value = null
  try {
    result.value = await redeemAPI.redeem(code.value.trim())
    code.value = ''
    await Promise.allSettled([authStore.refreshUser(), loadHistory()])
    toast.show({ title: '兑换成功' })
  } catch (err) {
    error.value = err instanceof Error ? err.message : '兑换失败，请检查兑换码。'
    toast.show({ title: '兑换失败', description: error.value, variant: 'destructive' })
  } finally {
    submitting.value = false
  }
}

onMounted(async () => {
  void loadHistory()
  try {
    const settings = await authAPI.getPublicSettings()
    contactInfo.value = settings.contact_info || ''
  } catch {
    contactInfo.value = ''
  }
})
</script>

<template>
  <AppShellLayout>
    <div class="mb-8">
      <Badge variant="outline">Redeem</Badge>
      <h1 class="mt-3 text-3xl font-semibold tracking-tight">兑换</h1>
      <p class="mt-2 text-zinc-500">使用兑换码添加余额、并发或订阅套餐。</p>
    </div>

    <div class="grid gap-6 lg:grid-cols-[0.8fr_1.2fr]">
      <div class="space-y-6">
        <Card>
          <CardSection role="header">
            <div class="flex items-center gap-3">
              <div class="rounded-md bg-zinc-100 p-2 dark:bg-zinc-800"><WalletCards class="h-5 w-5" /></div>
              <div>
                <div class="text-sm text-zinc-500">当前余额</div>
                <div class="text-3xl font-semibold">{{ formatCurrency(user?.balance || 0) }}</div>
              </div>
            </div>
          </CardSection>
          <CardSection>
            <div class="rounded-lg bg-zinc-50 p-4 text-sm dark:bg-zinc-900">
              <div class="text-zinc-500">并发额度</div>
              <div class="mt-1 text-xl font-semibold">{{ user?.concurrency || 0 }}</div>
            </div>
          </CardSection>
        </Card>

        <Card>
          <CardSection role="header">
            <div class="text-lg font-semibold">输入兑换码</div>
            <div class="text-sm text-zinc-500">兑换成功后会立即刷新账户状态。</div>
          </CardSection>
          <CardSection>
            <form class="space-y-4" @submit.prevent="submit">
              <Alert v-if="error" class="border-red-200 text-red-700 dark:border-red-900 dark:text-red-300">{{ error }}</Alert>
              <Alert v-if="result" class="border-emerald-200 text-emerald-700 dark:border-emerald-900 dark:text-emerald-300">
                <div class="flex items-start gap-2"><CheckCircle2 class="mt-0.5 h-4 w-4" /><span>{{ result.message }}</span></div>
              </Alert>
              <div class="space-y-2">
                <Label for="redeem-code">兑换码</Label>
                <div class="relative">
                  <Gift class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-zinc-400" />
                  <Input id="redeem-code" v-model="code" class="pl-9 font-mono" placeholder="XXXX-XXXX-XXXX" />
                </div>
              </div>
              <Button type="submit" class="w-full" :disabled="submitting || !code.trim()">
                <Loader2 v-if="submitting" class="h-4 w-4 animate-spin" />
                立即兑换
              </Button>
            </form>
          </CardSection>
        </Card>

        <Alert v-if="contactInfo">如兑换码有问题，请联系：{{ contactInfo }}</Alert>
      </div>

      <Card>
        <CardSection role="header">
          <div class="text-lg font-semibold">兑换记录</div>
          <div class="text-sm text-zinc-500">最近的兑换和管理员调整记录。</div>
        </CardSection>
        <CardSection>
          <div v-if="historyLoading" class="flex h-48 items-center justify-center"><Loader2 class="h-7 w-7 animate-spin text-zinc-500" /></div>
          <Table v-else>
            <thead class="border-b border-zinc-200 text-left text-xs uppercase text-zinc-500 dark:border-zinc-800">
              <tr>
                <th class="px-3 py-3">类型</th>
                <th class="px-3 py-3">值</th>
                <th class="px-3 py-3">时间</th>
                <th class="px-3 py-3">码</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-zinc-200 dark:divide-zinc-800">
              <tr v-for="item in history" :key="item.id">
                <td class="px-3 py-4 text-sm font-medium">{{ titleFor(item) }}</td>
                <td class="px-3 py-4 text-sm">{{ formatHistoryValue(item) }}</td>
                <td class="px-3 py-4 text-sm text-zinc-500">{{ formatDate(item.used_at || item.created_at) }}</td>
                <td class="px-3 py-4 font-mono text-xs text-zinc-500">{{ item.code ? `${item.code.slice(0, 8)}...` : '-' }}</td>
              </tr>
            </tbody>
          </Table>
          <div v-if="!historyLoading && !history.length" class="py-12 text-center text-sm text-zinc-500">暂无兑换记录。</div>
        </CardSection>
      </Card>
    </div>
  </AppShellLayout>
</template>
