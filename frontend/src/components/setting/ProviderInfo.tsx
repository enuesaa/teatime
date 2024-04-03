import { useGetProviderInfo } from '@/lib/api'
import { Code } from '@radix-ui/themes'

type Props = {
  id: string
}
export const ProviderInfo = ({ id }: Props) => {
  const { data: info } = useGetProviderInfo(id)

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
