import { css, useTheme } from '@emotion/react'
import Link from 'next/link'

type Props = {
  title: string,
  href: string,
}
export const TimelineItem = ({ title, href }: Props) => {
  const theme = useTheme()

  const styles = {
    li: css({
      padding: '10px',
      border: 'solid 1px rgba(255,255,255,0.2)',
    })
  }

  return (
    <li css={styles.li}>
      <Link href={href}>{title}</Link>
    </li>
  )
}