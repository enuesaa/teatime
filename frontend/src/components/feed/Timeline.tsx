import { css, useTheme } from '@emotion/react'
import { TimelineItem } from '@/components/feed/TimelineItem'
import { AiOutlineSwapRight } from 'react-icons/ai'
import Link from 'next/link'
import { useListItemsQuery } from '@/lib/feed'
import { ListItemsRequest } from '@/gen/v1/feed_pb'

export const Timeline = () => {
  const theme = useTheme()

  const feeditem = useListItemsQuery({ id: 'a', page: 1 } as ListItemsRequest)
  const styles = {
    main: css({
      margin: '20px',
      padding: '0 10px 10px 10px',
      color: theme.color.main,
    }),
    h2: css(theme.heading, {
      padding: '0 0 0 10px',
      'a': {
        margin: '10px',
        display: 'inline-block',
        color: theme.color.main,
        fontSize: theme.fontSize.large,
      },
    }),
    list: css({
      listStyleType: 'none',
      padding: '0',
    }),
  }

  return (
    <section css={styles.main}>
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