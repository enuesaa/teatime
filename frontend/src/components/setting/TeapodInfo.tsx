import { useGetTeapodInfo } from '@/lib/api/teapods'
import { Code } from '@radix-ui/themes'

type Props = {
  name: string
}
export const TeapodInfo = ({ name }: Props) => {
  const { data: info } = useGetTeapodInfo(name)

  return (
    <>
      {info !== undefined && (
        <pre>
          <Code color='gray' style={{ display: 'block', padding: '10px' }}>
            {JSON.stringify(info?.data, null, '  ')}
          </Code>
        </pre>
      )}
    </>
  )
}
