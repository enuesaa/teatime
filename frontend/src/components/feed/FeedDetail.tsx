import { css, useTheme } from '@emotion/react'
import { TimelineItem } from '@/components/feed/TimelineItem'

export const FeedDetail = () => {
  const theme = useTheme()

  const styles = {
    main: css({
      margin: '20px',
      padding: '0 10px 10px 10px',
      color: theme.color.main,
    }),
    h2: css(theme.heading, {
      padding: '0 0 0 10px',
    }),
    list: css({
      listStyleType: 'none',
      padding: '0',
    }),
  }

  return (
    <section css={styles.main}>
      <h2 css={styles.h2}>
        Feed Detail
      </h2>
      <ul css={styles.list}>
        <TimelineItem href='https://example.com/' title='aaa' />
      </ul>
    </section>
  )
}