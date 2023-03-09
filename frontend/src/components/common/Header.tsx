import Link from 'next/link'
import { css, useTheme } from '@emotion/react'
import { FaCrosshairs } from 'react-icons/fa'

export const Header = () => {
  const theme = useTheme()

  const styles = {
    top: css(theme.box, {
      boxShadow: '2px 2px 2px rgba(0, 0, 0, 0.7)',
    }),
    title: css(theme.heading, {
      color: '#fafafa',
    }),
    settingLink: css({
      padding: '10px',
      fontSize: theme.fontSize.x2large,
      color: '#fafafa',
      position: 'absolute',
      right: '10px',
      top: '15px',
    }),
  }

  return (
    <header css={styles.top}>
      <Link href={{ pathname: `/` }} css={styles.title}>
        teatime-app
      </Link>
      <Link href={{ pathname: `/setting` }} css={styles.settingLink}>
        <FaCrosshairs />
      </Link>
    </header>
  )
}
