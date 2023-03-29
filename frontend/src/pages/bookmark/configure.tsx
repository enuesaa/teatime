import { Header } from '@/components/common/Header'
import { Configure } from '@/components/bookmark/Cofigure'
import { Main } from '@/components/common/Main'

export default function () {
  return (
    <>
      <Header />
      <Main>
        <Configure />
      </Main>
    </>
  )
}
