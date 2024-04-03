import { useGetTeaInfo } from '@/lib/api'
import { Code } from '@radix-ui/themes'

type Props = {
  rid: string
}
export const TeaInfo = ({ rid }: Props) => {
  const { data: info } = useGetTeaInfo(rid)

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
