import { Header } from '@/components/common/Header'
import { useRouter } from 'next/router'
import { Detail } from '@/components/feed/Detail'
import { Main } from '@/components/common/Main'

export default function Page() {
  const router = useRouter()
  const { id } = router.query
  if (typeof id !== 'string') {
    return (<></>)
  }
  
  return (
    <>
      <Header />
      <Main>
        <Detail id={id} />
      </Main>
    </>
  )
}
