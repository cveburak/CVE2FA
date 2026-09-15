import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  server: {
    port: 2026,
    proxy: {
      '/api': 'http://localhost:2026'
    }
  },
  build: {
    outDir: 'dist',
    assetsDir: 'assets'
  }
})
