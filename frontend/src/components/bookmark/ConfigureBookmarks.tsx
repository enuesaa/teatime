import { useStyles } from '@/styles/use'
import { useListBookmarksQuery } from '@/lib/bookmark'
import Link from 'next/link'

export const ConfigureBookmarks = () => {
  const bookmarks = useListBookmarksQuery({})
  const styles = useStyles(theme => ({
    item: theme({}).css({
      padding: '10px',
      margin: '10px',
      borderRadius: '10px',
      border: 'solid 1px rgba(255,255,255,0.2)',
    }),
  }))

  return (
    <>
      {bookmarks?.items.map((v,i) => (
        <Link key={i} href={`/bookmark/items/${v.id}`} css={styles.item}>
          {v?.name}
        </Link>
      ))}
    </>
  )
}