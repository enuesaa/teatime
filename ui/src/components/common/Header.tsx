import styles from './Header.css'
import { Link } from './Link'
import { Flex, Box, Container, Separator } from '@radix-ui/themes'
import { BiCoffee } from 'react-icons/bi'
import { MdOutlineCoffeeMaker } from 'react-icons/md'
import { TbLogs } from 'react-icons/tb'

export const Header = () => {
  return (
    <header className={styles.main}>
      <Container size='4'>
        <Flex>
          <Box flexGrow='1'>
            <Link href='/' className={styles.heading}>
              <BiCoffee />
              teatime
            </Link>
          </Box>
          <Link href='/logs' mx='2' className={styles.setting}>
            <TbLogs />
          </Link>
          <Link href='/settings' className={styles.setting}>
            <MdOutlineCoffeeMaker />
          </Link>
        </Flex>
      </Container>
      <Separator size='4' mt='2' />
    </header>
  )
}
