import { useListTeapods } from '@/lib/api/teapods'
import { DescribeTeapod } from './DescribeTeapod'
import { ListCards } from './ListCards'

export const ListTeapods = () => {
  const teapods = useListTeapods()

  return (
    <>
      {teapods.data?.map((v, i) => (
        <ListCards key={i} teapod={v.name} />
      ))}
      {teapods.data?.map((v, i) => (
        <DescribeTeapod key={i} teapod={v.name} />
      ))}
    </>
  )
}
