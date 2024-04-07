import { Header } from '@/components/common/Header'
import { Container } from '@radix-ui/themes'
import { ListTeapods } from '@/components/top/ListTeapods'

export default function Page() {
  return (
    <>
      <Header />
      <Container size='4' m='3' pt='3'>
        <ListTeapods />
      </Container>
    </>
  )
}
