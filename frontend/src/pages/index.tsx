import { Header } from '@/components/common/Header'
import { ListTeas } from '@/components/top/ListTeas'
import { Container } from '@radix-ui/themes'

export default function Page() {
  return (
    <>
      <Header />
      <Container size='4' m='3' pt='3'>
        <ListTeas />
      </Container>
    </>
  )
}
