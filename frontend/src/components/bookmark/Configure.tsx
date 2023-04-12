import { PageTitle } from '@/components/common/PageTitle'
import Link from 'next/link'
import { FaPlus } from 'react-icons/fa'
import { ConfigureBookmarks } from '@/components/bookmark/ConfigureBookmarks'

export const Configure = () => {
  return (
    <>
      <PageTitle title='Bookmark' />
      <div>
        <Link href='/bookmark/configure/add'><FaPlus /></Link>
      </div>
      <ConfigureBookmarks />
    </>
  )
}