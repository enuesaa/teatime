import { PageTitle } from '@/components/common/PageTitle'
import Link from 'next/link'
import { FaPlus } from 'react-icons/fa'
import { Configure } from '@/components/feed/Configure'

export const List = () => {
  return (
    <>
      <PageTitle title='Feed' />
      <div>
        <Link href='/feeds?create'><FaPlus /></Link>
      </div>
      <Configure />
    </>
  )
}