import Link from 'next/link'
import { useStyles } from '@/styles/use'

type Props = {
  title: string,
  href: string,
}
export const TimelineItem = ({ title, href }: Props) => {
  const styles = useStyles(theme => ({
    li: theme().css({
      padding: '10px',
      border: 'solid 1px rgba(255,255,255,0.2)',
      a: {
        display: 'block',
        width: '100%',
        height: '100%',
      },
    })
  }))

  return (
    <li css={styles.li}>
      <Link href={href}>{title}</Link>
    </li>
  )
}