import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import { vanillaExtractPlugin } from '@vanilla-extract/vite-plugin'
import environmentPlugin from 'vite-plugin-environment'
import path from 'node:path'

export default defineConfig(({ mode }) => {
  const envvar = {
    API_BASE: mode === 'production' ? '/' : 'http://localhost:3000/',
  }

  return {
    plugins: [
      react(),
      vanillaExtractPlugin(),
      environmentPlugin(envvar, {
        defineOn: 'import.meta.env',
      }),
    ],
    resolve: {
      alias: {
        '@/': path.join(__dirname, './src/')
      }
    }
  }
})
