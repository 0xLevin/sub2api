<script setup lang="ts">
import { computed } from 'vue'
import { RouterLink } from 'vue-router'
import { ArrowRight, BarChart3, CheckCircle2, KeyRound, Layers3, ShieldCheck, WalletCards } from 'lucide-vue-next'
import { useAppStore } from '@/stores'
import { Badge, Button, Card, CardSection } from '@/brand/ui'

const appStore = useAppStore()
const siteName = computed(() => appStore.siteName || 'API Gateway')

const features = [
  { title: '统一 API Key', desc: '把多模型、多账号、多套餐收束为一个稳定入口。', icon: KeyRound },
  { title: '售卖闭环', desc: '注册、购买、创建 Key、查用量，面向客户完成重设计。', icon: WalletCards },
  { title: '用量透明', desc: '请求、Token、成本、订单状态集中展示，减少售后沟通。', icon: BarChart3 },
  { title: '稳定调度', desc: '保留原有账号池、限额和路由能力，前台只重塑体验。', icon: ShieldCheck }
]

const steps = ['选择套餐', '注册并完成支付', '创建 API Key', '接入 OpenAI 兼容接口']
</script>

<template>
  <MarketingLayout>
    <section class="border-b border-zinc-200 dark:border-zinc-800">
      <div class="mx-auto grid max-w-7xl gap-12 px-4 py-20 sm:px-6 lg:grid-cols-[1.05fr_0.95fr] lg:px-8 lg:py-24">
        <div class="flex flex-col justify-center">
          <Badge variant="outline" class="w-fit">AI API Commercial Gateway</Badge>
          <h1 class="mt-6 max-w-3xl text-5xl font-semibold tracking-tight text-zinc-950 dark:text-white md:text-6xl">
            {{ siteName }}，为对外售卖重新设计的 AI API 控制台。
          </h1>
          <p class="mt-6 max-w-2xl text-lg leading-8 text-zinc-600 dark:text-zinc-300">
            将模型接入、套餐购买、API Key 管理和用量查询放在同一条清晰路径里。客户无需理解后台系统，只需要购买、复制 Key、开始调用。
          </p>
          <div class="mt-8 flex flex-col gap-3 sm:flex-row">
            <RouterLink to="/pricing">
              <Button size="lg">查看套餐 <ArrowRight class="h-4 w-4" /></Button>
            </RouterLink>
            <RouterLink to="/login">
              <Button variant="outline" size="lg">已有账户登录</Button>
            </RouterLink>
          </div>
        </div>
        <div class="relative">
          <div class="absolute inset-0 rounded-3xl bg-zinc-900 blur-3xl opacity-10 dark:bg-white" />
          <Card class="relative overflow-hidden rounded-2xl">
            <CardSection role="header">
              <div class="flex items-center justify-between">
                <div>
                  <div class="text-sm font-medium text-zinc-500">Customer Console</div>
                  <div class="mt-1 text-2xl font-semibold">一站式交付</div>
                </div>
                <Badge variant="success">Live</Badge>
              </div>
            </CardSection>
            <CardSection>
              <div class="grid gap-3">
                <div class="rounded-lg border border-zinc-200 p-4 dark:border-zinc-800">
                  <div class="text-sm text-zinc-500">今日请求</div>
                  <div class="mt-2 text-3xl font-semibold">128,400</div>
                </div>
                <div class="grid grid-cols-2 gap-3">
                  <div class="rounded-lg bg-zinc-100 p-4 dark:bg-zinc-900">
                    <div class="text-xs text-zinc-500">Token</div>
                    <div class="mt-2 text-xl font-semibold">42.8M</div>
                  </div>
                  <div class="rounded-lg bg-zinc-100 p-4 dark:bg-zinc-900">
                    <div class="text-xs text-zinc-500">余额</div>
                    <div class="mt-2 text-xl font-semibold">$238.60</div>
                  </div>
                </div>
                <div class="rounded-lg border border-zinc-200 p-4 dark:border-zinc-800">
                  <div class="mb-3 flex items-center justify-between text-sm">
                    <span>API Key</span>
                    <span class="text-zinc-500">active</span>
                  </div>
                  <div class="rounded-md bg-zinc-950 px-3 py-2 font-mono text-xs text-zinc-100 dark:bg-zinc-900">
                    sk-live-••••••••••••••••••••
                  </div>
                </div>
              </div>
            </CardSection>
          </Card>
        </div>
      </div>
    </section>

    <section class="mx-auto max-w-7xl px-4 py-16 sm:px-6 lg:px-8">
      <div class="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        <Card v-for="feature in features" :key="feature.title">
          <CardSection role="header">
            <component :is="feature.icon" class="h-5 w-5" />
            <div class="text-lg font-semibold">{{ feature.title }}</div>
          </CardSection>
          <CardSection>
            <p class="text-sm leading-6 text-zinc-600 dark:text-zinc-300">{{ feature.desc }}</p>
          </CardSection>
        </Card>
      </div>
    </section>

    <section class="bg-zinc-50 py-16 dark:bg-zinc-900/40">
      <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div class="grid gap-10 lg:grid-cols-[0.8fr_1.2fr]">
          <div>
            <Badge variant="secondary">接入流程</Badge>
            <h2 class="mt-4 text-3xl font-semibold tracking-tight">从购买到调用，压缩到四步。</h2>
            <p class="mt-4 text-zinc-600 dark:text-zinc-300">新前台把客户最关心的路径放到第一层，后台运维能力继续留在管理端。</p>
          </div>
          <div class="grid gap-3 sm:grid-cols-2">
            <div v-for="(step, index) in steps" :key="step" class="rounded-lg border border-zinc-200 bg-white p-5 dark:border-zinc-800 dark:bg-zinc-950">
              <div class="flex items-center gap-3">
                <div class="flex h-8 w-8 items-center justify-center rounded-full bg-zinc-950 text-sm text-white dark:bg-white dark:text-zinc-950">{{ index + 1 }}</div>
                <div class="font-medium">{{ step }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <section class="mx-auto max-w-7xl px-4 py-16 sm:px-6 lg:px-8">
      <div class="rounded-2xl bg-zinc-950 p-8 text-white md:p-12">
        <div class="grid gap-8 md:grid-cols-[1fr_auto] md:items-center">
          <div>
            <div class="flex items-center gap-2 text-sm text-zinc-300"><Layers3 class="h-4 w-4" /> OpenAI-compatible</div>
            <h2 class="mt-3 text-3xl font-semibold">准备好开始售卖 AI API 了吗？</h2>
            <p class="mt-3 max-w-2xl text-zinc-300">先用默认品牌信息上线，后续可以替换文案、价格和视觉资产。</p>
          </div>
          <RouterLink to="/register">
            <Button variant="secondary" size="lg">创建账户 <CheckCircle2 class="h-4 w-4" /></Button>
          </RouterLink>
        </div>
      </div>
    </section>
  </MarketingLayout>
</template>
