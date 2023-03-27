import { Header } from '@/components/common/Header'
import { useRouter } from 'next/router'

export default function () {
  const router = useRouter()
  const { id } = router.query
  console.log(router)
  if (typeof id !== 'string') {
    return (<></>)
  }
  console.log(id)
  
  return (
    <>
      <Header />
      {id}
    </>
  )
}
