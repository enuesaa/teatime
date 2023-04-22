import { PageTitle } from '@/components/common/PageTitle'
import Link from 'next/link'
import { FaPlus } from 'react-icons/fa'
import { useStyles } from '@/styles/use'
import { useListBookmarksQuery } from '@/lib/bookmark'

type ItemProps = {
  id: string;
  name: string;
}
const Item = ({ id, name }: ItemProps) => {
  return (
    <Link href={`/bookmarks/${id}`}>
      {name}
    </Link>
  )
}

export const Configure = () => {
  const bookmarks = useListBookmarksQuery({})
  const styles = useStyles(theme => ({
    main: theme({ around: 'x1tb' }),
    list: theme({ around: 'x1' }).css({
      a: theme({ around: 'x1', decorate: 'card' }).to(),
    }),
  }))

  return (
    <section css={styles.main}>
      <PageTitle title='Bookmarks'>
        <Link href='/bookmarks?create'><FaPlus /></Link>
      </PageTitle>
      <div css={styles.list}>
        {bookmarks?.items.map((v,i) => (
          <Item id={v.id} name={v.name} key={i} />
        ))}
      </div>
    </section>
  )
}