import { Code } from '@radix-ui/themes'

type Props = {
  data: Record<string, any>
}
export const ListTeasTableCode = ({ data }: Props) => {
  return (
    <Code color='gray' style={{ display: 'block', padding: '10px' }}>
      {JSON.stringify(data, null, '  ')}
    </Code>
  )
}
