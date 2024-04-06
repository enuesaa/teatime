import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import { vanillaExtractPlugin } from '@vanilla-extract/vite-plugin'
import path from 'node:path'

export default defineConfig(({ mode }) => {
  return {
    define: {
      API_BASE_URL: mode === 'production' ? '/api' : 'http://localhost:3000/api'
    },
    esbuild: {
      target: 'esnext',
      format: 'esm',
    },
    plugins: [
      react(),
      vanillaExtractPlugin(),
    ],
    resolve: {
      alias: {
        '@/': path.join(__dirname, './src/')
      }
    }
  }
})
