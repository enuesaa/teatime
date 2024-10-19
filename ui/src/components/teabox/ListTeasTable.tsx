import { useGetTeapodInfo } from '@/lib/api/teapods'
import { useListTeas } from '@/lib/api/teas'
import { Table } from '@radix-ui/themes'
import { ListTeasTableCode } from './ListTeasTableCode'
import { DeleteTea } from './DeleteTea'
import { UpdateTea } from './UpdateTea'

type Props = {
  teapod: string
  teabox: string
}
export const ListTeasTable = ({ teapod, teabox: teaboxName }: Props) => {
  const info = useGetTeapodInfo(teapod)
  const teas = useListTeas(teapod, teaboxName)
  const teabox = info.data?.teaboxes.find((v) => v.name === teaboxName)

  if (teabox === undefined) {
    return <></>
  }

  return (
    <Table.Root variant='surface'>
      <Table.Header>
        <Table.Row>
          <Table.ColumnHeaderCell width='20%'>Id</Table.ColumnHeaderCell>
          <Table.ColumnHeaderCell>Data</Table.ColumnHeaderCell>
          <Table.ColumnHeaderCell width='5%'></Table.ColumnHeaderCell>
          <Table.ColumnHeaderCell width='5%'></Table.ColumnHeaderCell>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {teas.data !== undefined &&
          teas?.data?.map((tea, i) => (
            <Table.Row key={i}>
              <Table.RowHeaderCell>{tea.id}</Table.RowHeaderCell>
              <Table.Cell><ListTeasTableCode data={tea.data} /></Table.Cell>
              <Table.Cell><UpdateTea teapod={teapod} teabox={teabox.name} teaId={tea.id} /></Table.Cell>
              <Table.Cell><DeleteTea teapod={teapod} teabox={teabox.name} teaId={tea.id} /></Table.Cell>
            </Table.Row>
          ))}
      </Table.Body>
    </Table.Root>
  )
}
