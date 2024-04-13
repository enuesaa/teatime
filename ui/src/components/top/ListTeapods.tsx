import { useListTeapods } from '@/lib/api/teapods'
import { DescribeTeapod } from './DescribeTeapod'

export const ListTeapods = () => {
  const teapods = useListTeapods()

  return (
    <>
      {teapods.data?.map((v, i) => (
        <DescribeTeapod key={i} teapod={v.name} />
      ))}
    </>
  )
}
