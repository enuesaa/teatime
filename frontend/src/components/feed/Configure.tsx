import { PageTitle } from '@/components/common/PageTitle'
import Link from 'next/link'
import { FaPlus } from 'react-icons/fa'
import { useStyles } from '@/styles/use'
import { useListFeedsQuery } from '@/lib/feed'

type ItemProps = {
  id: string;
  name: string;
}
const Item = ({ id, name }: ItemProps) => {
  const styles = useStyles(theme => ({
    link: theme({ around: 'x2', decorate: 'card' }),
  }))

  return (
    <Link href={`/feeds/${id}`} css={styles.link}>
      {name}
    </Link>
  )
}

export const Configure = () => {
  const feeds = useListFeedsQuery({})
  const styles = useStyles(theme => ({
    main: theme({ around: 'x1tb' }),
  }))

  return (
    <section css={styles.main}>
      <PageTitle title='Feed'>
        <Link href='/feeds?create'><FaPlus /></Link>
      </PageTitle>
      <ul>
        {feeds?.items.map((v,i) => (
          <Item id={v.id} name={v?.name} key={i} />
        ))}
      </ul>
    </section>
  )
}