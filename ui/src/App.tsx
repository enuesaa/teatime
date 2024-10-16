import { Layout } from './components/common/Layout'
import TopPage from '@/pages/index'
import SettingsPage from '@/pages/settings'
import TeasPage from '@/pages/teas'
import LogsPage from '@/pages/logs'
import '@/style.css'
import { Theme } from '@radix-ui/themes'
import '@radix-ui/themes/styles.css'
import { QueryClient, QueryClientProvider } from 'react-query'
import { BrowserRouter, Routes, Route } from 'react-router-dom'

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
          </Layout>
        </Theme>
      </BrowserRouter>
    </QueryClientProvider>
  )
}
