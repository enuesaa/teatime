import { useDeleteFeedLazy, useGetFeedQuery, useFetchLazy, useRemoveAllItemsLazy } from '@/lib/feed'
import { MouseEventHandler } from 'react'
import { useStyles } from '@/styles/use'
import { PageTitle } from '@/components/common/PageTitle'

type Props = {
  id: string,
}
export const Detail = ({ id }: Props) => {
  const data = useGetFeedQuery({ id })
  const { invoke: invokeDeleteFeed } = useDeleteFeedLazy()
  const { invoke: invokeFetchFeed } = useFetchLazy()
  const { invoke: invokeRemoveAllItems } = useRemoveAllItemsLazy()

  const styles = useStyles(theme => ({
    btn: theme({ surf: 'sub', around: 'x2', decorate: 'card', hover: 'shadow' }),
  }))

  const handleDeleteFeed: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    invokeDeleteFeed({ id })
  }
  const handleFetchFeed: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    invokeFetchFeed({ id })
  }
  const handleRemoveAllItems: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    invokeRemoveAllItems({ id })
  } 

  return (
    <section>
      <PageTitle title={`Feed ${data?.name}`} />
      <button onClick={handleFetchFeed} css={styles.btn}>fetch</button>
      <button onClick={handleRemoveAllItems} css={styles.btn}>removeAllItems</button>
      <button onClick={handleDeleteFeed} css={styles.btn}>delete</button>
    </section>
  )
}
