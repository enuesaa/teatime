import { radixThemePreset } from 'radix-themes-tw'

/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './src/**/*.{ts,tsx}',
  ],
  theme: {
    extend: {},
  },
  plugins: [],
  presets: [radixThemePreset],
}
