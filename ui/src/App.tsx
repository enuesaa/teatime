import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { QueryClient, QueryClientProvider } from 'react-query'
import { Theme } from '@radix-ui/themes'
import TopPage from '@/pages/index'
import TeasPage from '@/pages/teas'
import SettingsPage from '@/pages/settings'
import { Layout } from './components/common/Layout'
import '@radix-ui/themes/styles.css'
import '@/style.css'

export const App = () => {
  const queryClient = new QueryClient()

  return (
    <QueryClientProvider client={queryClient}>
      <Theme appearance='dark' accentColor='purple'>
        <Layout>
          <BrowserRouter>
            <Routes>
              <Route path='/' element={<TopPage />} />
              <Route path='/settings' element={<SettingsPage />} />
              <Route path='/teapods/:teapod/teas' element={<TeasPage />} />
            </Routes>
          </BrowserRouter>
        </Layout>
      </Theme>
    </QueryClientProvider>
  )
}
