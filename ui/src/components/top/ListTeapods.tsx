import { useListTeapods } from '@/lib/api/teapods'
import { Card, Heading } from '@radix-ui/themes'
import { DescribeTeapod } from './DescribeTeapod'

export const ListTeapods = () => {
  const teapods = useListTeapods()

  return (
    <>
      <Heading mb='3'>teapods</Heading>
      {teapods.data?.map((v, i) => (
        <DescribeTeapod key={i} name={v.name} />
      ))}
    </>
  )
}
