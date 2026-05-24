<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'
import { Bell, Loader2, LockKeyhole, Mail, Save, Shield, UserRound } from 'lucide-vue-next'
import { userAPI } from '@/api/user'
import { useAppStore, useAuthStore } from '@/stores'
import type { NotifyEmailEntry } from '@/types'
import AppShellLayout from '@/brand/layouts/AppShellLayout.vue'
import { useBrandToast } from '@/brand/composables/useBrandToast'
import { formatCurrency, formatDate } from '@/brand/lib/format'
import { Alert, Badge, Button, Card, CardSection, Input, Label } from '@/brand/ui'

const authStore = useAuthStore()
const appStore = useAppStore()
const toast = useBrandToast()
const user = computed(() => authStore.user)

const loading = ref(true)
const savingProfile = ref(false)
const savingPassword = ref(false)
const savingNotify = ref(false)
const contactInfo = ref('')
const balanceNotifyEnabled = ref(false)
const systemDefaultThreshold = ref(0)

const profileForm = reactive({ username: '', avatar_url: '' })
const passwordForm = reactive({ oldPassword: '', newPassword: '', confirmPassword: '' })
const notifyForm = reactive({ enabled: true, threshold: '', extraEmailsText: '' })

function syncFromUser() {
  profileForm.username = user.value?.username || ''
  profileForm.avatar_url = user.value?.avatar_url || ''
  notifyForm.enabled = user.value?.balance_notify_enabled ?? true
  notifyForm.threshold = user.value?.balance_notify_threshold == null ? '' : String(user.value.balance_notify_threshold)
  notifyForm.extraEmailsText = (user.value?.balance_notify_extra_emails || [])
    .filter((entry) => !entry.disabled)
    .map((entry) => entry.email)
    .join('\n')
}

function notifyEntries(): NotifyEmailEntry[] {
  return notifyForm.extraEmailsText
    .split('\n')
    .map((email) => email.trim())
    .filter(Boolean)
    .map((email) => ({ email, disabled: false, verified: true }))
}

async function saveProfile() {
  savingProfile.value = true
  try {
    const updated = await userAPI.updateProfile({
      username: profileForm.username.trim() || undefined,
      avatar_url: profileForm.avatar_url.trim() || null
    })
    localStorage.setItem('auth_user', JSON.stringify(updated))
    await authStore.refreshUser()
    syncFromUser()
    toast.show({ title: '资料已保存' })
  } catch (err) {
    toast.show({ title: '保存失败', description: err instanceof Error ? err.message : '请稍后重试', variant: 'destructive' })
  } finally {
    savingProfile.value = false
  }
}

async function savePassword() {
  if (!passwordForm.oldPassword || !passwordForm.newPassword) {
    toast.show({ title: '请输入当前密码和新密码', variant: 'destructive' })
    return
  }
  if (passwordForm.newPassword !== passwordForm.confirmPassword) {
    toast.show({ title: '两次输入的新密码不一致', variant: 'destructive' })
    return
  }
  savingPassword.value = true
  try {
    await userAPI.changePassword(passwordForm.oldPassword, passwordForm.newPassword)
    passwordForm.oldPassword = ''
    passwordForm.newPassword = ''
    passwordForm.confirmPassword = ''
    toast.show({ title: '密码已更新' })
  } catch (err) {
    toast.show({ title: '密码更新失败', description: err instanceof Error ? err.message : '请稍后重试', variant: 'destructive' })
  } finally {
    savingPassword.value = false
  }
}

async function saveNotifySettings() {
  savingNotify.value = true
  try {
    await userAPI.updateProfile({
      balance_notify_enabled: notifyForm.enabled,
      balance_notify_threshold: notifyForm.threshold === '' ? null : Number(notifyForm.threshold),
      balance_notify_extra_emails: notifyEntries()
    })
    await authStore.refreshUser()
    syncFromUser()
    toast.show({ title: '通知设置已保存' })
  } catch (err) {
    toast.show({ title: '通知设置保存失败', description: err instanceof Error ? err.message : '请稍后重试', variant: 'destructive' })
  } finally {
    savingNotify.value = false
  }
}

onMounted(async () => {
  try {
    await Promise.allSettled([authStore.refreshUser(), appStore.fetchPublicSettings()])
    const settings = appStore.cachedPublicSettings
    contactInfo.value = settings?.contact_info || ''
    balanceNotifyEnabled.value = settings?.balance_low_notify_enabled ?? false
    systemDefaultThreshold.value = settings?.balance_low_notify_threshold ?? 0
    syncFromUser()
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <AppShellLayout>
    <div class="mb-8">
      <Badge variant="outline">Profile</Badge>
      <h1 class="mt-3 text-3xl font-semibold tracking-tight">个人资料</h1>
      <p class="mt-2 text-zinc-500">管理账户资料、安全密码和余额提醒。</p>
    </div>

    <div v-if="loading" class="flex h-48 items-center justify-center"><Loader2 class="h-8 w-8 animate-spin text-zinc-500" /></div>
    <div v-else class="grid gap-6 xl:grid-cols-[0.8fr_1.2fr]">
      <div class="space-y-6">
        <Card>
          <CardSection>
            <div class="flex items-center gap-4">
              <div class="flex h-16 w-16 items-center justify-center overflow-hidden rounded-full bg-zinc-100 dark:bg-zinc-800">
                <img v-if="user?.avatar_url" :src="user.avatar_url" alt="" class="h-full w-full object-cover">
                <UserRound v-else class="h-8 w-8 text-zinc-500" />
              </div>
              <div class="min-w-0">
                <div class="truncate text-xl font-semibold">{{ user?.username || user?.email }}</div>
                <div class="truncate text-sm text-zinc-500">{{ user?.email }}</div>
                <div class="mt-2 flex gap-2"><Badge>{{ user?.role }}</Badge><Badge :variant="user?.status === 'active' ? 'success' : 'warning'">{{ user?.status }}</Badge></div>
              </div>
            </div>
          </CardSection>
        </Card>

        <Card>
          <CardSection role="header"><div class="text-lg font-semibold">账户状态</div></CardSection>
          <CardSection>
            <div class="grid gap-3 text-sm">
              <div class="flex justify-between"><span class="text-zinc-500">余额</span><span class="font-medium">{{ formatCurrency(user?.balance || 0) }}</span></div>
              <div class="flex justify-between"><span class="text-zinc-500">并发</span><span class="font-medium">{{ user?.concurrency || 0 }}</span></div>
              <div class="flex justify-between"><span class="text-zinc-500">RPM</span><span class="font-medium">{{ user?.rpm_limit || '不限' }}</span></div>
              <div class="flex justify-between"><span class="text-zinc-500">创建时间</span><span class="font-medium">{{ formatDate(user?.created_at) }}</span></div>
            </div>
          </CardSection>
        </Card>

        <Alert v-if="contactInfo">需要账户支持：{{ contactInfo }}</Alert>
      </div>

      <div class="space-y-6">
        <Card>
          <CardSection role="header">
            <div class="flex items-center gap-2 text-lg font-semibold"><UserRound class="h-5 w-5" />基础资料</div>
          </CardSection>
          <CardSection>
            <form class="space-y-4" @submit.prevent="saveProfile">
              <div class="space-y-2">
                <Label for="username">用户名</Label>
                <Input id="username" v-model="profileForm.username" placeholder="显示名称" />
              </div>
              <div class="space-y-2">
                <Label for="avatar">头像 URL</Label>
                <Input id="avatar" v-model="profileForm.avatar_url" placeholder="https://..." />
              </div>
              <Button type="submit" :disabled="savingProfile"><Loader2 v-if="savingProfile" class="h-4 w-4 animate-spin" /><Save v-else class="h-4 w-4" />保存资料</Button>
            </form>
          </CardSection>
        </Card>

        <Card>
          <CardSection role="header">
            <div class="flex items-center gap-2 text-lg font-semibold"><LockKeyhole class="h-5 w-5" />修改密码</div>
          </CardSection>
          <CardSection>
            <form class="space-y-4" @submit.prevent="savePassword">
              <div class="space-y-2"><Label for="old-password">当前密码</Label><Input id="old-password" v-model="passwordForm.oldPassword" type="password" /></div>
              <div class="space-y-2"><Label for="new-password">新密码</Label><Input id="new-password" v-model="passwordForm.newPassword" type="password" /></div>
              <div class="space-y-2"><Label for="confirm-password">确认新密码</Label><Input id="confirm-password" v-model="passwordForm.confirmPassword" type="password" /></div>
              <Button type="submit" :disabled="savingPassword"><Loader2 v-if="savingPassword" class="h-4 w-4 animate-spin" /><Shield v-else class="h-4 w-4" />更新密码</Button>
            </form>
          </CardSection>
        </Card>

        <Card>
          <CardSection role="header">
            <div class="flex items-center gap-2 text-lg font-semibold"><Bell class="h-5 w-5" />余额提醒</div>
            <div class="text-sm text-zinc-500">系统默认阈值：{{ formatCurrency(systemDefaultThreshold) }}</div>
          </CardSection>
          <CardSection>
            <form class="space-y-4" @submit.prevent="saveNotifySettings">
              <Alert v-if="!balanceNotifyEnabled">系统未启用余额低额提醒，保存后仅会保留你的偏好。</Alert>
              <label class="flex items-center gap-3 text-sm">
                <input v-model="notifyForm.enabled" type="checkbox" class="h-4 w-4 rounded border-zinc-300">
                启用余额低额提醒
              </label>
              <div class="space-y-2">
                <Label for="threshold">自定义提醒阈值</Label>
                <Input id="threshold" v-model="notifyForm.threshold" type="number" min="0" step="0.01" placeholder="留空使用系统默认值" />
              </div>
              <div class="space-y-2">
                <Label for="notify-emails">额外通知邮箱</Label>
                <textarea id="notify-emails" v-model="notifyForm.extraEmailsText" rows="4" class="w-full rounded-md border border-zinc-200 bg-transparent px-3 py-2 text-sm shadow-sm outline-none focus-visible:ring-2 focus-visible:ring-zinc-950 dark:border-zinc-800 dark:focus-visible:ring-zinc-300" placeholder="每行一个邮箱" />
                <div class="flex items-center gap-2 text-xs text-zinc-500"><Mail class="h-3.5 w-3.5" />保存后将作为通知目标。</div>
              </div>
              <Button type="submit" :disabled="savingNotify"><Loader2 v-if="savingNotify" class="h-4 w-4 animate-spin" /><Save v-else class="h-4 w-4" />保存通知设置</Button>
            </form>
          </CardSection>
        </Card>
      </div>
    </div>
  </AppShellLayout>
</template>
