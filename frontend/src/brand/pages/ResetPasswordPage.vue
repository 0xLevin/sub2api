<script setup lang="ts">
import { computed, reactive, ref } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import { CheckCircle2, Eye, EyeOff, Loader2, TriangleAlert } from 'lucide-vue-next'
import { resetPassword } from '@/api/auth'
import AuthLayout from '@/brand/layouts/AuthLayout.vue'
import { useBrandToast } from '@/brand/composables/useBrandToast'
import { Alert, Button, Card, CardSection, Input, Label } from '@/brand/ui'

const route = useRoute()
const toast = useBrandToast()

const email = computed(() => typeof route.query.email === 'string' ? route.query.email : '')
const token = computed(() => typeof route.query.token === 'string' ? route.query.token : '')
const invalidLink = computed(() => !email.value || !token.value)

const form = reactive({
  password: '',
  confirmPassword: ''
})
const loading = ref(false)
const success = ref(false)
const error = ref('')
const showPassword = ref(false)
const showConfirmPassword = ref(false)

function validate(): boolean {
  if (!form.password) {
    error.value = '请输入新密码。'
    return false
  }
  if (form.password.length < 6) {
    error.value = '密码至少需要 6 个字符。'
    return false
  }
  if (form.password !== form.confirmPassword) {
    error.value = '两次输入的密码不一致。'
    return false
  }
  return true
}

async function submit() {
  error.value = ''
  if (invalidLink.value || !validate()) return

  loading.value = true
  try {
    await resetPassword({
      email: email.value,
      token: token.value,
      new_password: form.password
    })
    success.value = true
    toast.show({ title: '密码已重置' })
  } catch (err) {
    error.value = err instanceof Error ? err.message : '密码重置失败，链接可能已经过期。'
    toast.show({ title: '重置失败', description: error.value, variant: 'destructive' })
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <AuthLayout>
    <Card>
      <CardSection role="header">
        <h1 class="text-2xl font-semibold tracking-tight">重置密码</h1>
        <p class="text-sm text-zinc-500">设置一个新的登录密码。</p>
      </CardSection>
      <CardSection>
        <div v-if="invalidLink" class="space-y-6 text-center">
          <div class="mx-auto flex h-14 w-14 items-center justify-center rounded-full bg-amber-100 text-amber-700 dark:bg-amber-950 dark:text-amber-300">
            <TriangleAlert class="h-7 w-7" />
          </div>
          <div>
            <div class="text-lg font-semibold">链接无效或已过期</div>
            <p class="mt-2 text-sm text-zinc-500">请重新申请密码重置邮件。</p>
          </div>
          <RouterLink
            to="/forgot-password"
            class="inline-flex h-10 w-full items-center justify-center rounded-md bg-zinc-950 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-zinc-800 dark:bg-zinc-50 dark:text-zinc-950 dark:hover:bg-zinc-200"
          >
            重新发送重置链接
          </RouterLink>
        </div>

        <div v-else-if="success" class="space-y-6 text-center">
          <div class="mx-auto flex h-14 w-14 items-center justify-center rounded-full bg-emerald-100 text-emerald-700 dark:bg-emerald-950 dark:text-emerald-300">
            <CheckCircle2 class="h-7 w-7" />
          </div>
          <div>
            <div class="text-lg font-semibold">密码已更新</div>
            <p class="mt-2 text-sm text-zinc-500">现在可以使用新密码登录。</p>
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
          <div class="space-y-2">
            <Label for="email">邮箱</Label>
            <Input id="email" :model-value="email" type="email" disabled />
          </div>
          <div class="space-y-2">
            <Label for="password">新密码</Label>
            <div class="relative">
              <Input
                id="password"
                v-model="form.password"
                :type="showPassword ? 'text' : 'password'"
                class="pr-10"
                autocomplete="new-password"
                placeholder="至少 6 个字符"
                required
              />
              <button type="button" class="absolute right-3 top-1/2 -translate-y-1/2 text-zinc-400" @click="showPassword = !showPassword">
                <EyeOff v-if="showPassword" class="h-4 w-4" />
                <Eye v-else class="h-4 w-4" />
              </button>
            </div>
          </div>
          <div class="space-y-2">
            <Label for="confirm-password">确认密码</Label>
            <div class="relative">
              <Input
                id="confirm-password"
                v-model="form.confirmPassword"
                :type="showConfirmPassword ? 'text' : 'password'"
                class="pr-10"
                autocomplete="new-password"
                placeholder="再次输入新密码"
                required
              />
              <button type="button" class="absolute right-3 top-1/2 -translate-y-1/2 text-zinc-400" @click="showConfirmPassword = !showConfirmPassword">
                <EyeOff v-if="showConfirmPassword" class="h-4 w-4" />
                <Eye v-else class="h-4 w-4" />
              </button>
            </div>
          </div>
          <Button type="submit" class="w-full" :disabled="loading">
            <Loader2 v-if="loading" class="h-4 w-4 animate-spin" />
            重置密码
          </Button>
        </form>
      </CardSection>
    </Card>
  </AuthLayout>
</template>
