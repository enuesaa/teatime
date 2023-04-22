import { AiOutlineSwapRight } from 'react-icons/ai'
import Link from 'next/link'
import { useStyles } from '@/styles/use'
import { PageTitle } from '@/components/common/PageTitle'

type ItemProps = {
  title: string,
  id: string,
}
export const Item = ({ title, id }: ItemProps) => {
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

export const TopDashboard = () => {
  const styles = useStyles(theme => ({
    h2: theme({size: 'x3'}).css({
      padding: '0 0 0 10px',
      'a': {
        margin: '10px',
        display: 'inline-block',
      },
    }),
    list: theme().css({
      listStyleType: 'none',
      padding: '0',
    }),
  }))

  return (
    <section>
      <PageTitle title='Boards'>
        <Link href='/boards'><AiOutlineSwapRight /></Link>
      </PageTitle>
      <ul css={styles.list}>
        <Item title='aa' id='aa' />
      </ul>
    </section>
  )
} 
