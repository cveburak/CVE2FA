<template>
  <div class="card bg-white rounded-xl border border-base-300 shadow-sm transition-card hover:-translate-y-1 hover:shadow-lg group">
    <div class="card-body p-5">
      <div class="flex justify-between items-start mb-3">
        <div class="min-w-0">
          <h3 class="text-sm font-semibold text-secondary truncate">{{ account.issuer }}</h3>
          <p class="text-xs text-base-content/50 truncate">{{ account.accountName }}</p>
        </div>
        <button @click="$emit('delete', account.id)" class="btn btn-ghost btn-xs btn-circle text-base-content/30 transition-all duration-200 hover:bg-error/10 hover:text-error group-hover:opacity-100 opacity-60">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polyline points="3 6 5 6 21 6"/><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/></svg>
        </button>
      </div>

      <div class="text-center my-4">
        <span class="font-mono text-3xl font-semibold tracking-[0.25em] transition-colors duration-300" :class="codeColor">
          {{ code }}
        </span>
      </div>

      <div class="w-full bg-base-200 rounded-full h-1.5 mt-2 overflow-hidden">
        <div
          class="h-1.5 rounded-full transition-all duration-1000 ease-linear shadow-sm"
          :class="progressColor"
          :style="{ width: progress + '%' }"
        ></div>
      </div>

      <div class="flex justify-between text-[11px] text-base-content/40 mt-1.5">
        <span>{{ remaining }}s</span>
        <span>30s</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  account: { type: Object, required: true },
  code: { type: String, default: '------' },
  progress: { type: Number, default: 0 },
  remaining: { type: Number, default: 0 },
})

defineEmits(['delete'])

const codeColor = computed(() => {
  if (props.remaining <= 5) return 'text-error'
  if (props.remaining <= 10) return 'text-warning'
  return 'text-success'
})

const progressColor = computed(() => {
  if (props.remaining <= 5) return 'bg-error'
  if (props.remaining <= 10) return 'bg-warning'
  return 'bg-success'
})
</script>