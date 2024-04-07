import { useListTeas } from '@/lib/api/teas'
import { Table } from '@radix-ui/themes'
import { DescribeTea } from './DescribeTea'
import { DeleteTeaModal } from './DeleteTeaModal'

type Props = {
  teapod: string;
  teabox: string;
}
export const ListTeasTable = ({ teapod, teabox }: Props) => {
  const teas = useListTeas(teapod, teabox)

  return (
    <Table.Root variant='surface'>
      <Table.Header>
        <Table.Row>
          <Table.ColumnHeaderCell width='30%'>Id</Table.ColumnHeaderCell>
          <Table.ColumnHeaderCell>Value</Table.ColumnHeaderCell>
          <Table.ColumnHeaderCell width='5%'></Table.ColumnHeaderCell>
          <Table.ColumnHeaderCell width='5%'></Table.ColumnHeaderCell>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {teas.data !== undefined &&
          teas?.data?.map((tea, i) => (
            <Table.Row key={i}>
              <Table.RowHeaderCell>{tea.id}</Table.RowHeaderCell>
              <Table.Cell>
                <DescribeTea teapod={teapod} teaid={tea.id} />
              </Table.Cell>
              <Table.Cell></Table.Cell>
              <Table.Cell>
                <DeleteTeaModal teapod={teapod} teaid={tea.id} />
              </Table.Cell>
            </Table.Row>
          ))}
      </Table.Body>
    </Table.Root>
  )
}
