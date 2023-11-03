import { Inset, Table } from '@radix-ui/themes'
import { BaseCard } from './BaseCard'

export const TableCard = () => {
  return (
    <BaseCard>
      <Inset clip='padding-box'>
        <Table.Root>
          <Table.Header>
            <Table.Row>
              <Table.ColumnHeaderCell>Id</Table.ColumnHeaderCell>
              <Table.ColumnHeaderCell>Name</Table.ColumnHeaderCell>
              <Table.ColumnHeaderCell>Description</Table.ColumnHeaderCell>
            </Table.Row>
          </Table.Header>

          <Table.Body>
          <Table.Row>
              <Table.RowHeaderCell>1</Table.RowHeaderCell>
              <Table.Cell>aaa</Table.Cell>
              <Table.Cell>something something something</Table.Cell>
            </Table.Row>
            
            <Table.Row>
              <Table.RowHeaderCell>1</Table.RowHeaderCell>
              <Table.Cell>bbb</Table.Cell>
              <Table.Cell>something something something</Table.Cell>
            </Table.Row>

            <Table.Row>
              <Table.RowHeaderCell>1</Table.RowHeaderCell>
              <Table.Cell>ccc</Table.Cell>
              <Table.Cell>something something something</Table.Cell>
            </Table.Row>
          </Table.Body>
        </Table.Root>
      </Inset>
    </BaseCard>
  )
}
