import { Layout } from './components/common/Layout'
import TopPage from '@/pages/top'
import LogsPage from '@/pages/logs'
import SettingsPage from '@/pages/settings'
import TeasPage from '@/pages/teas'
import '@/style.css'
import { Theme } from '@radix-ui/themes'
import '@radix-ui/themes/styles.css'
import { QueryClient, QueryClientProvider } from 'react-query'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { ToastProvider } from '@/components/common/ToastProvider'

export const App = () => {
  const queryClient = new QueryClient()

  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <Theme appearance='dark' accentColor='purple'>
          <Layout>
            <Routes>
              <Route path='/' element={<TopPage />} />
              <Route path='/logs' element={<LogsPage />} />
              <Route path='/settings' element={<SettingsPage />} />
              <Route path='/teapods/:teapod/teas' element={<TeasPage />} />
            </Routes>
            <ToastProvider />
          </Layout>
        </Theme>
      </BrowserRouter>
    </QueryClientProvider>
  )
}
