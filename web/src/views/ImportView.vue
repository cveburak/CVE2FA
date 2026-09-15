<template>
  <div class="max-w-lg mx-auto mt-8">
    <div class="card bg-white border border-base-300 shadow-sm">
      <div class="card-body">
        <h2 class="card-title text-lg text-secondary mb-1">
          <div class="flex items-center justify-center h-9 w-9 rounded-lg bg-primary/10 text-primary">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/><polyline points="17 8 12 3 7 8"/><line x1="12" y1="3" x2="12" y2="15"/></svg>
          </div>
          Secure Import
        </h2>
        <p class="text-sm text-neutral mb-5">
          Import accounts from an encrypted CVE2fa file (.2fadata).
          Duplicate accounts (same issuer + account name) will be skipped.
        </p>

        <div v-if="error" class="alert alert-error text-sm mb-4 animate-slide-down">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><line x1="15" y1="9" x2="9" y2="15"/><line x1="9" y1="9" x2="15" y2="15"/></svg>
          <span>{{ error }}</span>
        </div>

        <div v-if="result" class="alert alert-success text-sm mb-4 animate-slide-down">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polyline points="20 6 9 17 4 12"/></svg>
          <span>{{ result }}</span>
        </div>

        <form @submit.prevent="handleImport">
          <div class="form-control mb-3">
            <label class="label"><span class="label-text font-medium">CVE2fa File (.2fadata)</span></label>
            <input ref="fileInput" type="file" accept=".2fadata,.json" class="file-input file-input-bordered w-full bg-white transition-all duration-200 focus:ring-2 focus:ring-primary/20 focus:border-primary" required />
          </div>

          <div class="form-control mb-5">
            <label class="label"><span class="label-text font-medium">Master Password</span></label>
            <input v-model="masterPassword" type="password" class="input input-bordered bg-white transition-all duration-200 focus:ring-2 focus:ring-primary/20 focus:border-primary" required minlength="4" />
          </div>

          <button type="submit" class="btn btn-primary w-full transition-all duration-200 hover:shadow-md hover:shadow-primary/20 active:scale-[0.98]" :disabled="loading">
            <span v-if="loading" class="loading loading-spinner"></span>
            {{ loading ? 'Decrypting...' : 'Import CVE2fa' }}
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const fileInput = ref(null)
const masterPassword = ref('')
const loading = ref(false)
const error = ref(null)
const result = ref(null)

async function handleImport() {
  error.value = null
  result.value = null
  loading.value = true

  try {
    const file = fileInput.value?.files?.[0]
    if (!file) throw new Error('Please select a file')

    const formData = new FormData()
    formData.append('file', file)
    formData.append('masterPassword', masterPassword.value)

    const res = await fetch('/api/import', {
      method: 'POST',
      body: formData,
    })

    const data = await res.json()
    if (!res.ok) throw new Error(data.error || 'Import failed')

    result.value = `Import complete! ${data.imported} accounts added, ${data.skipped} skipped.`
    masterPassword.value = ''
    fileInput.value.value = ''
  } catch (e) {
    error.value = e.message
  } finally {
    loading.value = false
  }
}
</script>