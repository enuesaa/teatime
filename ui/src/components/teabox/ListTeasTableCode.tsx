import { Code } from '@radix-ui/themes'

type Props = {
  data: Record<string, any>|string
}
export const ListTeasTableCode = ({ data }: Props) => {
  if (typeof data === 'string') {
    return (
      <Code color='gray' style={{ display: 'block', padding: '7px 7px 0 7px', minHeight: '28px' }}>
        {data}
      </Code>
    )
  }

  return (
    <Code color='gray' style={{ display: 'block', padding: '7px 7px 0 7px', minHeight: '28px' }}>
      {JSON.stringify(data, null, '  ')}
    </Code>
  )
}
