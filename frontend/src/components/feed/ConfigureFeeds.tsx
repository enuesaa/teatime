import { useStyles } from '@/styles/use'
import Link from 'next/link'
import { GiCoffeePot } from 'react-icons/gi'
import { useListFeedsQuery } from '@/lib/feed'

export const ConfigureFeeds = () => {
  const feeds = useListFeedsQuery({})

  const styles = useStyles(theme => ({
    list: theme().css({
      listStyleType: 'none',
      'li': {
        padding: '10px',
        border: 'solid 1px rgba(255,255,255,0.2)',
        a: {
          display: 'block',
          width: '100%',
          height: '100%',
        },
        'svg': {
          margin: '0 5px',
        }
      },
    }),
  }))

  return (
    <ul css={styles.list}>
      <li>
        {feeds?.items.map((v,i) => (
          <Link href={`/feeds/${v.id}`} key={i}>
            {v?.name}
            <GiCoffeePot />
          </Link>
        ))}
      </li>
    </ul>
  )
}