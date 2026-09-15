<template>
  <div>
    <div class="flex justify-between items-center gap-4 mb-6">
      <h1 class="text-2xl font-semibold text-secondary">Your Accounts</h1>
      <button class="btn btn-primary transition-all duration-200 hover:shadow-md hover:shadow-primary/20 hover:scale-[1.02] active:scale-[0.98]" onclick="add_modal.showModal()">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="12" y1="5" x2="12" y2="19"/><line x1="5" y1="12" x2="19" y2="12"/></svg>
        Add Account
      </button>
    </div>

    <div v-if="loading" class="flex justify-center py-16">
      <span class="loading loading-spinner loading-lg text-primary"></span>
    </div>

    <div v-else-if="accounts.length === 0" class="text-center py-16 animate-fade-in">
      <div class="inline-flex items-center justify-center h-24 w-24 rounded-full bg-base-200 mb-5">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-11 w-11 text-base-content/40" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><rect x="3" y="11" width="18" height="11" rx="2" ry="2"/><path d="M7 11V7a5 5 0 0 1 10 0v4"/></svg>
      </div>
      <p class="text-lg font-medium text-secondary">No accounts yet</p>
      <p class="text-sm text-base-content/50 mt-1">Click "Add Account" to get started</p>
      <button class="btn btn-primary mt-6 transition-all duration-200 hover:scale-[1.02] active:scale-[0.98]" onclick="add_modal.showModal()">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="12" y1="5" x2="12" y2="19"/><line x1="5" y1="12" x2="19" y2="12"/></svg>
        Add Account
      </button>
    </div>

    <div v-else>
      <div class="relative mb-5 max-w-md">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 absolute left-3.5 top-1/2 -translate-y-1/2 text-base-content/40 pointer-events-none" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/></svg>
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Search accounts…"
          class="input input-bordered w-full pl-10 transition-all duration-200 focus:ring-2 focus:ring-primary/20 focus:border-primary bg-white"
        />
      </div>

      <TransitionGroup
        v-if="filteredAccounts.length > 0"
        tag="div"
        name="card-list"
        class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4"
      >
        <TotpCard
          v-for="account in filteredAccounts"
          :key="account.id"
          :account="account"
          :code="getCode(account.id)"
          :progress="getProgress(account.id)"
          :remaining="getRemaining(account.id)"
          @delete="handleDelete"
        />
      </TransitionGroup>

      <div v-else class="text-center py-16 animate-fade-in">
        <div class="inline-flex items-center justify-center h-16 w-16 rounded-full bg-base-200 mb-4">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-7 w-7 text-base-content/40" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/></svg>
        </div>
        <p class="text-base font-medium text-secondary">No matching accounts</p>
        <p class="text-sm text-base-content/50 mt-1">Try a different search term</p>
      </div>
    </div>

    <dialog id="add_modal" class="modal">
      <div class="modal-box rounded-2xl">
        <h3 class="font-semibold text-lg mb-4 text-secondary">Add New Account</h3>

        <div v-if="addError" class="alert alert-error text-sm mb-3 animate-slide-down">
          <span>{{ addError }}</span>
        </div>

        <form @submit.prevent="handleAdd">
          <div class="form-control mb-3">
            <label class="label"><span class="label-text font-medium">Service Provider</span></label>
            <input v-model="newIssuer" type="text" placeholder="Google, GitHub, etc." class="input input-bordered bg-white transition-all duration-200 focus:ring-2 focus:ring-primary/20 focus:border-primary" required />
          </div>

          <div class="form-control mb-3">
            <label class="label"><span class="label-text font-medium">Account Identifier</span></label>
            <input v-model="newAccountName" type="text" placeholder="user@email.com" class="input input-bordered bg-white transition-all duration-200 focus:ring-2 focus:ring-primary/20 focus:border-primary" required />
          </div>

          <div class="form-control mb-4">
            <label class="label"><span class="label-text font-medium">Secret Key (Base32)</span></label>
            <input v-model="newSecret" type="text" placeholder="JBSWY3DPEHPK3PXP" class="input input-bordered font-mono bg-white transition-all duration-200 focus:ring-2 focus:ring-primary/20 focus:border-primary" required />
          </div>

          <div class="modal-action">
            <button type="button" class="btn transition-all duration-200 hover:bg-base-200" onclick="add_modal.close()">Cancel</button>
            <button type="submit" class="btn btn-primary transition-all duration-200 hover:shadow-md active:scale-[0.98]" :disabled="adding">
              <span v-if="adding" class="loading loading-spinner"></span>
              Add Account
            </button>
          </div>
        </form>
      </div>
      <form method="dialog" class="modal-backdrop"><button>close</button></form>
    </dialog>

    <dialog id="delete_modal" class="modal">
      <div class="modal-box rounded-2xl">
        <div class="flex items-start gap-3 mb-3">
          <div class="flex items-center justify-center h-10 w-10 rounded-full bg-error/10 shrink-0">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-error" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polyline points="3 6 5 6 21 6"/><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/></svg>
          </div>
          <div>
            <h3 class="font-semibold text-lg text-secondary">Delete account?</h3>
            <p class="text-sm text-base-content/50 mt-1">This action cannot be undone.</p>
          </div>
        </div>
        <div class="modal-action">
          <button type="button" class="btn transition-all duration-200 hover:bg-base-200 active:scale-[0.98]" onclick="delete_modal.close()">Cancel</button>
          <button type="button" class="btn btn-error transition-all duration-200 hover:shadow-md active:scale-[0.98]" @click="confirmDelete">Delete</button>
        </div>
      </div>
      <form method="dialog" class="modal-backdrop"><button>close</button></form>
    </dialog>
  </div>
</template>

<style scoped>
.card-list-enter-active {
  animation: card-enter 300ms ease-out;
}

.card-list-leave-active {
  animation: card-leave 200ms ease-in;
}

.card-list-move {
  transition: transform 300ms ease-out;
}

@keyframes card-enter {
  from {
    opacity: 0;
    transform: translateY(16px) scale(0.97);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

@keyframes card-leave {
  from {
    opacity: 1;
    transform: scale(1);
  }
  to {
    opacity: 0;
    transform: scale(0.95);
  }
}

@media (prefers-reduced-motion: reduce) {
  .card-list-enter-active,
  .card-list-leave-active,
  .card-list-move {
    animation: none !important;
    transition: none !important;
  }
}
</style>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import TotpCard from '../components/TotpCard.vue'
import { useTotp } from '../composables/useTotp'

const {
  accounts, loading, error,
  fetchAccounts, addAccount, deleteAccount,
  startAutoRefresh, stopAutoRefresh,
  getProgress, getCode, getRemaining,
} = useTotp()

const newIssuer = ref('')
const newAccountName = ref('')
const newSecret = ref('')
const adding = ref(false)
const addError = ref(null)

const searchQuery = ref('')
const deleteTarget = ref(null)

const filteredAccounts = computed(() => {
  const query = searchQuery.value.trim().toLowerCase()
  if (!query) return accounts.value
  return accounts.value.filter(a =>
    (a.issuer || '').toLowerCase().includes(query) ||
    (a.accountName || '').toLowerCase().includes(query)
  )
})

onMounted(async () => {
  await fetchAccounts()
  startAutoRefresh()
})

onUnmounted(() => stopAutoRefresh())

async function handleAdd() {
  adding.value = true
  addError.value = null
  try {
    await addAccount(newIssuer.value, newAccountName.value, newSecret.value)
    newIssuer.value = ''
    newAccountName.value = ''
    newSecret.value = ''
    document.getElementById('add_modal').close()
  } catch (e) {
    addError.value = e.message
  } finally {
    adding.value = false
  }
}

function handleDelete(id) {
  deleteTarget.value = id
  document.getElementById('delete_modal').showModal()
}

async function confirmDelete() {
  if (deleteTarget.value === null) return
  const id = deleteTarget.value
  deleteTarget.value = null
  document.getElementById('delete_modal').close()
  try {
    await deleteAccount(id)
  } catch (e) {
    alert(e.message)
  }
}
</script>