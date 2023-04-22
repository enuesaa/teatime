import { PageTitle } from '@/components/common/PageTitle'
import Link from 'next/link'
import { FaPlus } from 'react-icons/fa'
import { useStyles } from '@/styles/use'
import { useListBookmarksQuery } from '@/lib/bookmark'

const List = () => {
  const bookmarks = useListBookmarksQuery({})
  const styles = useStyles(theme => ({
    item: theme().css({
      padding: '10px',
      margin: '10px',
      borderRadius: '10px',
      border: 'solid 1px rgba(255,255,255,0.2)',
    }),
  }))

  return (
    <>
      {bookmarks?.items.map((v,i) => (
        <Link key={i} href={`/bookmarks/${v.id}`} css={styles.item}>
          {v?.name}
        </Link>
      ))}
    </>
  )
}

export const Configure = () => {
  return (
    <>
      <PageTitle title='Feed'>
        <Link href='/bookmarks?create'><FaPlus /></Link>
      </PageTitle>
      <List />
    </>
  )
}