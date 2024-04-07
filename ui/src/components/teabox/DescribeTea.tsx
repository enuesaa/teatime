import { useGetTea } from '@/lib/api/teas'
import { Code } from '@radix-ui/themes'

type Props = {
  teapod: string
  teabox: string
  teaid: string
}
export const DescribeTea = ({ teapod, teabox, teaid }: Props) => {
  const { data: info } = useGetTea(teapod, teabox, teaid)

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
