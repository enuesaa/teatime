import Link from 'next/link'
import { BiCoffee } from 'react-icons/bi'
import { FaSearch } from 'react-icons/fa'
import { MdOutlineCoffeeMaker } from 'react-icons/md'
import { Flex, Box, Container } from '@radix-ui/themes'

export const Header = () => {
  return (
    <header style={{ height: '50px', lineHeight: '50px', fontSize: '25px', fontWeight: 'bold' }}>
      <Container size='4'>
        <Flex>
          <Box grow='1'>
            <Link href={{ pathname: `/` }} style={{ color: 'white', margin: '0 10px', textDecoration: 'none' }}>
              <BiCoffee style={{ verticalAlign: 'middle' }} />
              teatime
            </Link>
            <input type='text' placeholder='search...' />
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
