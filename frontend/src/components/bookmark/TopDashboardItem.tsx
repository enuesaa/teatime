import { css, useTheme } from '@emotion/react'
import Link from 'next/link'

type Props = {
  title: string,
  href: string,
}
export const TopDashboardItem = ({ title, href }: Props) => {
  const theme = useTheme()

  const styles = {
    link: css({
      border: 'solid 1px rgba(255,255,255,0.2)',
      borderRadius: '10px',
      padding: '10px 20px',
      fontSize: theme.fontSize.large,
      color: theme.color.main,
    }),
  }
  return (
    <Link href={href} css={styles.link}>
      {title}
    </Link>
  )
}