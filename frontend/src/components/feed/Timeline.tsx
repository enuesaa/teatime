import { AiOutlineSwapRight } from 'react-icons/ai'
import Link from 'next/link'
import { useListItemsQuery } from '@/lib/feed'
import { ListItemsRequest } from '@/gen/v1/feed_pb'
import { useStyles } from '@/styles/use'
import { PageTitle } from '@/components/common/PageTitle'

type ItemProps = {
  title: string,
  href: string,
}
export const Item = ({ title, href }: ItemProps) => {
  const styles = useStyles(theme => ({
    li: theme().css({
      padding: '10px',
      border: 'solid 1px rgba(255,255,255,0.2)',
      a: {
        display: 'block',
        width: '100%',
        height: '100%',
      },
    })
  }))

  return (
    <li css={styles.li}>
      <Link href={href}>{title}</Link>
    </li>
  )
}

export const Timeline = () => {
  const feeditem = useListItemsQuery({ id: 'a', page: 1 } as ListItemsRequest)
  const styles = useStyles(theme => ({
    h2: theme().css({
      padding: '0 0 0 10px',
      'a': {
        margin: '10px',
        display: 'inline-block',
      },
    }),
    list: theme().css({
      listStyleType: 'none',
      padding: '0',
    }),
  }))

  return (
    <section>
      <PageTitle title='Feed'>
        <Link href='/feeds'><AiOutlineSwapRight /></Link>
      </PageTitle>
      <ul css={styles.list}>
        {
          feeditem?.items.map((v,i) => (
            <Item key={i} href={v.url} title={v.name} />
          ))
        }
      </ul>
    </section>
  )
}