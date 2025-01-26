import { Link } from '@/components/common/Link'
import { useGetTeapodInfo } from '@/lib/api/teapods'
import { Button, Heading } from '@radix-ui/themes'
import styles from './DescribeTeapod.css'

type Props = {
  teapod: string
}
export const DescribeTeapod = ({ teapod }: Props) => {
  const info = useGetTeapodInfo(teapod)
  const teaboxes = info.data?.teaboxes.map((v) => v.name) ?? []

  return (
    <div className={styles.main}>
      <Heading className={styles.heading}>{teapod}</Heading>

      {teaboxes.map((v, i) => (
        <Link href={`/teapods/${teapod}/teas?teabox=${v}`} asChild key={i}>
          <Button color='gray' variant='outline' highContrast m='2'>
            {v}
          </Button>
        </Link>
      ))}
    </div>
  )
}
