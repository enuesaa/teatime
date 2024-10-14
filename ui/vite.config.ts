import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import { vanillaExtractPlugin } from '@vanilla-extract/vite-plugin'
import path from 'node:path'

export default defineConfig(() => {
  return {
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
