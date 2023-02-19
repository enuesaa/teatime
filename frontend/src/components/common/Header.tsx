import Link from 'next/link'
import { css, useTheme } from '@emotion/react'

export function Header() {
  const theme = useTheme()

  const styles = {
    top: css({
      width: '100%',
      height: '1.0',
      minHeight: '100px',
      background: theme.color.highlight,
      boxShadow: '2px 2px 2px rgba(0, 0, 0, 0.7)',
      display: 'flex',
      justifyContent: 'center',
      alignItems: 'center',
    }),
    title: css({
      color: '#fafafa',
      textShadow: '2px 2px 2px #000',
      fontSize: '45px',
      height: '100px',
      lineHeight: '100px',
      fontWeight: '800',
      margin: '0 auto',
      textAlign: 'center',
      userSelect: 'none',
      '&:hover': {
        textShadow: '3px 3px 2px #000',
      },
    }),
  }

  return (
    <>
      <header css={styles.top}>
        <Link href={{ pathname: `/` }}>
          <div css={styles.title}>teatime-app</div>
        </Link>
      </header>
    </>
  )
}
