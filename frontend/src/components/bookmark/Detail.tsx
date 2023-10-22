import { useStyles } from '@/styles/use'
import { useDeleteBookmarkLazy, useGetBookmarkQuery } from '@/lib/bookmark'
import { MouseEventHandler } from 'react'

type Props = {
  id: string,
}
export const Detail = ({ id }: Props) => {
  const data = useGetBookmarkQuery({ id })
  const { invoke: invokeDeleteBookmark } = useDeleteBookmarkLazy()

  const styles = useStyles(theme => ({
    h2: theme().css({
      padding: '0 0 0 10px',
    }),
  }))

  const handleDeleteBookmark: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    invokeDeleteBookmark({ id })
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