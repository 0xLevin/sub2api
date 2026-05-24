<script setup lang="ts">
import { ref } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { Loader2 } from 'lucide-vue-next'
import { useAuthStore } from '@/stores'
import AuthLayout from '@/brand/layouts/AuthLayout.vue'
import { useBrandToast } from '@/brand/composables/useBrandToast'
import { Alert, Button, Card, CardSection, Input, Label } from '@/brand/ui'

const router = useRouter()
const authStore = useAuthStore()
const toast = useBrandToast()

const email = ref('')
const password = ref('')
const promoCode = ref('')
const loading = ref(false)
const error = ref('')

async function submit() {
  error.value = ''
  loading.value = true
  try {
    await authStore.register({
      email: email.value,
      password: password.value,
      promo_code: promoCode.value || undefined
    })
    toast.show({ title: '注册成功', description: '欢迎进入控制台。' })
    router.push('/app')
  } catch (err) {
    error.value = err instanceof Error ? err.message : '注册失败，请稍后重试。'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <AuthLayout>
    <Card>
      <CardSection role="header">
        <h1 class="text-2xl font-semibold tracking-tight">创建账户</h1>
        <p class="text-sm text-zinc-500">注册后即可购买套餐并创建 API Key。</p>
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
            <Input id="password" v-model="password" type="password" placeholder="至少 8 个字符" required />
          </div>
          <div class="space-y-2">
            <Label for="promo">优惠码（可选）</Label>
            <Input id="promo" v-model="promoCode" placeholder="如果有优惠码，在这里输入" />
          </div>
          <Button type="submit" class="w-full" :disabled="loading">
            <Loader2 v-if="loading" class="h-4 w-4 animate-spin" />
            注册并进入控制台
          </Button>
        </form>
        <p class="mt-6 text-sm text-zinc-500">
          已有账户？
          <RouterLink to="/login" class="underline underline-offset-4">登录</RouterLink>
        </p>
      </CardSection>
    </Card>
  </AuthLayout>
</template>
