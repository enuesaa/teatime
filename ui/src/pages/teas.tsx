import { Header } from '@/components/common/Header'
import { ListTeas } from '@/components/top/ListTeas'
import { Container } from '@radix-ui/themes'
import { useParams } from 'react-router-dom';

export default function Page() {
  const { teapod, teabox } = useParams()
  if (teapod === null || teabox === null) {
    return (<></>)
  }

  return (
    <>
      <Header />
      <Container size='4' m='3' pt='3'>
        <ListTeas />
      </Container>
    </>
  )
}
