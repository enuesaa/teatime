import { Header } from '@/components/common/Header'
import { ListTeas } from '@/components/teabox/ListTeas'
import { Container } from '@radix-ui/themes'
import { useParams, useSearchParams } from 'react-router-dom'

export default function Page() {
  const { teapod } = useParams()
  const [searchParams] = useSearchParams()
  if (teapod === null || teapod === undefined) {
    return <></>
  }
  const teabox = searchParams.get('teabox')

  return (
    <>
      <Header />
      <Container size='4' m='3' pt='3'>
        <ListTeas teapod={teapod} teabox={teabox} />
      </Container>
    </>
  )
}
