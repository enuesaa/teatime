import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { QueryClient, QueryClientProvider } from 'react-query'
import '@radix-ui/themes/styles.css'
import { Theme } from '@radix-ui/themes'
import TopPage from './pages/index'
import SettingPage from './pages/setting'

export const App = () => {
  const queryClient = new QueryClient()

  return (
    <QueryClientProvider client={queryClient}>
      <Theme appearance='dark'>
        <BrowserRouter>
          <Routes>
          <Route path='/' element={<TopPage />} />
          <Route path='/setting' element={<SettingPage />} />
          </Routes>
        </BrowserRouter>
      </Theme>
    </QueryClientProvider>
  )
}
