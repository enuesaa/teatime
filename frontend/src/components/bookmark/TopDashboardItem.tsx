import Link from 'next/link'
import { useStyles } from '@/styles/use'

type Props = {
  title: string,
  href: string,
}
export const TopDashboardItem = ({ title, href }: Props) => {
  const styles = useStyles(theme => ({
    link: theme({ size:'x1', around: 'x2' }).css({
      border: 'solid 1px rgba(255,255,255,0.2)',
      borderRadius: '10px',
    }),
  }))
  return (
    <Link href={href} css={styles.link} target='_blank'>
      {title}
    </Link>
  )
}