import { useListLogs } from '@/lib/api/logs'
import { Table } from '@radix-ui/themes'
import { format } from 'date-fns'

export const ListLogs = () => {
  const logs = useListLogs()

  return (
    <Table.Root variant='surface'>
      <Table.Header>
        <Table.Row>
          <Table.ColumnHeaderCell width='20%'>Created</Table.ColumnHeaderCell>
          <Table.ColumnHeaderCell>Message</Table.ColumnHeaderCell>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {logs !== undefined &&
          logs.data?.items.map((l, i) => (
            <Table.Row key={i}>
              <Table.RowHeaderCell>{format(l.created, 'yyyy-MM-dd HH:mm:ss')}</Table.RowHeaderCell>
              <Table.Cell>{l.message}</Table.Cell>
            </Table.Row>
          ))}
      </Table.Body>
    </Table.Root>
  )
}
