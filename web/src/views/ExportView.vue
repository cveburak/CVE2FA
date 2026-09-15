<template>
  <div class="max-w-lg mx-auto mt-8">
    <div class="card bg-white border border-base-300 shadow-sm">
      <div class="card-body">
        <h2 class="card-title text-lg text-secondary mb-1">
          <div class="flex items-center justify-center h-9 w-9 rounded-lg bg-primary/10 text-primary">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/><polyline points="7 10 12 15 17 10"/><line x1="12" y1="15" x2="12" y2="3"/></svg>
          </div>
          Secure Export
        </h2>
        <p class="text-sm text-neutral mb-5">
          Export all your TOTP accounts as an encrypted CVE2fa file.
          You will need the master password to import it later.
        </p>

        <div v-if="error" class="alert alert-error text-sm mb-4 animate-slide-down">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><line x1="15" y1="9" x2="9" y2="15"/><line x1="9" y1="9" x2="15" y2="15"/></svg>
          <span>{{ error }}</span>
        </div>

        <div v-if="success" class="alert alert-success text-sm mb-4 animate-slide-down">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polyline points="20 6 9 17 4 12"/></svg>
          <span>{{ success }}</span>
        </div>

        <form @submit.prevent="handleExport">
          <div class="form-control mb-5">
            <label class="label"><span class="label-text font-medium">Master Password</span></label>
            <input v-model="masterPassword" type="password" class="input input-bordered bg-white transition-all duration-200 focus:ring-2 focus:ring-primary/20 focus:border-primary" required minlength="4" />
            <label class="label"><span class="label-text-alt text-xs">Used to encrypt your data. Do not forget this!</span></label>
          </div>

          <button type="submit" class="btn btn-primary w-full transition-all duration-200 hover:shadow-md hover:shadow-primary/20 active:scale-[0.98]" :disabled="loading">
            <span v-if="loading" class="loading loading-spinner"></span>
            {{ loading ? 'Encrypting...' : 'Download Encrypted CVE2fa' }}
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const masterPassword = ref('')
const loading = ref(false)
const error = ref(null)
const success = ref(null)

async function handleExport() {
  error.value = null
  success.value = null
  loading.value = true

  try {
    const res = await fetch('/api/export', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ masterPassword: masterPassword.value }),
    })

    if (!res.ok) {
      const data = await res.json()
      throw new Error(data.error || 'Export failed')
    }

    const blob = await res.blob()
    const url = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = 'CVE2fa.2fadata'
    a.click()
    URL.revokeObjectURL(url)

    success.value = 'CVE2fa exported successfully!'
    masterPassword.value = ''
  } catch (e) {
    error.value = e.message
  } finally {
    loading.value = false
  }
}
</script>