import { useGetTeapodCard, useListTeapods } from '@/lib/api/teapods'
import { DescribeTeapod } from './DescribeTeapod'
import { Card, DataList } from '@radix-ui/themes'

export const ListTeapods = () => {
  const teapods = useListTeapods()
  const card = useGetTeapodCard('links', 'summary')

  return (
    <>
      <Card style={{width: '300px'}} mb='5'>
        <DataList.Root>
          <DataList.Item>
            <DataList.Label>name</DataList.Label>
            <DataList.Value>{card.data?.name}</DataList.Value>
          </DataList.Item>
          <DataList.Item>
            <DataList.Label>title</DataList.Label>
            <DataList.Value>{card.data?.title}</DataList.Value>
          </DataList.Item>
          <DataList.Item>
            <DataList.Label>text</DataList.Label>
            <DataList.Value>{card.data?.text}</DataList.Value>
          </DataList.Item>
        </DataList.Root>
      </Card>
      {teapods.data?.map((v, i) => (
        <DescribeTeapod key={i} teapod={v.name} />
      ))}
    </>
  )
}
