import { useListLogs } from '@/lib/api/logs'
import { Table } from '@radix-ui/themes'

export const ListLogs = () => {
  const logs = useListLogs()

  return (
    <Table.Root variant='surface'>
      <Table.Header>
        <Table.Row>
          <Table.ColumnHeaderCell width='10%'>Created</Table.ColumnHeaderCell>
          <Table.ColumnHeaderCell>Message</Table.ColumnHeaderCell>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {logs !== undefined &&
          logs.data?.map((l, i) => (
            <Table.Row key={i}>
              <Table.RowHeaderCell>{l.created}</Table.RowHeaderCell>
              <Table.Cell>{l.message}</Table.Cell>
            </Table.Row>
          ))}
      </Table.Body>
    </Table.Root>
  )
}
