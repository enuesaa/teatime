import { Button, Card } from '@radix-ui/themes'
import { Link } from '../common/Link'

type Props = {
  name: string
}
export const DescribeTeapod = ({ name }: Props) => {
  return (
    <Link href={`/teapods/${name}/teas`} asChild>
      <Button color='gray' variant='outline' highContrast>
        {name}
      </Button>
    </Link>
  )
}
