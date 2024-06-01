import { useGetTeapodInfo } from '@/lib/api/teapods'
import { useListTeas } from '@/lib/api/teas'
import { Table } from '@radix-ui/themes'

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
          <Table.ColumnHeaderCell width='10%'>Id</Table.ColumnHeaderCell>
          {teabox.valdefs.map((v, i) => (
            <Table.ColumnHeaderCell key={i}>{v.name.toUpperCase()}</Table.ColumnHeaderCell>
          ))}
          <Table.ColumnHeaderCell width='5%'></Table.ColumnHeaderCell>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {teas.data !== undefined &&
          teas?.data?.map((tea, i) => (
            <Table.Row key={i}>
              <Table.RowHeaderCell>{tea.teaid}</Table.RowHeaderCell>
              {teabox.valdefs.map((v, i) => (
                <Table.Cell key={i}>{tea.vals.find(t => t.name === v.name)?.value ?? ''}</Table.Cell>
              ))}
              <Table.Cell>
                {/* <DeleteTeaModal teapod={teapod} teaid={tea.teaid} /> */}
              </Table.Cell>
            </Table.Row>
          ))}
      </Table.Body>
    </Table.Root>
  )
}
