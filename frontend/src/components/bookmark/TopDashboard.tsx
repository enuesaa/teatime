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

export const TopDashboard = () => {
  const list = useListBookmarksQuery({})

  const styles = useStyles(theme => ({
    h2: theme().css({
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
      <PageTitle title='Bookmark'>
        <Link href='/bookmarks'><AiOutlineSwapRight /></Link>
      </PageTitle>
      <ul css={styles.list}>
        {
          list?.items.map((v,i) => {
            return (<Item title={v?.name} href={v?.url ?? 'https://example.com/'} key={i} />)
          })
        }
      </ul>
    </section>
  )
} 
