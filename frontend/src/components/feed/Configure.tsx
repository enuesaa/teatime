import { PageTitle } from '@/components/common/PageTitle'
import Link from 'next/link'
import { FaPlus } from 'react-icons/fa'
import { useStyles } from '@/styles/use'
import { GiCoffeePot } from 'react-icons/gi'
import { useListFeedsQuery } from '@/lib/feed'

const List = () => {
  const feeds = useListFeedsQuery({})

  const styles = useStyles(theme => ({
    list: theme().css({
      'li': {
        padding: '10px',
        border: 'solid 1px #fafafa',
      },
    }),
  }))

  return (
    <ul css={styles.list}>
      <li>
        {feeds?.items.map((v,i) => (
          <Link href={`/feeds/${v.id}`} key={i}>
            {v?.name}
            <GiCoffeePot />
          </Link>
        ))}
      </li>
    </ul>
  )
}

export const Configure = () => {
  return (
    <>
      <PageTitle title='Feed'>
        <Link href='/bookmarks?create'><FaPlus /></Link>
      </PageTitle>
      <List />
    </>
  )
}