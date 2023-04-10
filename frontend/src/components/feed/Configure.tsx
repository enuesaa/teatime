import { PageTitle } from '@/components/common/PageTitle'
import Link from 'next/link'
import { FaPlus } from 'react-icons/fa'
import { ConfigureFeeds } from '@/components/feed/ConfigureFeeds'

export const Configure = () => {
  return (
    <>
      <PageTitle title='Feed' />
      <div>
        <Link href='/feed/configure/add'><FaPlus /></Link>
      </div>
      <ConfigureFeeds />
    </>
  )
}