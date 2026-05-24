<script setup lang="ts">
import { ref } from 'vue'
import { RouterLink, useRoute, useRouter } from 'vue-router'
import { Loader2 } from 'lucide-vue-next'
import { isTotp2FARequired } from '@/api'
import { useAuthStore } from '@/stores'
import type { TotpLoginResponse } from '@/types'
import AuthLayout from '@/brand/layouts/AuthLayout.vue'
import { useBrandToast } from '@/brand/composables/useBrandToast'
import { Alert, Button, Card, CardSection, Dialog, Input, Label } from '@/brand/ui'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
const toast = useBrandToast()

const email = ref('')
const password = ref('')
const loading = ref(false)
const error = ref('')
const showTwoFactor = ref(false)
const twoFactorCode = ref('')
const twoFactorError = ref('')
const twoFactorLoading = ref(false)
const twoFactorTempToken = ref('')
const twoFactorMaskedEmail = ref('')

function resolveRedirect(): string {
  return typeof route.query.redirect === 'string' ? route.query.redirect : '/app'
}

async function submit() {
  error.value = ''
  twoFactorError.value = ''
  loading.value = true
  try {
    const response = await authStore.login({ email: email.value, password: password.value })
    if (isTotp2FARequired(response)) {
      const totpResponse = response as TotpLoginResponse
      twoFactorTempToken.value = totpResponse.temp_token || ''
      twoFactorMaskedEmail.value = totpResponse.user_email_masked || email.value
      showTwoFactor.value = true
      return
    }
    toast.show({ title: '登录成功' })
    router.push(resolveRedirect())
  } catch (err) {
    error.value = err instanceof Error ? err.message : '登录失败，请检查邮箱和密码。'
  } finally {
    loading.value = false
  }
}

async function submitTwoFactor() {
  if (!twoFactorTempToken.value || !twoFactorCode.value.trim()) {
    twoFactorError.value = '请输入 6 位验证码。'
    return
  }

  twoFactorError.value = ''
  twoFactorLoading.value = true
  try {
    await authStore.login2FA(twoFactorTempToken.value, twoFactorCode.value.trim())
    toast.show({ title: '登录成功' })
    showTwoFactor.value = false
    router.push(resolveRedirect())
  } catch (err) {
    twoFactorError.value = err instanceof Error ? err.message : '验证码验证失败，请重试。'
  } finally {
    twoFactorLoading.value = false
  }
}

function cancelTwoFactor() {
  showTwoFactor.value = false
  twoFactorCode.value = ''
  twoFactorError.value = ''
  twoFactorTempToken.value = ''
  twoFactorMaskedEmail.value = ''
}
</script>

<template>
  <AuthLayout>
    <Card>
      <CardSection role="header">
        <h1 class="text-2xl font-semibold tracking-tight">登录账户</h1>
        <p class="text-sm text-zinc-500">进入你的 API 控制台，管理 Key、用量与订单。</p>
      </CardSection>
      <CardSection>
        <form class="space-y-4" @submit.prevent="submit">
          <Alert v-if="error" class="border-red-200 text-red-700 dark:border-red-900 dark:text-red-300">{{ error }}</Alert>
          <div class="space-y-2">
            <Label for="email">邮箱</Label>
            <Input id="email" v-model="email" type="email" placeholder="you@example.com" required />
          </div>
          <div class="space-y-2">
            <Label for="password">密码</Label>
            <Input id="password" v-model="password" type="password" placeholder="输入密码" required />
          </div>
          <Button type="submit" class="w-full" :disabled="loading">
            <Loader2 v-if="loading" class="h-4 w-4 animate-spin" />
            登录
          </Button>
        </form>
        <div class="mt-6 flex items-center justify-between text-sm text-zinc-500">
          <RouterLink to="/register" class="underline underline-offset-4">创建账户</RouterLink>
          <RouterLink to="/forgot-password" class="underline underline-offset-4">忘记密码</RouterLink>
        </div>
      </CardSection>
    </Card>

    <Dialog :open="showTwoFactor" title="二次验证" @update:open="(value) => { if (!value) cancelTwoFactor() }">
      <form class="space-y-4" @submit.prevent="submitTwoFactor">
        <p class="text-sm text-zinc-500">
          {{ twoFactorMaskedEmail }} 已启用二次验证，请输入认证器中的 6 位验证码。
        </p>
        <Alert v-if="twoFactorError" class="border-red-200 text-red-700 dark:border-red-900 dark:text-red-300">
          {{ twoFactorError }}
        </Alert>
        <div class="space-y-2">
          <Label for="two-factor-code">验证码</Label>
          <Input
            id="two-factor-code"
            v-model="twoFactorCode"
            inputmode="numeric"
            autocomplete="one-time-code"
            maxlength="6"
            placeholder="000000"
            required
          />
        </div>
        <div class="flex justify-end gap-2">
          <Button type="button" variant="outline" @click="cancelTwoFactor">取消</Button>
          <Button type="submit" :disabled="twoFactorLoading">
            <Loader2 v-if="twoFactorLoading" class="h-4 w-4 animate-spin" />
            验证并登录
          </Button>
        </div>
      </form>
    </Dialog>
  </AuthLayout>
</template>
