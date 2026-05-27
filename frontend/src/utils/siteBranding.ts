export const DEFAULT_SITE_SUBTITLE = 'Metered OpenAI API access with transparent balance'

const LEGACY_SITE_SUBTITLES = new Set([
  'Subscription to API Conversion Platform',
  'Subscription to API',
  '订阅转 API 转换平台',
  'AI API Gateway Platform',
  'Buy OpenAI tokens and call GPT models directly',
  '充值 OpenAI Token，直接调用 GPT 模型'
])

export function resolveSiteSubtitle(value: string | null | undefined, fallback = DEFAULT_SITE_SUBTITLE): string {
  const subtitle = value?.trim() ?? ''
  if (!subtitle || LEGACY_SITE_SUBTITLES.has(subtitle)) {
    return fallback
  }
  return subtitle
}
