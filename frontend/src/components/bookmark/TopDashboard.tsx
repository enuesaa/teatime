import { TopDashboardItem } from '@/components/bookmark/TopDashboardItem'
import { AiOutlineSwapRight } from 'react-icons/ai'
import Link from 'next/link'
import { useListBookmarksQuery } from '@/lib/bookmark'
import { useStyles } from '@/styles/use'

export const TopDashboard = () => {
  const list = useListBookmarksQuery({})

  const styles = useStyles(theme => ({
    h2: theme().css({
      padding: '0 0 0 10px',
      'a': {
        margin: '10px',
        display: 'inline-block',
        // color: theme.color.main,
        // fontSize: theme.fontSize.large,
      },
    }),
    list: theme().css({
      listStyleType: 'none',
      padding: '0',
    }),
  }))

  return (
    <section>
      <h2 css={styles.h2}>
        Bookmark
        <Link href='/bookmark/configure'><AiOutlineSwapRight /></Link>
      </h2>
      <ul css={styles.list}>
        {
          list?.items.map((v,i) => {
            return (<TopDashboardItem title={v?.name} href={v?.url ?? 'https://example.com/'} key={i} />)
          })
        }
      </ul>
    </section>
  )
} 
