<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Check, Loader2 } from 'lucide-vue-next'
import { paymentAPI } from '@/api'
import {
  PAYMENT_RECOVERY_STORAGE_KEY,
  buildCreateOrderPayload,
  decidePaymentLaunch,
  normalizeVisibleMethod,
  writePaymentRecoverySnapshot,
} from '@/components/payment/paymentFlow'
import { getPaymentPopupFeatures } from '@/components/payment/providerConfig'
import { isMobileDevice } from '@/utils/device'
import type { CreateOrderResult, PaymentOrder, PaymentType, SubscriptionPlan } from '@/types/payment'
import AppShellLayout from '@/brand/layouts/AppShellLayout.vue'
import { useBrandToast } from '@/brand/composables/useBrandToast'
import { formatCurrency, formatDate } from '@/brand/lib/format'
import { Alert, Badge, Button, Card, CardSection, Table } from '@/brand/ui'

const route = useRoute()
const router = useRouter()
const toast = useBrandToast()

const loading = ref(true)
const creatingPlanId = ref<number | null>(null)
const plans = ref<SubscriptionPlan[]>([])
const orders = ref<PaymentOrder[]>([])
const methods = ref<PaymentType[]>([])
const checkoutResult = ref<CreateOrderResult | null>(null)
const error = ref('')
const alipayForceQrCode = ref(false)

const visiblePlans = computed(() => plans.value.filter((plan) => plan.for_sale !== false))
const selectedMethod = computed<PaymentType>(() => methods.value[0] || 'stripe')

onMounted(async () => {
  try {
    const [checkoutInfo, orderResult] = await Promise.allSettled([
      paymentAPI.getCheckoutInfo(),
      paymentAPI.getMyOrders({ page: 1, page_size: 10 })
    ])
    if (checkoutInfo.status === 'fulfilled') {
      plans.value = checkoutInfo.value.data.plans || []
      methods.value = Object.keys(checkoutInfo.value.data.methods || {}) as PaymentType[]
      alipayForceQrCode.value = checkoutInfo.value.data.alipay_force_qrcode === true
    }
    if (orderResult.status === 'fulfilled') {
      orders.value = orderResult.value.data.items || []
    }
  } catch (err) {
    error.value = err instanceof Error ? err.message : '支付信息加载失败'
  } finally {
    loading.value = false
  }
})

function openPaymentWindow(url: string) {
  const win = window.open(url, 'paymentPopup', getPaymentPopupFeatures())
  if (!win || win.closed) {
    window.location.href = url
  }
}

function persistPaymentRecovery(result: CreateOrderResult, payUrl: string, visibleMethod: string) {
  if (typeof window === 'undefined' || !result.order_id) return

  writePaymentRecoverySnapshot(window.localStorage, {
    orderId: result.order_id,
    amount: result.amount,
    qrCode: result.qr_code || '',
    expiresAt: result.expires_at || '',
    paymentType: visibleMethod,
    payUrl,
    outTradeNo: result.out_trade_no || '',
    clientSecret: result.client_secret || '',
    intentId: result.intent_id || '',
    currency: result.currency || '',
    countryCode: result.country_code || '',
    paymentEnv: result.payment_env || '',
    payAmount: result.pay_amount,
    orderType: 'subscription',
    paymentMode: (result.payment_mode || '').trim(),
    resumeToken: result.resume_token || '',
    createdAt: Date.now(),
  }, PAYMENT_RECOVERY_STORAGE_KEY)
}

async function createOrder(plan: SubscriptionPlan) {
  creatingPlanId.value = plan.id
  checkoutResult.value = null
  try {
    const requestType = selectedMethod.value
    const visibleMethod = normalizeVisibleMethod(requestType) || requestType
    const response = await paymentAPI.createOrder(buildCreateOrderPayload({
      amount: plan.price,
      paymentType: requestType,
      orderType: 'subscription',
      planId: plan.id,
      origin: typeof window !== 'undefined' ? window.location.origin : '',
      isMobile: isMobileDevice(),
      isWechatBrowser: typeof window !== 'undefined' && /MicroMessenger/i.test(window.navigator.userAgent),
      forceQRCode: alipayForceQrCode.value && visibleMethod === 'alipay',
    }))
    const result = response.data
    checkoutResult.value = result

    const stripeMethod = visibleMethod === 'stripe'
      ? ''
      : visibleMethod === 'wxpay' ? 'wechat_pay' : 'alipay'
    const stripeRouteUrl = result.client_secret && visibleMethod !== 'airwallex'
      ? router.resolve({
        path: '/payment/stripe',
        query: {
          order_id: String(result.order_id),
          client_secret: result.client_secret,
          method: stripeMethod || undefined,
          resume_token: result.resume_token || undefined,
        },
      }).href
      : ''
    const airwallexRouteUrl = result.client_secret && result.intent_id
      ? router.resolve({
        path: '/payment/airwallex',
        query: {
          order_id: String(result.order_id),
          out_trade_no: result.out_trade_no || undefined,
          resume_token: result.resume_token || undefined,
        },
      }).href
      : ''
    const decision = decidePaymentLaunch(result, {
      visibleMethod,
      orderType: 'subscription',
      isMobile: isMobileDevice(),
      isWechatBrowser: typeof window !== 'undefined' && /MicroMessenger/i.test(window.navigator.userAgent),
      forceQRCode: alipayForceQrCode.value && visibleMethod === 'alipay',
      stripePopupUrl: stripeRouteUrl,
      stripeRouteUrl,
      airwallexRouteUrl,
    })

    if (decision.kind === 'wechat_oauth' && decision.oauth?.authorize_url) {
      window.location.href = decision.oauth.authorize_url
      return
    }
    if (decision.kind === 'unhandled') {
      toast.show({ title: '订单已创建', description: '请使用下方支付信息完成付款。' })
      return
    }

    persistPaymentRecovery(result, decision.paymentState.payUrl, visibleMethod)
    toast.show({ title: '订单已创建', description: '正在打开支付页面。' })

    if (decision.kind === 'stripe_popup') {
      openPaymentWindow(decision.paymentState.payUrl)
      return
    }
    if (['stripe_route', 'airwallex_route'].includes(decision.kind)) {
      window.location.href = decision.paymentState.payUrl
      return
    }
    if (decision.kind === 'redirect_waiting' && decision.paymentState.payUrl) {
      if (isMobileDevice()) {
        window.location.href = decision.paymentState.payUrl
        return
      }
      openPaymentWindow(decision.paymentState.payUrl)
    }
  } catch (err) {
    toast.show({ title: '创建订单失败', description: err instanceof Error ? err.message : '请稍后重试', variant: 'destructive' })
  } finally {
    creatingPlanId.value = null
  }
}
</script>

<template>
  <AppShellLayout>
    <div class="mb-8">
      <Badge variant="outline">Billing</Badge>
      <h1 class="mt-3 text-3xl font-semibold tracking-tight">购买与订单</h1>
      <p class="mt-2 text-zinc-500">购买套餐，查看最近订单状态。</p>
    </div>

    <div v-if="loading" class="flex h-48 items-center justify-center">
      <Loader2 class="h-8 w-8 animate-spin text-zinc-500" />
    </div>
    <div v-else class="space-y-6">
      <Alert v-if="error" class="border-red-200 text-red-700 dark:border-red-900 dark:text-red-300">{{ error }}</Alert>
      <Alert v-if="route.query.plan" class="border-zinc-200 dark:border-zinc-800">已从价格页带入套餐选择，可直接在下方创建订单。</Alert>
      <Alert v-if="checkoutResult">
        <div class="space-y-2">
          <div class="font-medium">订单创建成功</div>
          <a v-if="checkoutResult.pay_url" :href="checkoutResult.pay_url" target="_blank" rel="noopener noreferrer" class="text-sm underline underline-offset-4">打开支付链接</a>
          <div v-if="checkoutResult.qr_code" class="break-all text-sm text-zinc-500">二维码地址：{{ checkoutResult.qr_code }}</div>
          <div v-if="checkoutResult.out_trade_no" class="font-mono text-xs text-zinc-500">订单号：{{ checkoutResult.out_trade_no }}</div>
        </div>
      </Alert>

      <div class="grid gap-6 md:grid-cols-2 xl:grid-cols-3">
        <Card v-for="plan in visiblePlans" :key="plan.id">
          <CardSection role="header">
            <div class="flex items-start justify-between gap-4">
              <div>
                <div class="text-xl font-semibold">{{ plan.name }}</div>
                <p class="mt-2 text-sm text-zinc-500">{{ plan.description || plan.group_name }}</p>
              </div>
              <Badge v-if="String(route.query.plan || '') === String(plan.id)" variant="success">已选择</Badge>
            </div>
            <div class="mt-6 text-4xl font-semibold">{{ formatCurrency(plan.price) }}</div>
          </CardSection>
          <CardSection>
            <ul class="space-y-3 text-sm">
              <li v-for="feature in plan.features || []" :key="feature" class="flex gap-2">
                <Check class="mt-0.5 h-4 w-4 text-emerald-600" />
                <span>{{ feature }}</span>
              </li>
            </ul>
          </CardSection>
          <CardSection role="footer">
            <Button class="w-full" :disabled="creatingPlanId === plan.id || !methods.length" @click="createOrder(plan)">
              <Loader2 v-if="creatingPlanId === plan.id" class="h-4 w-4 animate-spin" />
              创建订单
            </Button>
          </CardSection>
        </Card>
      </div>

      <Card>
        <CardSection role="header">
          <div class="text-lg font-semibold">最近订单</div>
          <div class="text-sm text-zinc-500">显示最近 10 条订单</div>
        </CardSection>
        <CardSection>
          <Table>
            <thead class="border-b border-zinc-200 text-left text-xs uppercase text-zinc-500 dark:border-zinc-800">
              <tr>
                <th class="px-3 py-3">订单号</th>
                <th class="px-3 py-3">金额</th>
                <th class="px-3 py-3">状态</th>
                <th class="px-3 py-3">创建时间</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-zinc-200 dark:divide-zinc-800">
              <tr v-for="order in orders" :key="order.id">
                <td class="px-3 py-4 font-mono text-xs">{{ order.out_trade_no }}</td>
                <td class="px-3 py-4 text-sm">{{ formatCurrency(order.amount) }}</td>
                <td class="px-3 py-4"><Badge variant="secondary">{{ order.status }}</Badge></td>
                <td class="px-3 py-4 text-sm text-zinc-500">{{ formatDate(order.created_at) }}</td>
              </tr>
            </tbody>
          </Table>
          <div v-if="!orders.length" class="py-12 text-center text-sm text-zinc-500">
            暂无订单。
          </div>
        </CardSection>
      </Card>
    </div>
  </AppShellLayout>
</template>
