import { css, useTheme } from '@emotion/react'
import { TopDashboardItem } from '@/components/board/TopDashboardItem'
import { AiOutlineSwapRight } from 'react-icons/ai'
import Link from 'next/link'

export const TopDashboard = () => {
  const theme = useTheme()

  const styles = {
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
