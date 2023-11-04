import { Inset, Table } from '@radix-ui/themes'
import { BaseCard } from './BaseCard'
import { useGetProviderPanel } from '@/lib/api'

export const TableCard = () => {
  const {data, isLoading} = useGetProviderPanel('pinit', 'main')
  console.log(data)

  // tablePanel.records をみて api を呼ぶ
  // tablePanel.records[].values[].readonly が false なら編集画面が出てきてデータを更新できる
  // tablePanel.creation.enable が true なら create button と name の入力フォームがありデータを作成できる 

  return (
    <BaseCard>
      <Inset clip='padding-box'>
        <Table.Root>
          <Table.Header>
            <Table.Row>
              <Table.ColumnHeaderCell>{data?.tablePanel.head[0]}</Table.ColumnHeaderCell>
              <Table.ColumnHeaderCell>{data?.tablePanel.head[1]}</Table.ColumnHeaderCell>
              <Table.ColumnHeaderCell>{data?.tablePanel.head[2]}</Table.ColumnHeaderCell>
            </Table.Row>
          </Table.Header>

          <Table.Body>
            <Table.Row>
              <Table.RowHeaderCell>1</Table.RowHeaderCell>
              <Table.Cell>aaa</Table.Cell>
              <Table.Cell>something something something</Table.Cell>
            </Table.Row>
          </Table.Body>
        </Table.Root>
      </Inset>
    </BaseCard>
  )
}
