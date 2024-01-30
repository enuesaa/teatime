import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { QueryClient, QueryClientProvider } from 'react-query'
import '@radix-ui/themes/styles.css'
import { Theme } from '@radix-ui/themes'
import TopPage from './pages/index'
import SettingsPage from './pages/settings'
import '@/styles/app.css'

export const App = () => {
  const queryClient = new QueryClient()

  return (
    <QueryClientProvider client={queryClient}>
      <Theme appearance='dark' accentColor='purple'>
        <BrowserRouter>
          <Routes>
            <Route path='/' element={<TopPage />} />
            <Route path='/settings' element={<SettingsPage />} />
          </Routes>
        </BrowserRouter>
      </Theme>
    </QueryClientProvider>
  )
}
