import { BiCoffee } from 'react-icons/bi'
import { MdOutlineCoffeeMaker } from 'react-icons/md'
import { Flex, Box, Container, Separator, Link } from '@radix-ui/themes'
import { css } from '@emotion/react'

export const Header = () => {
  const styles = {
    main: css({
      fontSize: '25px',
      fontWeight: 'bold',
      'a': {
        textDecoration: 'none',
      },
    }),
    heading: css({
      color: 'white',
      margin: '0 10px',
      'svg': {
        verticalAlign: 'middle',
        margin: '0 10px',
      },
    }),
    setting: css({
      color: 'white',
      'svg': {
        verticalAlign: 'middle',
      },
    }),
  }

  return (
    <header css={styles.main}>
      <Container size='4'>
        <Flex>
          <Box grow='1'>
            <Link href='/' css={styles.heading}><BiCoffee />teatime</Link>
          </Box>
          <Link href='/settings' css={styles.setting}><MdOutlineCoffeeMaker /></Link>
        </Flex>
      </Container>
      <Separator size='4' mt='2' />
    </header>
  )
}
