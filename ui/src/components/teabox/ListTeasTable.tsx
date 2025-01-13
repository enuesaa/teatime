import { DeleteTea } from './DeleteTea'
import { ListTeasTableCode } from './ListTeasTableCode'
import { UpdateTea } from './UpdateTea'
import { useGetTeapodInfo } from '@/lib/api/teapods'
import { useListTeas } from '@/lib/api/teas'
import { useGetTeasFilter } from '@/states/teasfilter'
import { Table } from '@radix-ui/themes'

export const ListTeasTable = () => {
  const filter = useGetTeasFilter()
  const info = useGetTeapodInfo(filter.teapod)
  const teas = useListTeas(filter.teapod, filter.teabox)
  const teabox = info.data?.teaboxes.find((v) => v.name === filter.teabox)
  const inputs = info.data?.teaboxes.find((v) => v.name === filter.teabox)?.inputs ?? []

  if (teabox === undefined) {
    return <></>
  }

  return (
    <Table.Root variant='surface'>
      <Table.Header>
        <Table.Row>
          <Table.ColumnHeaderCell width='20%'>Id</Table.ColumnHeaderCell>
          {inputs.map(v => (
            <Table.ColumnHeaderCell key={v.name}>{v.name.toUpperCase()}</Table.ColumnHeaderCell>
          ))}
          <Table.ColumnHeaderCell width='20%'>Updated</Table.ColumnHeaderCell>
          <Table.ColumnHeaderCell width='5%'></Table.ColumnHeaderCell>
          <Table.ColumnHeaderCell width='5%'></Table.ColumnHeaderCell>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {teas.data !== undefined &&
          teas?.data.items?.map((tea, i) => (
            <Table.Row key={i}>
              <Table.RowHeaderCell>{tea.id}</Table.RowHeaderCell>
              {inputs.map(v => (
                <Table.Cell key={v.name}>
                  <ListTeasTableCode data={tea.data[v.name] ?? ''} />
                </Table.Cell>
              ))}
              <Table.Cell>{tea.updated}</Table.Cell>
              <Table.Cell>
                <UpdateTea teaId={tea.id} />
              </Table.Cell>
              <Table.Cell>
                <DeleteTea teaId={tea.id} />
              </Table.Cell>
            </Table.Row>
          ))}
      </Table.Body>
    </Table.Root>
  )
}
