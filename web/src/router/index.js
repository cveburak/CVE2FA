import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue'
import Dashboard from '../views/Dashboard.vue'
import Profile from '../views/Profile.vue'
import ExportView from '../views/ExportView.vue'
import ImportView from '../views/ImportView.vue'

const routes = [
  { path: '/', redirect: '/dashboard' },
  { path: '/login', name: 'Login', component: Login },
  { path: '/dashboard', name: 'Dashboard', component: Dashboard, meta: { requiresAuth: true } },
  { path: '/profile', name: 'Profile', component: Profile, meta: { requiresAuth: true } },
  { path: '/export', name: 'Export', component: ExportView, meta: { requiresAuth: true } },
  { path: '/import', name: 'Import', component: ImportView, meta: { requiresAuth: true } },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, from, next) => {
  const authenticated = hasAuth()
  if (to.meta.requiresAuth && !authenticated) {
    next('/login')
  } else if (to.name === 'Login' && authenticated) {
    next('/dashboard')
  } else {
    next()
  }
})

function hasAuth() {
  const match = document.cookie.match(new RegExp('(^| )auth=([^;]+)'))
  return match !== null
}

export default router
