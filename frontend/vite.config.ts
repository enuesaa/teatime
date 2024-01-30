import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import { vanillaExtractPlugin } from '@vanilla-extract/vite-plugin'
import path from 'node:path'

export default defineConfig({
  root: './',
  envDir: '../',
  esbuild: {
    target: 'esnext',
    format: 'esm',
  },
  build: {
    outDir: './dist',
    rollupOptions: {
      output: {
        entryFileNames: '[name].js',
        assetFileNames: 'assets/[name][extname]',
      },
		},
  },
  plugins: [
    react(),
    vanillaExtractPlugin(),
  ],
  resolve: {
    alias: {
      '@/': path.join(__dirname, './src/')
    }
  },
})