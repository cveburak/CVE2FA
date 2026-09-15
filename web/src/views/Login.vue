<template>
  <div class="min-h-[80vh] flex items-center justify-center">
    <div class="w-full max-w-md px-4 animate-fade-in">
      <div class="text-center mb-8">
        <h1 class="text-3xl font-semibold text-secondary">CVE2fa</h1>
        <p class="text-sm text-neutral mt-2">Your secure TOTP manager</p>
      </div>

      <div class="card bg-white border border-base-300 shadow-lg">
        <div class="card-body p-7">
          <h2 class="text-lg font-semibold text-secondary text-center mb-2">Sign In</h2>

          <div v-if="error" class="alert alert-error text-sm mb-4 animate-slide-down">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><line x1="15" y1="9" x2="9" y2="15"/><line x1="9" y1="9" x2="15" y2="15"/></svg>
            <span>{{ error }}</span>
          </div>

          <form @submit.prevent="handleLogin">
            <div class="form-control mb-3">
              <label class="label"><span class="label-text font-medium text-secondary">Username</span></label>
              <input v-model="username" type="text" placeholder="admin" class="input input-bordered transition-all duration-200 focus:ring-2 focus:ring-primary/20 focus:border-primary bg-white" required />
            </div>

            <div class="form-control mb-5">
              <label class="label"><span class="label-text font-medium text-secondary">Password</span></label>
              <input v-model="password" type="password" placeholder="••••••" class="input input-bordered transition-all duration-200 focus:ring-2 focus:ring-primary/20 focus:border-primary bg-white" required />
            </div>

            <button type="submit" class="btn btn-primary w-full transition-all duration-200 hover:shadow-md hover:shadow-primary/20 active:scale-[0.98]" :disabled="loading">
              <span v-if="loading" class="loading loading-spinner"></span>
              {{ loading ? 'Signing in...' : 'Sign In' }}
            </button>
          </form>
        </div>
      </div>

      <p class="text-center text-xs text-base-content/40 mt-6">End-to-end encrypted &middot; Self-hosted</p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '../composables/useAuth'

const router = useRouter()
const { login, loading } = useAuth()

const username = ref('admin')
const password = ref('')
const error = ref(null)

async function handleLogin() {
  error.value = null
  try {
    const result = await login(username.value, password.value)
    if (result.mustChangePassword) {
      router.push('/profile')
    } else {
      router.push('/dashboard')
    }
  } catch (e) {
    error.value = e.message
  }
}
</script>