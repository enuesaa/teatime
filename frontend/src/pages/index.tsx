import { Header } from '@/components/common/Header'
import { Container } from '@radix-ui/themes'
import { useFetchCards } from '@/lib/api';

export default function TopPage() {
  const { data, error, isLoading } = useFetchCards()
  if (isLoading) {
    return (<>a</>)
  }

  return (
    <>
      <Header />
      <Container size='4'>
        {data?.items.map(v => (
          <>{v.layout}{v.rid}</>
        ))}
      </Container>
    </>
  )
}
