import { useDeleteFeedLazy, useGetFeedQuery, useFetchLazy } from '@/lib/feed'
import { GetFeedRequest, DeleteFeedRequest, FetchRequest } from '@/gen/v1/feed_pb'
import { MouseEventHandler } from 'react'
import { useStyles } from '@/styles/use'

type Props = {
  id: string,
}
export const FeedDetail = ({ id }: Props) => {
  const data = useGetFeedQuery({ id } as GetFeedRequest)
  const { invoke: invokeDeleteFeed } = useDeleteFeedLazy()
  const { invoke: invokeFetchFeed } = useFetchLazy()

  const styles = useStyles(theme => ({
    h2: theme().css({
      padding: '0 0 0 10px',
    }),
    list: theme().css({
      listStyleType: 'none',
      padding: '0',
    }),
  }))

  const handleDeleteFeed: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    invokeDeleteFeed({ id } as DeleteFeedRequest)
  }
  const handleFetchFeed: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    invokeFetchFeed({ id } as FetchRequest)
  }

  return (
    <section>
      <h2 css={styles.h2}>
        Feed Detail: {data?.name}
      </h2>
      <button onClick={handleDeleteFeed}>delete</button>
      <button onClick={handleFetchFeed}>fetch</button>
      <ul css={styles.list}>
        {/* <TimelineItem href='https://example.com/' title='aaa' /> */}
      </ul>
    </section>
  )
}