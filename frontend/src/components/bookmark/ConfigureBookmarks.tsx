import { css, useTheme } from '@emotion/react'
import { useListBookmarksQuery } from '@/lib/bookmark'
import Link from 'next/link'

export const ConfigureBookmarks = () => {
  const theme = useTheme()

  const bookmarks = useListBookmarksQuery({})
  const styles = {
    item: css({
      padding: '10px',
      margin: '10px',
      borderRadius: '10px',
      border: 'solid 1px rgba(255,255,255,0.2)',
    }),
  }

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