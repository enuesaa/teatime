import { css, useTheme } from '@emotion/react'
import { TopDashboardItem } from '@/components/bookmark/TopDashboardItem'

export const TopDashboard = () => {
  const theme = useTheme()

  const styles = {
    main: css({
      margin: '20px',
      padding: '10px',
      color: theme.color.main,
    }),
    h2: css(theme.heading),
    list: css({
      padding: '10px',
      listStyleType: 'none',
    }),
  }

  return (
    <section css={styles.main}>
      <h2 css={styles.h2}>Bookmarks</h2>
      <ul css={styles.list}>
        <TopDashboardItem title='aa' href='https://example.com' />
      </ul>
    </section>
  )
} 
