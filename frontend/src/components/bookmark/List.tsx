import { PageTitle } from '@/components/common/PageTitle'
import Link from 'next/link'
import { FaPlus } from 'react-icons/fa'
import { Configure } from '@/components/bookmark/Configure'

export const List = () => {
  return (
    <>
      <PageTitle title='Bookmark' />
      <div>
        <Link href='/bookmarks?create'><FaPlus /></Link>
      </div>
      <Configure />
    </>
  )
}