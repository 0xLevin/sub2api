<template>
  <div v-if="homeContent" class="min-h-screen">
    <iframe
      v-if="isHomeContentUrl"
      :src="homeContent.trim()"
      class="h-screen w-full border-0"
      allowfullscreen
    ></iframe>
    <div v-else v-html="homeContent"></div>
  </div>

  <div v-else class="public-home-shell">
    <div class="public-grid-bg"></div>

    <header class="public-home-header">
      <router-link to="/home" class="public-home-brand">
        <span class="public-brand-mark">
          <img :src="siteLogo || '/logo.png'" alt="" />
        </span>
        <span class="public-brand-copy">
          <span class="public-brand-name">{{ siteName }}</span>
          <span class="public-brand-subtitle">{{ siteSubtitle }}</span>
        </span>
      </router-link>

      <div class="public-home-actions">
        <LocaleSwitcher />
        <a
          v-if="docUrl"
          :href="docUrl"
          target="_blank"
          rel="noopener noreferrer"
          class="public-icon-button"
          :title="t('home.viewDocs')"
        >
          <Icon name="book" size="md" />
        </a>
        <button
          class="public-icon-button"
          type="button"
          :title="isDark ? t('home.switchToLight') : t('home.switchToDark')"
          @click="toggleTheme"
        >
          <Icon v-if="isDark" name="sun" size="md" />
          <Icon v-else name="moon" size="md" />
        </button>
        <router-link :to="isAuthenticated ? dashboardPath : '/login'" class="public-nav-button">
          {{ isAuthenticated ? t('home.dashboard') : t('home.login') }}
        </router-link>
      </div>
    </header>

    <main class="public-home-main">
      <section class="public-home-hero">
        <div class="public-kicker">{{ t('home.v2.kicker') }}</div>
        <h1>{{ siteName }}</h1>
        <p>{{ siteSubtitle }}</p>
        <p class="public-hero-description">{{ t('home.heroDescription') }}</p>
        <div class="public-home-cta">
          <router-link :to="isAuthenticated ? dashboardPath : '/register'" class="public-primary-button">
            {{ isAuthenticated ? t('home.goToDashboard') : t('home.getStarted') }}
            <Icon name="arrowRight" size="md" />
          </router-link>
          <a
            v-if="docUrl"
            :href="docUrl"
            target="_blank"
            rel="noopener noreferrer"
            class="public-secondary-button"
          >
            {{ t('home.docs') }}
          </a>
        </div>
      </section>

      <section class="public-terminal-card" aria-label="API terminal preview">
        <div class="public-terminal-bar">
          <span></span>
          <span>openai.token</span>
        </div>
        <div class="public-terminal-body">
          <p><b>$</b> curl -X POST /v1/responses</p>
          <p><i>key</i> user balance verified</p>
          <p><i>usage</i> tokens deducted in realtime</p>
          <p><strong>200 OK</strong> { "content": "Hello!" }</p>
        </div>
      </section>

      <section class="public-feature-strip">
        <article>
          <Icon name="swap" size="lg" />
          <h2>{{ t('home.features.unifiedGateway') }}</h2>
          <p>{{ t('home.features.unifiedGatewayDesc') }}</p>
        </article>
        <article>
          <Icon name="shield" size="lg" />
          <h2>{{ t('home.features.multiAccount') }}</h2>
          <p>{{ t('home.features.multiAccountDesc') }}</p>
        </article>
        <article>
          <Icon name="chart" size="lg" />
          <h2>{{ t('home.features.balanceQuota') }}</h2>
          <p>{{ t('home.features.balanceQuotaDesc') }}</p>
        </article>
      </section>

      <section class="public-buy-section">
        <div class="public-section-heading">
          <span>{{ t('home.v2.buying.kicker') }}</span>
          <h2>{{ t('home.v2.buying.title') }}</h2>
          <p>{{ t('home.v2.buying.description') }}</p>
        </div>

        <div class="public-buy-grid">
          <article>
            <Icon name="creditCard" size="lg" />
            <span>{{ t('home.v2.buying.items.topUp.label') }}</span>
            <h3>{{ t('home.v2.buying.items.topUp.title') }}</h3>
            <p>{{ t('home.v2.buying.items.topUp.description') }}</p>
          </article>
          <article>
            <Icon name="key" size="lg" />
            <span>{{ t('home.v2.buying.items.key.label') }}</span>
            <h3>{{ t('home.v2.buying.items.key.title') }}</h3>
            <p>{{ t('home.v2.buying.items.key.description') }}</p>
          </article>
          <article>
            <Icon name="chartBar" size="lg" />
            <span>{{ t('home.v2.buying.items.usage.label') }}</span>
            <h3>{{ t('home.v2.buying.items.usage.title') }}</h3>
            <p>{{ t('home.v2.buying.items.usage.description') }}</p>
          </article>
        </div>
      </section>

      <section class="public-capability-section">
        <div class="public-section-heading">
          <span>{{ t('home.v2.capabilities.kicker') }}</span>
          <h2>{{ t('home.v2.capabilities.title') }}</h2>
          <p>{{ t('home.v2.capabilities.description') }}</p>
        </div>

        <div class="public-capability-grid">
          <article>
            <Icon name="terminal" size="lg" />
            <h3>{{ t('home.v2.capabilities.items.api.title') }}</h3>
            <p>{{ t('home.v2.capabilities.items.api.description') }}</p>
          </article>
          <article>
            <Icon name="server" size="lg" />
            <h3>{{ t('home.v2.capabilities.items.routing.title') }}</h3>
            <p>{{ t('home.v2.capabilities.items.routing.description') }}</p>
          </article>
          <article>
            <Icon name="database" size="lg" />
            <h3>{{ t('home.v2.capabilities.items.accountPool.title') }}</h3>
            <p>{{ t('home.v2.capabilities.items.accountPool.description') }}</p>
          </article>
          <article>
            <Icon name="calculator" size="lg" />
            <h3>{{ t('home.v2.capabilities.items.billing.title') }}</h3>
            <p>{{ t('home.v2.capabilities.items.billing.description') }}</p>
          </article>
          <article>
            <Icon name="shield" size="lg" />
            <h3>{{ t('home.v2.capabilities.items.guardrails.title') }}</h3>
            <p>{{ t('home.v2.capabilities.items.guardrails.description') }}</p>
          </article>
          <article>
            <Icon name="chartBar" size="lg" />
            <h3>{{ t('home.v2.capabilities.items.observability.title') }}</h3>
            <p>{{ t('home.v2.capabilities.items.observability.description') }}</p>
          </article>
        </div>
      </section>

      <section class="public-workflow-section">
        <div class="public-section-heading">
          <span>{{ t('home.v2.workflow.kicker') }}</span>
          <h2>{{ t('home.v2.workflow.title') }}</h2>
          <p>{{ t('home.v2.workflow.description') }}</p>
        </div>

        <div class="public-workflow-grid">
          <article>
            <b>01</b>
            <h3>{{ t('home.v2.workflow.steps.connect.title') }}</h3>
            <p>{{ t('home.v2.workflow.steps.connect.description') }}</p>
          </article>
          <article>
            <b>02</b>
            <h3>{{ t('home.v2.workflow.steps.issue.title') }}</h3>
            <p>{{ t('home.v2.workflow.steps.issue.description') }}</p>
          </article>
          <article>
            <b>03</b>
            <h3>{{ t('home.v2.workflow.steps.route.title') }}</h3>
            <p>{{ t('home.v2.workflow.steps.route.description') }}</p>
          </article>
          <article>
            <b>04</b>
            <h3>{{ t('home.v2.workflow.steps.measure.title') }}</h3>
            <p>{{ t('home.v2.workflow.steps.measure.description') }}</p>
          </article>
        </div>
      </section>

      <section class="public-model-section">
        <div class="public-model-copy">
          <span>{{ t('home.v2.modelAccess.kicker') }}</span>
          <h2>{{ t('home.v2.modelAccess.title') }}</h2>
          <p>{{ t('home.v2.modelAccess.description') }}</p>
        </div>
        <div class="public-model-list">
          <article>
            <strong>{{ t('home.v2.modelAccess.items.chat.title') }}</strong>
            <p>{{ t('home.v2.modelAccess.items.chat.description') }}</p>
          </article>
          <article>
            <strong>{{ t('home.v2.modelAccess.items.code.title') }}</strong>
            <p>{{ t('home.v2.modelAccess.items.code.description') }}</p>
          </article>
          <article>
            <strong>{{ t('home.v2.modelAccess.items.automation.title') }}</strong>
            <p>{{ t('home.v2.modelAccess.items.automation.description') }}</p>
          </article>
        </div>
      </section>

      <section class="public-control-section">
        <div class="public-control-panel">
          <div class="public-panel-header">
            <span>{{ t('home.v2.controlPlane.kicker') }}</span>
            <strong>token.console</strong>
          </div>
          <div class="public-metric-grid">
            <div>
              <span>{{ t('home.v2.controlPlane.metrics.access') }}</span>
              <strong>API</strong>
            </div>
            <div>
              <span>{{ t('home.v2.controlPlane.metrics.models') }}</span>
              <strong>OPENAI</strong>
            </div>
            <div>
              <span>{{ t('home.v2.controlPlane.metrics.usage') }}</span>
              <strong>LIVE</strong>
            </div>
            <div>
              <span>{{ t('home.v2.controlPlane.metrics.balance') }}</span>
              <strong>TOKEN</strong>
            </div>
          </div>
          <div class="public-signal-list">
            <p><i></i>{{ t('home.v2.controlPlane.signals.failover') }}</p>
            <p><i></i>{{ t('home.v2.controlPlane.signals.usage') }}</p>
            <p><i></i>{{ t('home.v2.controlPlane.signals.audit') }}</p>
          </div>
        </div>

        <div class="public-usecase-panel">
          <span>{{ t('home.v2.useCases.kicker') }}</span>
          <h2>{{ t('home.v2.useCases.title') }}</h2>
          <div class="public-usecase-list">
            <article>
              <Icon name="users" size="md" />
              <div>
                <h3>{{ t('home.v2.useCases.items.teams.title') }}</h3>
                <p>{{ t('home.v2.useCases.items.teams.description') }}</p>
              </div>
            </article>
            <article>
              <Icon name="cpu" size="md" />
              <div>
                <h3>{{ t('home.v2.useCases.items.products.title') }}</h3>
                <p>{{ t('home.v2.useCases.items.products.description') }}</p>
              </div>
            </article>
            <article>
              <Icon name="key" size="md" />
              <div>
                <h3>{{ t('home.v2.useCases.items.resellers.title') }}</h3>
                <p>{{ t('home.v2.useCases.items.resellers.description') }}</p>
              </div>
            </article>
          </div>
        </div>
      </section>

      <section class="public-faq-section">
        <div class="public-section-heading">
          <span>{{ t('home.v2.faq.kicker') }}</span>
          <h2>{{ t('home.v2.faq.title') }}</h2>
        </div>
        <div class="public-faq-grid">
          <article>
            <h3>{{ t('home.v2.faq.items.billing.question') }}</h3>
            <p>{{ t('home.v2.faq.items.billing.answer') }}</p>
          </article>
          <article>
            <h3>{{ t('home.v2.faq.items.compatibility.question') }}</h3>
            <p>{{ t('home.v2.faq.items.compatibility.answer') }}</p>
          </article>
          <article>
            <h3>{{ t('home.v2.faq.items.balance.question') }}</h3>
            <p>{{ t('home.v2.faq.items.balance.answer') }}</p>
          </article>
          <article>
            <h3>{{ t('home.v2.faq.items.keys.question') }}</h3>
            <p>{{ t('home.v2.faq.items.keys.answer') }}</p>
          </article>
        </div>
      </section>

      <section class="public-final-cta">
        <div>
          <span>{{ t('home.v2.finalCta.kicker') }}</span>
          <h2>{{ t('home.cta.title') }}</h2>
          <p>{{ t('home.cta.description') }}</p>
        </div>
        <router-link :to="isAuthenticated ? dashboardPath : '/register'" class="public-primary-button">
          {{ isAuthenticated ? t('home.goToDashboard') : t('home.cta.button') }}
          <Icon name="arrowRight" size="md" />
        </router-link>
      </section>
    </main>

    <footer class="public-home-footer">
      <span>&copy; {{ currentYear }} {{ siteName }}. {{ t('home.footer.allRightsReserved') }}</span>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useAuthStore, useAppStore } from '@/stores'
import LocaleSwitcher from '@/components/common/LocaleSwitcher.vue'
import Icon from '@/components/icons/Icon.vue'
import { resolveSiteSubtitle } from '@/utils/siteBranding'

const { t } = useI18n()
const authStore = useAuthStore()
const appStore = useAppStore()

const siteName = computed(() => appStore.cachedPublicSettings?.site_name || appStore.siteName || 'Sub2API')
const siteLogo = computed(() => appStore.cachedPublicSettings?.site_logo || appStore.siteLogo || '')
const siteSubtitle = computed(() => resolveSiteSubtitle(appStore.cachedPublicSettings?.site_subtitle, t('home.heroSubtitle')))
const docUrl = computed(() => appStore.cachedPublicSettings?.doc_url || appStore.docUrl || '')
const homeContent = computed(() => appStore.cachedPublicSettings?.home_content || '')
const isHomeContentUrl = computed(() => {
  const content = homeContent.value.trim()
  return content.startsWith('http://') || content.startsWith('https://')
})

const isDark = ref(document.documentElement.classList.contains('dark'))
const isAuthenticated = computed(() => authStore.isAuthenticated)
const isAdmin = computed(() => authStore.isAdmin)
const dashboardPath = computed(() => (isAdmin.value ? '/admin/dashboard' : '/dashboard'))
const currentYear = computed(() => new Date().getFullYear())

function toggleTheme() {
  isDark.value = !isDark.value
  document.documentElement.classList.toggle('dark', isDark.value)
  localStorage.setItem('theme', isDark.value ? 'dark' : 'light')
}

function initTheme() {
  const savedTheme = localStorage.getItem('theme')
  if (
    savedTheme === 'dark' ||
    (!savedTheme && window.matchMedia('(prefers-color-scheme: dark)').matches)
  ) {
    isDark.value = true
    document.documentElement.classList.add('dark')
  }
}

onMounted(() => {
  initTheme()
  authStore.checkAuth()
  if (!appStore.publicSettingsLoaded) {
    appStore.fetchPublicSettings()
  }
})
</script>

<style scoped>
.public-home-main {
  position: relative;
  z-index: 1;
  display: grid;
  grid-template-columns: minmax(0, 1fr) minmax(20rem, 31rem);
  gap: clamp(2rem, 6vw, 5rem);
  align-items: center;
  padding: clamp(2rem, 7vw, 6rem) clamp(1rem, 5vw, 5rem) 2rem;
}

.public-home-hero h1 {
  margin-top: 1rem;
  max-width: 11ch;
  font-size: clamp(3.8rem, 10vw, 9rem);
  font-weight: 800;
  line-height: 0.86;
  text-transform: uppercase;
}

.public-home-hero p {
  margin-top: 1.25rem;
  max-width: 42rem;
  font-size: clamp(1rem, 2vw, 1.4rem);
  color: color-mix(in srgb, var(--foreground) 76%, transparent);
}

.public-home-hero .public-hero-description {
  margin-top: 0.85rem;
  max-width: 38rem;
  font-size: 1rem;
  color: color-mix(in srgb, var(--foreground) 66%, transparent);
}

.public-home-cta {
  margin-top: 2rem;
  display: flex;
  flex-wrap: wrap;
  gap: 0.875rem;
}

.public-icon-button,
.public-nav-button,
.public-primary-button,
.public-secondary-button {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  border: 1px solid var(--border);
  background: var(--card);
  color: var(--foreground);
  text-decoration: none;
  transition:
    transform 0.18s ease,
    box-shadow 0.18s ease;
}

.public-icon-button {
  height: 2.4rem;
  width: 2.4rem;
}

.public-nav-button,
.public-primary-button,
.public-secondary-button {
  min-height: 2.7rem;
  padding: 0.65rem 1rem;
  font-weight: 800;
  text-transform: uppercase;
}

.public-primary-button,
.public-nav-button {
  background: var(--primary);
  color: var(--primary-foreground);
  box-shadow: 4px 4px 0 color-mix(in srgb, var(--border) 22%, transparent);
}

.public-icon-button:hover,
.public-nav-button:hover,
.public-primary-button:hover,
.public-secondary-button:hover {
  transform: translate(-2px, -2px);
  box-shadow: 5px 5px 0 color-mix(in srgb, var(--border) 22%, transparent);
}

.public-terminal-card {
  border: 1px solid var(--border);
  background: #11130f;
  color: #f8f0dd;
  box-shadow: var(--public-shadow);
}

.public-terminal-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-bottom: 1px solid #f8f0dd;
  padding: 0.75rem 1rem;
  font-family: var(--font-mono-public);
  font-size: 0.75rem;
}

.public-terminal-bar span:first-child {
  height: 0.7rem;
  width: 0.7rem;
  background: var(--accent);
}

.public-terminal-body {
  display: grid;
  gap: 0.85rem;
  padding: 1.25rem;
  font-family: var(--font-mono-public);
  font-size: clamp(0.78rem, 1.4vw, 0.95rem);
}

.public-terminal-body b {
  color: var(--primary);
}

.public-terminal-body i {
  color: var(--accent);
  font-style: normal;
}

.public-terminal-body strong {
  color: #75e6a9;
}

.public-feature-strip {
  grid-column: 1 / -1;
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 1rem;
  margin-top: clamp(2rem, 5vw, 4rem);
}

.public-feature-strip article {
  border: 1px solid var(--border);
  background: color-mix(in srgb, var(--card) 90%, transparent);
  padding: 1.25rem;
  box-shadow: 4px 4px 0 color-mix(in srgb, var(--border) 14%, transparent);
}

.public-feature-strip h2 {
  margin-top: 0.85rem;
  font-size: 1rem;
  font-weight: 800;
}

.public-feature-strip p {
  margin-top: 0.5rem;
  font-size: 0.9rem;
  color: color-mix(in srgb, var(--foreground) 70%, transparent);
}

.public-capability-section,
.public-workflow-section,
.public-control-section,
.public-buy-section,
.public-model-section,
.public-faq-section,
.public-final-cta {
  grid-column: 1 / -1;
  margin-top: clamp(2rem, 5vw, 4rem);
}

.public-section-heading {
  max-width: 52rem;
}

.public-section-heading span,
.public-model-copy > span,
.public-usecase-panel > span,
.public-final-cta span {
  display: inline-flex;
  border-left: 3px solid var(--primary);
  padding-left: 0.75rem;
  color: var(--primary);
  font-family: var(--font-mono-public);
  font-size: 0.78rem;
  font-weight: 800;
  letter-spacing: 0;
  text-transform: uppercase;
}

.public-section-heading h2,
.public-model-copy h2,
.public-usecase-panel h2,
.public-final-cta h2 {
  margin-top: 0.8rem;
  font-size: clamp(1.8rem, 4vw, 3.8rem);
  font-weight: 800;
  line-height: 0.95;
  text-transform: uppercase;
}

.public-section-heading p,
.public-model-copy p,
.public-final-cta p {
  margin-top: 0.85rem;
  max-width: 42rem;
  color: color-mix(in srgb, var(--foreground) 70%, transparent);
}

.public-buy-grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 1rem;
  margin-top: 1.5rem;
}

.public-buy-grid article {
  border: 1px solid var(--border);
  background: var(--card);
  padding: 1.25rem;
  box-shadow: 4px 4px 0 color-mix(in srgb, var(--border) 14%, transparent);
}

.public-buy-grid svg {
  color: var(--primary);
}

.public-buy-grid span {
  display: block;
  margin-top: 1rem;
  color: var(--accent);
  font-family: var(--font-mono-public);
  font-size: 0.72rem;
  font-weight: 800;
  text-transform: uppercase;
}

.public-buy-grid h3 {
  margin-top: 0.5rem;
  font-size: 1.2rem;
  font-weight: 800;
  text-transform: uppercase;
}

.public-buy-grid p {
  margin-top: 0.65rem;
  color: color-mix(in srgb, var(--foreground) 68%, transparent);
  font-size: 0.92rem;
}

.public-capability-grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 1rem;
  margin-top: 1.5rem;
}

.public-capability-grid article,
.public-workflow-grid article,
.public-usecase-list article,
.public-control-panel,
.public-usecase-panel,
.public-final-cta {
  border: 1px solid var(--border);
  background: color-mix(in srgb, var(--card) 92%, transparent);
  box-shadow: 4px 4px 0 color-mix(in srgb, var(--border) 14%, transparent);
}

.public-capability-grid article {
  min-height: 13rem;
  padding: 1.25rem;
}

.public-capability-grid svg,
.public-feature-strip svg,
.public-usecase-list svg {
  color: var(--primary);
}

.public-capability-grid h3,
.public-workflow-grid h3,
.public-usecase-list h3 {
  margin-top: 1rem;
  font-size: 1rem;
  font-weight: 800;
  text-transform: uppercase;
}

.public-capability-grid p,
.public-workflow-grid p,
.public-usecase-list p {
  margin-top: 0.55rem;
  color: color-mix(in srgb, var(--foreground) 68%, transparent);
  font-size: 0.9rem;
}

.public-workflow-grid {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 1rem;
  margin-top: 1.5rem;
}

.public-workflow-grid article {
  position: relative;
  min-height: 12rem;
  padding: 1.25rem;
}

.public-workflow-grid b {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  height: 2.3rem;
  min-width: 2.9rem;
  border: 1px solid var(--border);
  background: var(--primary);
  color: var(--primary-foreground);
  font-family: var(--font-mono-public);
  font-size: 0.85rem;
}

.public-model-section {
  display: grid;
  grid-template-columns: minmax(0, 0.85fr) minmax(0, 1.15fr);
  gap: 1rem;
  align-items: stretch;
}

.public-model-copy,
.public-model-list article,
.public-faq-grid article {
  border: 1px solid var(--border);
  background: color-mix(in srgb, var(--card) 92%, transparent);
  box-shadow: 4px 4px 0 color-mix(in srgb, var(--border) 14%, transparent);
}

.public-model-copy {
  padding: 1.25rem;
}

.public-model-list {
  display: grid;
  gap: 1rem;
}

.public-model-list article {
  padding: 1.1rem;
}

.public-model-list strong {
  font-size: 1rem;
  text-transform: uppercase;
}

.public-model-list p {
  margin-top: 0.5rem;
  color: color-mix(in srgb, var(--foreground) 68%, transparent);
  font-family: var(--font-mono-public);
  font-size: 0.86rem;
}

.public-control-section {
  display: grid;
  grid-template-columns: minmax(0, 1.1fr) minmax(19rem, 0.9fr);
  gap: 1rem;
  align-items: stretch;
}

.public-control-panel,
.public-usecase-panel {
  padding: 1.25rem;
}

.public-panel-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-bottom: 1px solid color-mix(in srgb, var(--border) 38%, transparent);
  padding-bottom: 0.85rem;
  font-family: var(--font-mono-public);
}

.public-panel-header span {
  color: var(--primary);
  font-size: 0.78rem;
  font-weight: 800;
  text-transform: uppercase;
}

.public-panel-header strong {
  color: color-mix(in srgb, var(--foreground) 72%, transparent);
  font-size: 0.82rem;
}

.public-metric-grid {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 0.75rem;
  margin-top: 1rem;
}

.public-metric-grid div {
  border: 1px solid color-mix(in srgb, var(--border) 34%, transparent);
  background: var(--surface);
  padding: 1rem;
}

.public-metric-grid span,
.public-signal-list,
.public-usecase-list p {
  font-family: var(--font-mono-public);
}

.public-metric-grid span {
  display: block;
  color: var(--text-muted);
  font-size: 0.72rem;
  text-transform: uppercase;
}

.public-metric-grid strong {
  display: block;
  margin-top: 0.5rem;
  font-size: clamp(1.1rem, 2vw, 1.75rem);
  color: var(--foreground);
}

.public-signal-list {
  display: grid;
  gap: 0.65rem;
  margin-top: 1rem;
  color: color-mix(in srgb, var(--foreground) 72%, transparent);
  font-size: 0.86rem;
}

.public-signal-list p {
  display: flex;
  align-items: center;
  gap: 0.65rem;
  margin: 0;
}

.public-signal-list i {
  height: 0.65rem;
  width: 0.65rem;
  background: var(--primary);
  box-shadow: 0 0 0 3px color-mix(in srgb, var(--primary) 16%, transparent);
}

.public-usecase-list {
  display: grid;
  gap: 0.75rem;
  margin-top: 1.25rem;
}

.public-usecase-list article {
  display: flex;
  gap: 0.9rem;
  padding: 1rem;
  box-shadow: none;
}

.public-usecase-list h3 {
  margin-top: 0;
}

.public-faq-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 1rem;
  margin-top: 1.5rem;
}

.public-faq-grid article {
  padding: 1.25rem;
}

.public-faq-grid h3 {
  font-size: 1rem;
  font-weight: 800;
  text-transform: uppercase;
}

.public-faq-grid p {
  margin-top: 0.55rem;
  color: color-mix(in srgb, var(--foreground) 68%, transparent);
}

.public-final-cta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1.5rem;
  padding: clamp(1.25rem, 4vw, 2rem);
  background:
    linear-gradient(90deg, color-mix(in srgb, var(--primary) 10%, var(--card)), var(--card)) !important;
}

.public-home-footer {
  position: relative;
  z-index: 1;
  display: flex;
  justify-content: center;
  gap: 1rem;
  padding: 1.5rem;
  color: color-mix(in srgb, var(--foreground) 68%, transparent);
  font-size: 0.85rem;
}

.public-home-footer a {
  color: inherit;
}

@media (max-width: 900px) {
  .public-home-shell {
    overflow: auto;
  }

  .public-home-main {
    grid-template-columns: 1fr;
  }

  .public-feature-strip {
    grid-template-columns: 1fr;
  }

  .public-capability-grid,
  .public-workflow-grid,
  .public-control-section,
  .public-buy-grid,
  .public-model-section,
  .public-faq-grid,
  .public-metric-grid {
    grid-template-columns: 1fr;
  }

  .public-final-cta {
    align-items: flex-start;
    flex-direction: column;
  }
}

@media (max-width: 560px) {
  .public-home-actions {
    gap: 0.35rem;
  }

  .public-nav-button {
    padding-inline: 0.7rem;
    font-size: 0.78rem;
  }
}
</style>
