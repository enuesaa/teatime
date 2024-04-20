import { useGetTeapodCard } from '@/lib/api/teapods'
import { Card, DataList } from '@radix-ui/themes'

type Props = {
  teapod: string
  name: string
}
export const TeapodCard = ({ teapod, name }: Props) => {
  const card = useGetTeapodCard(teapod, name)

  return (
    <Card style={{ width: '300px' }} mb='5'>
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
  )
}
