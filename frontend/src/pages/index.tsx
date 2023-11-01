import { Header } from '@/components/common/Header'
import { Container } from '@radix-ui/themes'
import { useQuery } from 'react-query'

type Cards = {
  Items: any
}

const fetchCards = async (): Promise<Cards> => {
  const res = await fetch('/cards')
  const body = await res.json()
  return body as Cards
}

export default function TopPage() {
  const { data, error, isLoading } = useQuery('posts', fetchCards);
  console.log(data)

  return (
    <>
      <Header />
      <Container size='4'>

      </Container>
    </>
  )
}
