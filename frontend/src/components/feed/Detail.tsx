import { useDeleteFeedLazy, useGetFeedQuery, useFetchLazy } from '@/lib/feed'
import { GetFeedRequest, DeleteFeedRequest, FetchRequest } from '@/gen/v1/feed_pb'
import { MouseEventHandler } from 'react'
import { useStyles } from '@/styles/use'
import { PageTitle } from '@/components/common/PageTitle'

type Props = {
  id: string,
}
export const Detail = ({ id }: Props) => {
  const data = useGetFeedQuery({ id } as GetFeedRequest)
  const { invoke: invokeDeleteFeed } = useDeleteFeedLazy()
  const { invoke: invokeFetchFeed } = useFetchLazy()

  const styles = useStyles(theme => ({
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
      <PageTitle title={`Feed Detail ${data?.name}`} />
      <button onClick={handleDeleteFeed}>delete</button>
      <button onClick={handleFetchFeed}>fetch</button>
      <ul css={styles.list}>
        {/* <TimelineItem href='https://example.com/' title='aaa' /> */}
      </ul>
    </section>
  )
}