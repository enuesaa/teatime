import { TopDashboardItem } from '@/components/board/TopDashboardItem'
import { AiOutlineSwapRight } from 'react-icons/ai'
import Link from 'next/link'
import { useStyles } from '@/styles/use'

export const TopDashboard = () => {
  const styles = useStyles(theme => ({
    h2: theme({size: 'x3'}).css({
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
      <h2 css={styles.h2}>
        Board
        <Link href='/board/configure'><AiOutlineSwapRight /></Link>
      </h2>
      <ul css={styles.list}>
        <TopDashboardItem title='aa' id='aa' />
      </ul>
    </section>
  )
} 
