import Link from 'next/link'
import { css, useTheme } from '@emotion/react'

export const Header = () => {
  const theme = useTheme()

  const styles = {
    top: css(theme.box, {
      boxShadow: '2px 2px 2px rgba(0, 0, 0, 0.7)',
    }),
    title: css(theme.heading, {
      color: '#fafafa',
    }),
  }

  return (
    <header css={styles.top}>
      <Link href={{ pathname: `/` }} css={styles.title}>
        teatime-app
      </Link>
    </header>
  )
}
