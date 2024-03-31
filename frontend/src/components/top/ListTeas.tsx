import { Heading, Text, Table } from '@radix-ui/themes'
import { useListTeas } from '@/lib/api'
import { TeaInfo } from './TeaInfo'
import { DeleteTeaModal } from './DeleteTeaModal'
import { AddTeaModal } from './AddTeaModal'

export const ListTeas = () => {
  const { data: teas } = useListTeas()

  return (
    <>
      <Heading as='h3'>Teas <AddTeaModal /></Heading>
      <Text as='p' size='4' mt='2' mb='6' color='gray'></Text>

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
          {teas !== undefined && teas?.data?.map((tea, i) => (
            <Table.Row key={i}>
              <Table.RowHeaderCell>{tea.id}</Table.RowHeaderCell>
              <Table.Cell><TeaInfo rid={tea.id} /></Table.Cell>
              <Table.Cell>{/*<EditPluginModal id={p.id} />*/}</Table.Cell>
              <Table.Cell><DeleteTeaModal rid={tea.id} /></Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table.Root>
    </>
  )
}