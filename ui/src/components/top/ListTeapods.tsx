import { DescribeTeapod } from './DescribeTeapod'
import { useListTeapods } from '@/lib/api/teapods'

export const ListTeapods = () => {
  const teapods = useListTeapods()

  return <>{teapods.data?.items.map((v, i) => <DescribeTeapod key={i} teapod={v.name} />)}</>
}
