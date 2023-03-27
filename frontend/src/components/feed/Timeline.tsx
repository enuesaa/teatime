import { css, useTheme } from '@emotion/react'
import { TimelineItem } from '@/components/feed/TimelineItem'

export const Timeline = () => {
  const theme = useTheme()

  const styles = {
    main: css({
      margin: '20px',
      padding: '10px',
      color: theme.color.main,
    }),
    h2: css(theme.heading),
    list: css({
      listStyleType: 'none',
    }),
  }

  return (
    <section css={styles.main}>
      <h2 css={styles.h2}>Feeds</h2>
      <ul css={styles.list}>
        <TimelineItem href='https://example.com/' title='aaa' />
      </ul>
    </section>
  )
}