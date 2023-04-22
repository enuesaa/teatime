import { Header } from '@/components/common/Header'
import { Main } from '@/components/common/Main'
import { List } from '@/components/board/List'
import { Add } from '@/components/board/Add'
import { useRouter } from 'next/router'

export default function Page() {
  const router = useRouter()
  const { create } = router.query

  return (
    <>
      <Header />
      <Main>
        {create === undefined ? <List /> : <Add />}
      </Main>
    </>
  )
}
