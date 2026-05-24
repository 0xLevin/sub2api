<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { RouterLink } from 'vue-router'
import { CheckCircle2, Loader2, Mail } from 'lucide-vue-next'
import { forgotPassword, getPublicSettings } from '@/api/auth'
import AuthLayout from '@/brand/layouts/AuthLayout.vue'
import { useBrandToast } from '@/brand/composables/useBrandToast'
import { Alert, Button, Card, CardSection, Input, Label } from '@/brand/ui'

const toast = useBrandToast()

const email = ref('')
const loading = ref(false)
const submitted = ref(false)
const error = ref('')
const turnstileEnabled = ref(false)
const turnstileUnavailable = ref(false)

onMounted(async () => {
  try {
    const settings = await getPublicSettings()
    turnstileEnabled.value = settings.turnstile_enabled === true
    turnstileUnavailable.value = turnstileEnabled.value
  } catch {
    turnstileEnabled.value = false
  }
})

async function submit() {
  error.value = ''
  if (!email.value.trim()) {
    error.value = '请输入邮箱地址。'
    return
  }
  if (turnstileUnavailable.value) {
    error.value = '当前站点启用了人机验证，新品牌页暂未接入该验证组件，请暂时联系管理员重置密码。'
    return
  }

  loading.value = true
  try {
    await forgotPassword({ email: email.value.trim() })
    submitted.value = true
    toast.show({ title: '重置邮件已发送' })
  } catch (err) {
    error.value = err instanceof Error ? err.message : '发送重置邮件失败，请稍后重试。'
    toast.show({ title: '发送失败', description: error.value, variant: 'destructive' })
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <AuthLayout>
    <Card>
      <CardSection role="header">
        <h1 class="text-2xl font-semibold tracking-tight">找回密码</h1>
        <p class="text-sm text-zinc-500">输入注册邮箱，我们会发送密码重置链接。</p>
      </CardSection>
      <CardSection>
        <div v-if="submitted" class="space-y-6 text-center">
          <div class="mx-auto flex h-14 w-14 items-center justify-center rounded-full bg-emerald-100 text-emerald-700 dark:bg-emerald-950 dark:text-emerald-300">
            <CheckCircle2 class="h-7 w-7" />
          </div>
          <div>
            <div class="text-lg font-semibold">请检查邮箱</div>
            <p class="mt-2 text-sm text-zinc-500">如果该邮箱存在账户，你会收到一封密码重置邮件。</p>
          </div>
          <RouterLink
            to="/login"
            class="inline-flex h-10 w-full items-center justify-center rounded-md bg-zinc-950 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-zinc-800 dark:bg-zinc-50 dark:text-zinc-950 dark:hover:bg-zinc-200"
          >
            返回登录
          </RouterLink>
        </div>

        <form v-else class="space-y-4" @submit.prevent="submit">
          <Alert v-if="error" class="border-red-200 text-red-700 dark:border-red-900 dark:text-red-300">{{ error }}</Alert>
          <Alert v-if="turnstileUnavailable" class="border-amber-200 text-amber-700 dark:border-amber-900 dark:text-amber-300">
            当前站点启用了人机验证，新品牌页暂未接入验证组件。
          </Alert>
          <div class="space-y-2">
            <Label for="email">邮箱</Label>
            <div class="relative">
              <Mail class="pointer-events-none absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-zinc-400" />
              <Input id="email" v-model="email" class="pl-9" type="email" placeholder="you@example.com" required />
            </div>
          </div>
          <Button type="submit" class="w-full" :disabled="loading || turnstileUnavailable">
            <Loader2 v-if="loading" class="h-4 w-4 animate-spin" />
            发送重置链接
          </Button>
        </form>

        <p class="mt-6 text-center text-sm text-zinc-500">
          想起来了？
          <RouterLink to="/login" class="underline underline-offset-4">返回登录</RouterLink>
        </p>
      </CardSection>
    </Card>
  </AuthLayout>
</template>
