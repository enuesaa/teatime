import { css, useTheme } from '@emotion/react'
import Link from 'next/link'
import { GiCoffeePot } from 'react-icons/gi'
import { useListFeedsQuery } from '@/lib/feed'

export const ConfigureFeeds = () => {
  const theme = useTheme()
  const feeds = useListFeedsQuery({})

  const styles = {
    list: css({
      listStyleType: 'none',
      'li': {
        padding: '10px',
        border: 'solid 1px rgba(255,255,255,0.2)',
        a: {
          display: 'block',
          width: '100%',
          height: '100%',
        },
        '&:hover': {
          background: theme.color.sub,
        },
        'svg': {
          margin: '0 5px',
        }
      },
    })
  }

  return (
    <ul css={styles.list}>
      <li>
        {feeds?.items.map((v,i) => (
          <Link href={`/feed/items/${v.id}`} key={i}>
            {v?.name}
            <GiCoffeePot />
          </Link>
        ))}
      </li>
    </ul>
  )
}