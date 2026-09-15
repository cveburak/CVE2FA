import { ref } from 'vue'
import router from '../router'

const API = ''

export function useAuth() {
  const user = ref(null)
  const loading = ref(false)
  const error = ref(null)

  async function login(username, password) {
    loading.value = true
    error.value = null
    try {
      const res = await fetch(`${API}/api/auth/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password }),
      })
      const data = await res.json()
      if (!res.ok) throw new Error(data.error || 'Login failed')
      return data
    } catch (e) {
      error.value = e.message
      throw e
    } finally {
      loading.value = false
    }
  }

  async function logout() {
    await fetch(`${API}/api/auth/logout`, { method: 'POST' })
    router.push('/login')
  }

  async function me() {
    try {
      const res = await fetch(`${API}/api/auth/me`)
      if (res.ok) {
        user.value = await res.json()
        return user.value
      }
      return null
    } catch {
      return null
    }
  }

  async function changePassword(oldPassword, newPassword) {
    loading.value = true
    error.value = null
    try {
      const res = await fetch(`${API}/api/auth/password`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ oldPassword, newPassword }),
      })
      const data = await res.json()
      if (!res.ok) throw new Error(data.error || 'Failed to change password')
      return data
    } catch (e) {
      error.value = e.message
      throw e
    } finally {
      loading.value = false
    }
  }

  return { user, loading, error, login, logout, me, changePassword }
}
