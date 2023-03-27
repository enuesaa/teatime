import { css, useTheme } from '@emotion/react'
import Link from 'next/link'

type Props = {
  title: string,
  id: string,
}
export const TopDashboardItem = ({ title, id }: Props) => {
  const theme = useTheme()

  const styles = {
    link: css({
      border: 'solid 1px rgba(255,255,255,0.2)',
      borderRadius: '10px',
      padding: '10px 20px',
      fontSize: theme.fontSize.large,
      color: theme.color.main,
      '&:hover': {
        background: theme.color.sub,
      },
    }),
  }
  return (
    <Link href={`/board/items/${id}`} css={styles.link}>
      {title}
    </Link>
  )
}