<template>
  <div class="min-h-screen bg-base-100">
    <Navbar v-if="isAuthenticated" />
    <main class="container mx-auto px-4 py-8">
      <router-view v-slot="{ Component }">
        <transition name="page" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </main>
  </div>
</template>

<style>
.page-enter-active,
.page-leave-active {
  transition: opacity 200ms ease-out, transform 200ms ease-out;
}

.page-enter-from {
  opacity: 0;
  transform: translateY(8px);
}

.page-leave-to {
  opacity: 0;
  transform: translateY(-4px);
}

@media (prefers-reduced-motion: reduce) {
  .page-enter-active,
  .page-leave-active {
    transition: none;
  }
}
</style>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import Navbar from './components/Navbar.vue'

const router = useRouter()
const isAuthenticated = ref(false)

function hasAuth() {
  return document.cookie.includes('auth=1')
}

onMounted(() => {
  isAuthenticated.value = hasAuth()
})

router.afterEach(() => {
  isAuthenticated.value = hasAuth()
})
</script>