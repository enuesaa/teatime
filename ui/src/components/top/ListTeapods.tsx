import { useListTeapods } from '@/lib/api/teapods'
import { Card } from '@radix-ui/themes'

export const ListTeapods = () => {
  const teapods = useListTeapods()

  return (
    <>
      {teapods.data?.map((v, i) => (
        <Card key={i} variant='surface' style={{width:'100px', textAlign: 'center'}}>
          {v.name}
        </Card>
      ))}
    </>
  )
}
