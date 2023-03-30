import { css, useTheme } from '@emotion/react'
import { TimelineItem } from '@/components/feed/TimelineItem'
import { useDeleteFeedLazy, useGetFeedQuery, useFetchLazy } from '@/lib/feed'
import { GetFeedRequest, DeleteFeedRequest, FetchRequest } from '@/gen/v1/feed_pb'
import { MouseEventHandler } from 'react'

type Props = {
  id: string,
}
export const FeedDetail = ({ id }: Props) => {
  const theme = useTheme()

  const data = useGetFeedQuery({ id } as GetFeedRequest)
  const { invoke: invokeDeleteFeed } = useDeleteFeedLazy()
  const { invoke: invokeFetchFeed } = useFetchLazy()

  const styles = {
    main: css({
      margin: '20px',
      padding: '0 10px 10px 10px',
      color: theme.color.main,
    }),
    h2: css(theme.heading, {
      padding: '0 0 0 10px',
    }),
    list: css({
      listStyleType: 'none',
      padding: '0',
    }),
  }

  const handleDeleteFeed: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    invokeDeleteFeed({ id } as DeleteFeedRequest)
  }
  const handleFetchFeed: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    invokeFetchFeed({ id } as FetchRequest)
  }

  return (
    <section css={styles.main}>
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