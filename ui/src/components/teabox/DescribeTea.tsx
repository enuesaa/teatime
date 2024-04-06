import { useGetTea } from '@/lib/api/teas'
import { Code } from '@radix-ui/themes'

type Props = {
  teaid: string
}
export const DescribeTea = ({ teaid }: Props) => {
  const { data: info } = useGetTea(teaid)

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
