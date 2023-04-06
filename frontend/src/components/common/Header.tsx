import Link from 'next/link'
import { css, useTheme } from '@emotion/react'
import { BiCoffee } from 'react-icons/bi'
import { FaSearch } from 'react-icons/fa'
import { MdOutlineCoffeeMaker } from 'react-icons/md'

export const Header = () => {
  const theme = useTheme()

  const styles = {
    top: css(theme.box, {
      boxShadow: '2px 2px 2px rgba(0, 0, 0, 0.7)',
      padding: '0',
      height: '50px',
      display: 'flex',
    }),
    title: css(theme.heading, {
      flex: '1 0 auto',
      color: '#fafafa',
      padding: '0',
      margin: '0 30px',
      lineHeight: '50px',
      fontSize: theme.fontSize.xlarge,
      'svg': {
        verticalAlign: 'middle',
        color: theme.color.main,
        margin: '0 5px',
        cursor: 'pointer',
      },
    }),
    search: css({
      display: 'inline-block',
      margin: '3px 30px',
      fontSize: theme.fontSize.normal,
      'input': {
        ...theme.input,
        margin: '10px',
        width: '300px',
        padding: '5px',
        borderRadius: '5px',
        display: 'inline-block',
        color: theme.color.contrast,
      },
    }),
    config: css(theme.input, {
      color: theme.color.main,
      fontSize: theme.fontSize.xlarge,
      background: 'rgba(0,0,0,0)',
      lineHeight: '50px',
      cursor: 'pointer',
      'svg': {
        verticalAlign: 'middle',
        margin: '0 10px',
      },
    }),
  }

  return (
    <header css={styles.top}>
      <div css={styles.title}>
        <Link href={{ pathname: `/` }}>
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
