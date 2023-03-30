import { Header } from '@/components/common/Header'
import { Configure } from '@/components/board/Cofigure'
import { Main } from '@/components/common/Main'

export default function Page() {
  return (
    <>
      <Header />
      <Main>
        <Configure />        
      </Main>
    </>
  )
}
