import { ref, onUnmounted } from 'vue'

const API = ''

export function useTotp() {
  const accounts = ref([])
  const codes = ref({})
  const loading = ref(false)
  const error = ref(null)
  let interval = null

  async function fetchAccounts() {
    loading.value = true
    try {
      const res = await fetch(`${API}/api/accounts`)
      if (!res.ok) throw new Error('Failed to fetch accounts')
      accounts.value = await res.json()
      accounts.value.forEach(a => fetchCode(a.id))
    } catch (e) {
      error.value = e.message
    } finally {
      loading.value = false
    }
  }

  async function fetchCode(accountId) {
    try {
      const res = await fetch(`${API}/api/accounts/${accountId}/code`)
      if (!res.ok) throw new Error('Failed to fetch code')
      codes.value[accountId] = await res.json()
    } catch (e) {
      console.error('fetch code error:', e)
    }
  }

  function startAutoRefresh() {
    stopAutoRefresh()
    interval = setInterval(async () => {
      for (const account of accounts.value) {
        await fetchCode(account.id)
      }
    }, 1000)
  }

  function stopAutoRefresh() {
    if (interval) {
      clearInterval(interval)
      interval = null
    }
  }

  async function addAccount(issuer, accountName, secretBase32) {
    const res = await fetch(`${API}/api/accounts`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ issuer, accountName, secretBase32 }),
    })
    const data = await res.json()
    if (!res.ok) throw new Error(data.error || 'Failed to add account')
    await fetchAccounts()
    return data
  }

  async function deleteAccount(id) {
    const res = await fetch(`${API}/api/accounts/${id}`, { method: 'DELETE' })
    if (!res.ok) throw new Error('Failed to delete account')
    accounts.value = accounts.value.filter(a => a.id !== id)
    delete codes.value[id]
  }

  function getProgress(accountId) {
    const c = codes.value[accountId]
    if (!c) return 0
    return (c.remaining / c.period) * 100
  }

  function getCode(accountId) {
    const c = codes.value[accountId]
    return c ? c.code : '------'
  }

  function getRemaining(accountId) {
    const c = codes.value[accountId]
    return c ? c.remaining : 0
  }

  onUnmounted(() => stopAutoRefresh())

  return {
    accounts, codes, loading, error,
    fetchAccounts, addAccount, deleteAccount,
    startAutoRefresh, stopAutoRefresh,
    getProgress, getCode, getRemaining,
  }
}
