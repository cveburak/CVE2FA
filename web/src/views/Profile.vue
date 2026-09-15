<template>
  <div class="max-w-md mx-auto mt-8">
    <div class="card bg-white border border-base-300 shadow-sm">
      <div class="card-body">
        <h2 class="card-title text-lg text-secondary mb-5">
          <div class="flex items-center justify-center h-9 w-9 rounded-lg bg-primary/10 text-primary">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/></svg>
          </div>
          Profile &amp; Security
        </h2>

        <div v-if="error" class="alert alert-error text-sm mb-4 animate-slide-down">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><line x1="15" y1="9" x2="9" y2="15"/><line x1="9" y1="9" x2="15" y2="15"/></svg>
          <span>{{ error }}</span>
        </div>

        <div v-if="success" class="alert alert-success text-sm mb-4 animate-slide-down">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polyline points="20 6 9 17 4 12"/></svg>
          <span>{{ success }}</span>
        </div>

        <form @submit.prevent="handleChangePassword">
          <div class="form-control mb-3">
            <label class="label"><span class="label-text font-medium">Current Password</span></label>
            <input v-model="oldPassword" type="password" class="input input-bordered bg-white transition-all duration-200 focus:ring-2 focus:ring-primary/20 focus:border-primary" required />
          </div>

          <div class="form-control mb-3">
            <label class="label"><span class="label-text font-medium">New Password</span></label>
            <input v-model="newPassword" type="password" class="input input-bordered bg-white transition-all duration-200 focus:ring-2 focus:ring-primary/20 focus:border-primary" required />
          </div>

          <div class="form-control mb-5">
            <label class="label"><span class="label-text font-medium">Confirm New Password</span></label>
            <input v-model="confirmPassword" type="password" class="input input-bordered bg-white transition-all duration-200 focus:ring-2 focus:ring-primary/20 focus:border-primary" required />
          </div>

          <button type="submit" class="btn btn-primary w-full transition-all duration-200 hover:shadow-md hover:shadow-primary/20 active:scale-[0.98]" :disabled="loading">
            <span v-if="loading" class="loading loading-spinner"></span>
            Change Password
          </button>
        </form>
      </div>
    </div>

    <div class="card bg-white border border-base-300 shadow-sm mt-5">
      <div class="card-body">
        <h3 class="card-title text-base text-secondary">
          <div class="flex items-center justify-center h-9 w-9 rounded-lg bg-base-200 text-secondary">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="3"/><path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"/></svg>
          </div>
          Preferences
        </h3>

        <div class="mt-2 divide-y divide-base-200">
          <div class="flex items-center justify-between py-3">
            <div>
              <p class="text-sm font-medium text-secondary">Theme</p>
              <p class="text-xs text-base-content/50">Light interface</p>
            </div>
            <span class="badge badge-outline badge-sm bg-primary/10 border-primary/20 text-primary gap-1">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="5"/><path d="M12 1v2M12 21v2M4.22 4.22l1.42 1.42M18.36 18.36l1.42 1.42M1 12h2M21 12h2M4.22 19.78l1.42-1.42M18.36 5.64l1.42-1.42"/></svg>
              Light
            </span>
          </div>
          <div class="flex items-center justify-between py-3">
            <div>
              <p class="text-sm font-medium text-secondary">Session Timeout</p>
              <p class="text-xs text-base-content/50">Auto sign-out after inactivity</p>
            </div>
            <span class="badge badge-ghost badge-sm">30 min</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useAuth } from '../composables/useAuth'

const { changePassword, loading } = useAuth()

const oldPassword = ref('')
const newPassword = ref('')
const confirmPassword = ref('')
const error = ref(null)
const success = ref(null)

async function handleChangePassword() {
  error.value = null
  success.value = null

  if (newPassword.value !== confirmPassword.value) {
    error.value = 'Passwords do not match'
    return
  }

  if (newPassword.value.length < 4) {
    error.value = 'Password must be at least 4 characters'
    return
  }

  try {
    await changePassword(oldPassword.value, newPassword.value)
    success.value = 'Password changed successfully'
    oldPassword.value = ''
    newPassword.value = ''
    confirmPassword.value = ''
  } catch (e) {
    error.value = e.message
  }
}
</script>