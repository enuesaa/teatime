import { AiOutlineSwapRight } from 'react-icons/ai'
import Link from 'next/link'
import { useListBookmarksQuery } from '@/lib/bookmark'
import { useStyles } from '@/styles/use'
import { PageTitle } from '@/components/common/PageTitle'

type ItemProps = {
  title: string,
  href: string,
}
export const Item = ({ title, href }: ItemProps) => {
  const styles = useStyles(theme => ({
    link: theme({ size:'x1', around: 'x2', decorate: 'card' }),
  }))

  return (
    <Link href={href} css={styles.link} target='_blank'>
      {title}
    </Link>
  )
}

export const Dashboard = () => {
  const list = useListBookmarksQuery({})

  const styles = useStyles(theme => ({
    main: theme({ around: 'x1tb' }),
  }))

  return (
    <section css={styles.main}>
      <PageTitle title='Bookmark'>
        <Link href='/bookmarks'><AiOutlineSwapRight /></Link>
      </PageTitle>
      <ul>
        {list?.items.map((v,i) => (
          <Item title={v?.name} href={v?.url ?? 'https://example.com/'} key={i} />)
        )}
      </ul>
    </section>
  )
} 
