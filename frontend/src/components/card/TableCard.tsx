import { Inset, Table } from '@radix-ui/themes'
import { BaseCard } from './BaseCard'
import { useGetProviderPanel } from '@/lib/api'
import { type TablePanelRecordValue } from '@/lib/schema'

type CellValueProps = {
  value: TablePanelRecordValue
}
const CellValue = ({value}: CellValueProps) => {
  // tablePanel.records をみて api を呼ぶ
  // tablePanel.records[].values[].readonly が false なら編集画面が出てきてデータを更新できる
  // tablePanel.creation.enable が true なら create button と name の入力フォームがありデータを作成できる 

  return (<>{value.key}</>)
}

export const TableCard = () => {
  const {data, isLoading} = useGetProviderPanel('pinit', 'main')
  console.log(data)

  return (
    <BaseCard>
      <Inset clip='padding-box'>
        <Table.Root>
          <Table.Header>
            <Table.Row>
              {data?.tablePanel.head.map((v, i) => (
                <Table.ColumnHeaderCell key={i}>{v}</Table.ColumnHeaderCell>
              ))}
            </Table.Row>
          </Table.Header>

          <Table.Body>
            {data?.tablePanel.records.map((r, i) => (
              <Table.Row key={i}>
                {r.values.map((v, k) => (
                  <Table.Cell key={k}><CellValue value={v} /></Table.Cell>
                ))}
              </Table.Row>
            ))}
          </Table.Body>
        </Table.Root>
      </Inset>
    </BaseCard>
  )
}
