import { TimelineItem } from '@/components/feed/TimelineItem'
import { AiOutlineSwapRight } from 'react-icons/ai'
import Link from 'next/link'
import { useListItemsQuery } from '@/lib/feed'
import { ListItemsRequest } from '@/gen/v1/feed_pb'
import { useStyles } from '@/styles/use'

export const Timeline = () => {
  const feeditem = useListItemsQuery({ id: 'a', page: 1 } as ListItemsRequest)
  const styles = useStyles(theme => ({
    h2: theme({}).css({
      padding: '0 0 0 10px',
      'a': {
        margin: '10px',
        display: 'inline-block',
        // color: theme.color.main,
        // fontSize: theme.fontSize.large,
      },
    }),
    list: theme({}).css({
      listStyleType: 'none',
      padding: '0',
    }),
  }))

  return (
    <section>
      <h2 css={styles.h2}>
        Feed
        <Link href='/feed/configure'><AiOutlineSwapRight /></Link>
      </h2>
      <ul css={styles.list}>
        {
          feeditem?.items.map((v,i) => (
            <TimelineItem key={i} href={v.url} title={v.name} />
          ))
        }
      </ul>
    </section>
  )
}