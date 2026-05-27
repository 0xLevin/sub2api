<template>
  <div class="public-auth-shell">
    <div class="public-grid-bg"></div>

    <header class="public-auth-header">
      <router-link to="/home" class="public-brand-link" aria-label="Home">
        <span class="public-brand-mark">
          <img :src="siteLogo || '/logo.png'" alt="" />
        </span>
        <span class="public-brand-copy">
          <span class="public-brand-name">{{ siteName }}</span>
          <span class="public-brand-subtitle">{{ siteSubtitle }}</span>
        </span>
      </router-link>
      <div class="public-auth-actions">
        <LocaleSwitcher />
      </div>
    </header>

    <main class="public-auth-main">
      <section class="public-auth-aside" aria-hidden="true">
        <div class="public-kicker">{{ t('auth.publicLayout.kicker') }}</div>
        <h1>{{ siteName }}</h1>
        <p>{{ siteSubtitle }}</p>
        <div class="public-metric-grid">
          <div>
            <span>{{ t('auth.publicLayout.metrics.api.label') }}</span>
            <strong>{{ t('auth.publicLayout.metrics.api.value') }}</strong>
          </div>
          <div>
            <span>{{ t('auth.publicLayout.metrics.auth.label') }}</span>
            <strong>{{ t('auth.publicLayout.metrics.auth.value') }}</strong>
          </div>
          <div>
            <span>{{ t('auth.publicLayout.metrics.mode.label') }}</span>
            <strong>{{ t('auth.publicLayout.metrics.mode.value') }}</strong>
          </div>
        </div>
      </section>

      <section class="public-auth-panel">
        <slot />
        <div class="public-auth-footer">
          <slot name="footer" />
        </div>
      </section>
    </main>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { useAppStore } from '@/stores'
import LocaleSwitcher from '@/components/common/LocaleSwitcher.vue'
import { sanitizeUrl } from '@/utils/url'
import { resolveSiteSubtitle } from '@/utils/siteBranding'

const { t } = useI18n()
const appStore = useAppStore()

const siteName = computed(() => appStore.siteName || appStore.cachedPublicSettings?.site_name || 'Sub2API')
const siteLogo = computed(() =>
  sanitizeUrl(appStore.siteLogo || appStore.cachedPublicSettings?.site_logo || '', {
    allowRelative: true,
    allowDataUrl: true
  })
)
const siteSubtitle = computed(
  () => resolveSiteSubtitle(appStore.cachedPublicSettings?.site_subtitle, t('home.heroSubtitle'))
)

onMounted(() => {
  if (!appStore.publicSettingsLoaded) {
    appStore.fetchPublicSettings()
  }
})
</script>
