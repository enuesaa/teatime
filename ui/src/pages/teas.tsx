import { Layout } from '@/components/common/Layout'
import { ListTeas } from '@/components/teabox/ListTeas'
import { useParams, useSearchParams } from 'react-router-dom'

export default function Page() {
  const { teapod } = useParams()
  const [searchParams] = useSearchParams()
  if (teapod === null || teapod === undefined) {
    return <></>
  }
  const teabox = searchParams.get('teabox')

  return (
    <Layout>
      <ListTeas teapod={teapod} teabox={teabox} />
    </Layout>
  )
}
