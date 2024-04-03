import { BiCoffee } from 'react-icons/bi'
import { MdOutlineCoffeeMaker } from 'react-icons/md'
import { Flex, Box, Container, Separator, Link } from '@radix-ui/themes'
import styles from './Header.css'

export const Header = () => {
  return (
    <header className={styles.main}>
      <Container size='4'>
        <Flex>
          <Box flexGrow='1'>
            <Link href='/' className={styles.heading}><BiCoffee />teatime</Link>
          </Box>
          <Link href='/settings' className={styles.setting}><MdOutlineCoffeeMaker /></Link>
        </Flex>
      </Container>
      <Separator size='4' mt='2' />
    </header>
  )
}
