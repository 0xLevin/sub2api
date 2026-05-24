<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { Check, Loader2 } from 'lucide-vue-next'
import { paymentAPI } from '@/api'
import { useAuthStore } from '@/stores'
import type { SubscriptionPlan } from '@/types/payment'
import { Badge, Button, Card, CardSection, Alert } from '@/brand/ui'

const router = useRouter()
const authStore = useAuthStore()
const plans = ref<SubscriptionPlan[]>([])
const loading = ref(true)
const error = ref('')

const visiblePlans = computed(() => plans.value.filter((plan) => plan.for_sale !== false))

onMounted(async () => {
  try {
    const response = await paymentAPI.getPlans()
    plans.value = response.data
  } catch (err) {
    error.value = err instanceof Error ? err.message : '套餐加载失败'
  } finally {
    loading.value = false
  }
})

function selectPlan(plan: SubscriptionPlan) {
  if (!authStore.isAuthenticated) {
    router.push({ path: '/register', query: { plan: String(plan.id) } })
    return
  }
  router.push({ path: '/app/billing', query: { plan: String(plan.id) } })
}
</script>

<template>
  <MarketingLayout>
    <section class="mx-auto max-w-7xl px-4 py-16 sm:px-6 lg:px-8">
      <div class="mx-auto max-w-3xl text-center">
        <Badge variant="outline">Pricing</Badge>
        <h1 class="mt-5 text-4xl font-semibold tracking-tight md:text-5xl">选择适合你的 API 套餐</h1>
        <p class="mt-4 text-zinc-600 dark:text-zinc-300">套餐信息来自后台配置。商业前台只负责更清晰地展示和转化。</p>
      </div>

      <div v-if="loading" class="mt-12 flex justify-center">
        <Loader2 class="h-8 w-8 animate-spin text-zinc-500" />
      </div>
      <Alert v-else-if="error" class="mx-auto mt-10 max-w-xl border-red-200 text-red-700 dark:border-red-900 dark:text-red-300">
        {{ error }}
      </Alert>
      <div v-else class="mt-12 grid gap-6 md:grid-cols-2 lg:grid-cols-3">
        <Card v-for="plan in visiblePlans" :key="plan.id" class="flex flex-col">
          <CardSection role="header">
            <div class="flex items-start justify-between gap-4">
              <div>
                <div class="text-xl font-semibold">{{ plan.name }}</div>
                <p class="mt-2 text-sm text-zinc-500">{{ plan.description || plan.group_name }}</p>
              </div>
              <Badge v-if="plan.original_price && plan.original_price > plan.price" variant="warning">优惠</Badge>
            </div>
            <div class="mt-6">
              <span class="text-4xl font-semibold">${{ plan.price.toFixed(2) }}</span>
              <span class="text-sm text-zinc-500"> / {{ plan.validity_days }} {{ plan.validity_unit || 'days' }}</span>
            </div>
          </CardSection>
          <CardSection class="flex-1">
            <ul class="space-y-3 text-sm">
              <li v-for="feature in plan.features || []" :key="feature" class="flex gap-2">
                <Check class="mt-0.5 h-4 w-4 text-emerald-600" />
                <span>{{ feature }}</span>
              </li>
              <li v-if="plan.monthly_limit_usd" class="flex gap-2">
                <Check class="mt-0.5 h-4 w-4 text-emerald-600" />
                <span>月用量上限 ${{ plan.monthly_limit_usd }}</span>
              </li>
            </ul>
          </CardSection>
          <CardSection role="footer">
            <Button class="w-full" @click="selectPlan(plan)">选择套餐</Button>
          </CardSection>
        </Card>
      </div>

      <div class="mt-8 text-center text-sm text-zinc-500">
        <RouterLink to="/login" class="underline underline-offset-4">已有账户？登录后继续购买</RouterLink>
      </div>
    </section>
  </MarketingLayout>
</template>
