import Link from 'next/link'
import { useStyles } from '@/styles/use'
import { BiCoffee } from 'react-icons/bi'
import { FaSearch } from 'react-icons/fa'
import { MdOutlineCoffeeMaker } from 'react-icons/md'

export const Header = () => {

  const styles = useStyles(theme => ({
    top: theme({surf: 'main', decorate: 'shadow'}).css({
      height: '50px',
      display: 'flex',
      lineHeight: '50px',
      'svg': {
        verticalAlign: 'middle',
        margin: '0 5px',
        cursor: 'pointer',
      },
    }),
    title: theme({size: 'x3'}).css({
      flex: '1 0 auto',
      margin: '0 30px',
    }),
    search: theme({size: 'x1'}).css({
      display: 'inline-block',
      margin: '0 30px',
      'input': {
        margin: '10px',
        width: '300px',
        padding: '5px',
        borderRadius: '5px',
        display: 'inline-block',
      },
    }),
    config: theme({size: 'x3'}),
  }))

  return (
    <header css={styles.top}>
      <div css={styles.title}>
        <Link href={{ pathname: `/` }} >
          <BiCoffee />
          teatime
        </Link>
        <div css={styles.search}>
          <input type='text' placeholder='search...' />
          <FaSearch />
        </div>
      </div>
      <Link href='/setting' css={styles.config}>
        <MdOutlineCoffeeMaker />
      </Link>
    </header>
  )
}
