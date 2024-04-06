import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { QueryClient, QueryClientProvider } from 'react-query'
import '@radix-ui/themes/styles.css'
import { Theme } from '@radix-ui/themes'
import TopPage from '@/pages/index'
import TeasPage from '@/pages/teas'
import SettingsPage from '@/pages/settings'
import '@/style.css'

export const App = () => {
  const queryClient = new QueryClient()

  return (
    <QueryClientProvider client={queryClient}>
      <Theme appearance='dark' accentColor='purple'>
        <BrowserRouter>
          <Routes>
            <Route path='/' element={<TopPage />} />
            <Route path='/settings' element={<SettingsPage />} />
            <Route path='/teapods/:teapod/teas/:teabox' element={<TeasPage />} />
          </Routes>
        </BrowserRouter>
      </Theme>
    </QueryClientProvider>
  )
}
