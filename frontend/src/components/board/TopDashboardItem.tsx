import Link from 'next/link'
import { useStyles } from '@/styles/use'

type Props = {
  title: string,
  id: string,
}
export const TopDashboardItem = ({ title, id }: Props) => {
  const styles = useStyles(theme => ({
    link: theme({ size:'x1', around: 'x2' }).css({
      border: 'solid 1px rgba(255,255,255,0.2)',
      borderRadius: '10px',
    }),
  }))

  return (
    <Link href={`/boards/${id}`} css={styles.link}>
      {title}
    </Link>
  )
}