import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import path from 'path'

// https://vite.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      '@': path.resolve(import.meta.dirname, './src'),
    },
  },
  plugins: [react()],
  server: {
    proxy: {
      '/api': 'http://localhost:8080',
      '/webhook': 'http://localhost:8080',
      '/tg': 'http://localhost:8080',
    },
  },
})
