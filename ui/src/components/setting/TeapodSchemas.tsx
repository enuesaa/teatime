import { useGetTeapodInfo } from '@/lib/api/teapods'
import { Code } from '@radix-ui/themes'

type Props = {
  name: string
}
export const TeapodSchemas = ({ name }: Props) => {
  const { data: info } = useGetTeapodInfo(name)

  return (
    <>
      {info !== undefined &&
        info.data.schemas.map((s, i) => (
          <pre key={i}>
            <Code color='gray' style={{ display: 'block', padding: '10px' }}>
              {JSON.stringify(s, null, '  ')}
            </Code>
          </pre>
        ))}
    </>
  )
}
