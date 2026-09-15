<template>
  <div class="navbar bg-white/95 backdrop-blur-sm border-b border-base-300 sticky top-0 z-50 px-4">
    <div class="flex-1">
      <router-link to="/dashboard" class="font-semibold text-secondary text-lg transition-colors duration-200 hover:text-primary">
        CVE2fa
      </router-link>
    </div>
    <div class="flex-none gap-1">
      <router-link to="/dashboard" custom v-slot="{ isActive, navigate, href }">
        <a :href="href" @click="navigate" :class="[
          'btn btn-sm gap-2 rounded-full px-4 transition-all duration-200',
          isActive
            ? 'bg-primary/10 text-primary border-transparent hover:bg-primary/15'
            : 'btn-ghost text-neutral hover:bg-base-200 hover:text-secondary'
        ]">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><polyline points="9 22 9 12 15 12 15 22"/></svg>
          Dashboard
        </a>
      </router-link>
      <router-link to="/export" custom v-slot="{ isActive, navigate, href }">
        <a :href="href" @click="navigate" :class="[
          'btn btn-sm gap-2 rounded-full px-4 transition-all duration-200',
          isActive
            ? 'bg-primary/10 text-primary border-transparent hover:bg-primary/15'
            : 'btn-ghost text-neutral hover:bg-base-200 hover:text-secondary'
        ]">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/><polyline points="7 10 12 15 17 10"/><line x1="12" y1="15" x2="12" y2="3"/></svg>
          Export
        </a>
      </router-link>
      <router-link to="/import" custom v-slot="{ isActive, navigate, href }">
        <a :href="href" @click="navigate" :class="[
          'btn btn-sm gap-2 rounded-full px-4 transition-all duration-200',
          isActive
            ? 'bg-primary/10 text-primary border-transparent hover:bg-primary/15'
            : 'btn-ghost text-neutral hover:bg-base-200 hover:text-secondary'
        ]">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/><polyline points="17 8 12 3 7 8"/><line x1="12" y1="3" x2="12" y2="15"/></svg>
          Import
        </a>
      </router-link>
      <router-link to="/profile" custom v-slot="{ isActive, navigate, href }">
        <a :href="href" @click="navigate" :class="[
          'btn btn-sm gap-2 rounded-full px-4 transition-all duration-200',
          isActive
            ? 'bg-primary/10 text-primary border-transparent hover:bg-primary/15'
            : 'btn-ghost text-neutral hover:bg-base-200 hover:text-secondary'
        ]">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/></svg>
          Profile
        </a>
      </router-link>
      <div class="w-px h-6 bg-base-300 my-auto mx-1"></div>
      <button @click="handleLogout" class="btn btn-ghost btn-sm gap-2 px-4 rounded-full text-error transition-all duration-200 hover:bg-error/10 group">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 transition-transform duration-200 group-hover:rotate-90" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/><polyline points="16 17 21 12 16 7"/><line x1="21" y1="12" x2="9" y2="12"/></svg>
        Logout
      </button>
    </div>
  </div>
</template>

<script setup>
import { useAuth } from '../composables/useAuth'

const { logout } = useAuth()

async function handleLogout() {
  await logout()
}
</script>