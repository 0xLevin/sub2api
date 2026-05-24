<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { Copy, Loader2, Plus, Trash2 } from 'lucide-vue-next'
import { keysAPI } from '@/api'
import type { ApiKey } from '@/types'
import AppShellLayout from '@/brand/layouts/AppShellLayout.vue'
import { useBrandToast } from '@/brand/composables/useBrandToast'
import { formatCurrency, formatDate } from '@/brand/lib/format'
import { Badge, Button, Card, CardSection, Dialog, Input, Label, Table } from '@/brand/ui'

const toast = useBrandToast()
const loading = ref(true)
const saving = ref(false)
const dialogOpen = ref(false)
const keys = ref<ApiKey[]>([])
const name = ref('Production Key')

async function loadKeys() {
  loading.value = true
  try {
    const response = await keysAPI.list(1, 20)
    keys.value = response.items
  } finally {
    loading.value = false
  }
}

async function createKey() {
  saving.value = true
  try {
    await keysAPI.create(name.value || 'API Key')
    dialogOpen.value = false
    toast.show({ title: 'API Key 已创建' })
    await loadKeys()
  } catch (err) {
    toast.show({ title: '创建失败', description: err instanceof Error ? err.message : '请稍后重试', variant: 'destructive' })
  } finally {
    saving.value = false
  }
}

async function copyKey(value: string) {
  await navigator.clipboard.writeText(value)
  toast.show({ title: '已复制到剪贴板' })
}

async function deleteKey(key: ApiKey) {
  if (!window.confirm(`删除 ${key.name}？`)) return
  await keysAPI.delete(key.id)
  toast.show({ title: 'API Key 已删除' })
  await loadKeys()
}

onMounted(loadKeys)
</script>

<template>
  <AppShellLayout>
    <div class="mb-8 flex flex-col justify-between gap-4 sm:flex-row sm:items-end">
      <div>
        <Badge variant="outline">API Keys</Badge>
        <h1 class="mt-3 text-3xl font-semibold tracking-tight">API Key 管理</h1>
        <p class="mt-2 text-zinc-500">创建、复制和停用客户侧调用密钥。</p>
      </div>
      <Button @click="dialogOpen = true"><Plus class="h-4 w-4" /> 创建 Key</Button>
    </div>

    <Card>
      <CardSection>
        <div v-if="loading" class="flex h-40 items-center justify-center">
          <Loader2 class="h-8 w-8 animate-spin text-zinc-500" />
        </div>
        <Table v-else>
          <thead class="border-b border-zinc-200 text-left text-xs uppercase text-zinc-500 dark:border-zinc-800">
            <tr>
              <th class="px-3 py-3">名称</th>
              <th class="px-3 py-3">Key</th>
              <th class="px-3 py-3">状态</th>
              <th class="px-3 py-3">额度</th>
              <th class="px-3 py-3">创建时间</th>
              <th class="px-3 py-3 text-right">操作</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-zinc-200 dark:divide-zinc-800">
            <tr v-for="key in keys" :key="key.id">
              <td class="px-3 py-4 font-medium">{{ key.name }}</td>
              <td class="px-3 py-4 font-mono text-xs text-zinc-500">{{ key.key }}</td>
              <td class="px-3 py-4"><Badge :variant="key.status === 'active' ? 'success' : 'secondary'">{{ key.status }}</Badge></td>
              <td class="px-3 py-4 text-sm">{{ formatCurrency(key.quota_used) }} / {{ key.quota ? formatCurrency(key.quota) : '不限' }}</td>
              <td class="px-3 py-4 text-sm text-zinc-500">{{ formatDate(key.created_at) }}</td>
              <td class="px-3 py-4">
                <div class="flex justify-end gap-2">
                  <Button variant="outline" size="icon" @click="copyKey(key.key)"><Copy class="h-4 w-4" /></Button>
                  <Button variant="ghost" size="icon" @click="deleteKey(key)"><Trash2 class="h-4 w-4" /></Button>
                </div>
              </td>
            </tr>
          </tbody>
        </Table>
        <div v-if="!loading && !keys.length" class="py-12 text-center text-sm text-zinc-500">
          还没有 API Key，创建一个开始接入。
        </div>
      </CardSection>
    </Card>

    <Dialog v-model:open="dialogOpen" title="创建 API Key">
      <div class="space-y-4">
        <div class="space-y-2">
          <Label for="key-name">Key 名称</Label>
          <Input id="key-name" v-model="name" />
        </div>
        <Button class="w-full" :disabled="saving" @click="createKey">
          <Loader2 v-if="saving" class="h-4 w-4 animate-spin" />
          创建
        </Button>
      </div>
    </Dialog>
  </AppShellLayout>
</template>
