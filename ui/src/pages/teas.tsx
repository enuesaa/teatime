import { ListTeas } from '@/components/teabox/ListTeas'
import { useQueryStr } from '@/lib/navigation/querystr'
import { useParams } from 'react-router-dom'

export default function Page() {
  const { teapod } = useParams()
  const teabox = useQueryStr('teabox')
  if (teapod === null || teapod === undefined) {
    return <></>
  }

  return <ListTeas teapod={teapod} teabox={teabox} />
}
