import { Code } from '@radix-ui/themes'

type Props = {
  data: Record<string, any>
}
export const ListTeasTableData = ({ data }: Props) => {
  return (
    <Code color='gray' style={{ display: 'block', padding: '10px' }}>
      {JSON.stringify(data, null, '  ')}
    </Code>
  )
}
