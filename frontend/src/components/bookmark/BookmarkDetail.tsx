import { css, useTheme } from '@emotion/react'
import { useDeleteBookmarkLazy, useGetBookmarkQuery } from '@/lib/bookmark'
import { MouseEventHandler } from 'react'
import { DeleteBookmarkRequest, GetBookmarkRequest } from '@/gen/v1/bookmark_pb'

type Props = {
  id: string,
}
export const BookmarkDetail = ({ id }: Props) => {
  const theme = useTheme()

  const data = useGetBookmarkQuery({ id } as GetBookmarkRequest)
  const { invoke: invokeDeleteBookmark } = useDeleteBookmarkLazy()

  const styles = {
    h2: css(theme.heading, {
      padding: '0 0 0 10px',
    }),
  }

  const handleDeleteBookmark: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    invokeDeleteBookmark({ id } as DeleteBookmarkRequest)
  }

  return (
    <section>
      <h2 css={styles.h2}>
        Bookmark Detail: {data?.name}
      </h2>
      <button onClick={handleDeleteBookmark}>delete</button>
    </section>
  )
}