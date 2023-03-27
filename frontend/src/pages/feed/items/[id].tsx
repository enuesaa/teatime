import { Header } from '@/components/common/Header'
import { useRouter } from 'next/router'
import { FeedDetail } from '@/components/feed/FeedDetail'

export default function () {
  const router = useRouter()
  const { id } = router.query
  if (typeof id !== 'string') {
    return (<></>)
  }
  
  return (
    <>
      <Header />
      <FeedDetail />
    </>
  )
}
