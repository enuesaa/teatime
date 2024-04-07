import { useListTeapods } from '@/lib/api/teapods'
import { Card } from '@radix-ui/themes'
import { Link } from '@/components/common/Link'

export const ListTeapods = () => {
  const teapods = useListTeapods()

  return (
    <>
      {teapods.data?.map((v, i) => (
        <Link key={i} href='/teapods/links/teas' asChild>
          <Card variant='surface' style={{width:'100px', textAlign: 'center'}}>
            {v.name}
          </Card>
        </Link>
      ))}
    </>
  )
}
