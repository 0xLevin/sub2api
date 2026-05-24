import { readonly, ref } from 'vue'

export interface BrandToast {
  id: number
  title: string
  description?: string
  variant?: 'default' | 'destructive'
}

const toasts = ref<BrandToast[]>([])
let nextId = 1

export function useBrandToast() {
  function show(toast: Omit<BrandToast, 'id'>) {
    const id = nextId++
    toasts.value = [...toasts.value, { id, ...toast }]
    window.setTimeout(() => dismiss(id), 4000)
  }

  function dismiss(id: number) {
    toasts.value = toasts.value.filter((toast) => toast.id !== id)
  }

  return {
    toasts: readonly(toasts),
    show,
    dismiss
  }
}
