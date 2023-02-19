import type { AppProps } from 'next/app'
import { Global, ThemeProvider } from '@emotion/react'
import { globalStyle } from '@/styles/global'
import { baseTheme } from '@/styles/theme'
import { QueryClient, QueryClientProvider } from 'react-query'

export default function App({ Component, pageProps }: AppProps) {
  const queryClient = new QueryClient()

  return (
    <>
      <Global styles={globalStyle} />
      <ThemeProvider theme={baseTheme}>
        <QueryClientProvider client={queryClient}>
          <Component {...pageProps} />
        </QueryClientProvider>
      </ThemeProvider>
    </>
  )
}
