import Link from 'next/link'
import { BiCoffee } from 'react-icons/bi'
import { FaSearch } from 'react-icons/fa'
import { MdOutlineCoffeeMaker } from 'react-icons/md'
import { Flex, Box, Container, TextField } from '@radix-ui/themes'
import { css } from '@emotion/react'

export const Header = () => {
  const styles = {
    main: css({
      height: '50px',
      lineHeight: '50px',
      fontSize: '25px',
      fontWeight: 'bold',
    }),
    search: css({
      display: 'inline-block',
      maxWidth: '200px',
    })
  }

  return (
    <header css={styles.main}>
      <Container size='4'>
        <Flex>
          <Box grow='1'>
            <Link href={{ pathname: `/` }} style={{ color: 'white', margin: '0 10px', textDecoration: 'none' }}>
              <BiCoffee style={{ verticalAlign: 'middle' }} />
              teatime
            </Link>
            <TextField.Input placeholder='search...' css={styles.search} />
            <FaSearch style={{ verticalAlign: 'middle', margin: '0 10px' }} />
          </Box>

          <Link href='/setting' style={{ display: 'block', width: '100px', color: 'white', textDecoration: 'none' }}>
            <MdOutlineCoffeeMaker />
          </Link>
        </Flex>
      </Container>
    </header>
  )
}
