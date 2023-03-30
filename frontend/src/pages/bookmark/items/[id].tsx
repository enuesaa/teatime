import { Header } from '@/components/common/Header'
import { useRouter } from 'next/router'
import { BookmarkDetail } from '@/components/bookmark/BookmarkDetail'

export default function Page() {
  const router = useRouter()
  const { id } = router.query
  if (typeof id !== 'string') {
    return (<></>)
  }
  
  return (
    <>
      <Header />
      <BookmarkDetail id={id} />
    </>
  )
}
