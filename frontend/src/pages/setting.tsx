import { Header } from '@/components/common/Header'
import { Main } from '@/components/common/Main'
import { Dashboard } from '@/components/setting/Dashboard'

export default function Page() {
  return (
    <>
      <Header />
      <Main>
        <Dashboard />
      </Main>
    </>
  )
}
