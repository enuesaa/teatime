import { Header } from '@/components/common/Header'
import { Container } from '@radix-ui/themes'
import { Link } from '@/components/common/Link'

export default function Page() {
  return (
    <>
      <Header />
      <Container size='4' m='3' pt='3'>
        <Link href='/teapods/links/teas/links'>LINKS</Link>
      </Container>
    </>
  )
}
