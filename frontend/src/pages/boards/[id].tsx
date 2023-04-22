import { Header } from '@/components/common/Header'
import { Main } from '@/components/common/Main'
import { useRouter } from 'next/router'

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
        {id}
      </Main>
    </>
  )
}
