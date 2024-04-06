import { Heading, Text, Table } from '@radix-ui/themes'
import { useListTeas } from '@/lib/api/teas'
import { DescribeTea } from './DescribeTea'
import { DeleteTeaModal } from './DeleteTeaModal'
import { AddTeaModal } from './AddTeaModal'

export const ListTeas = () => {
  const { data: teas } = useListTeas()

  return (
    <>
      <Heading as='h3'>
        Teas <AddTeaModal />
      </Heading>
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
          {teas !== undefined &&
            teas?.data?.map((tea, i) => (
              <Table.Row key={i}>
                <Table.RowHeaderCell>{tea.id}</Table.RowHeaderCell>
                <Table.Cell>
                  <DescribeTea teaid={tea.id} />
                </Table.Cell>
                <Table.Cell></Table.Cell>
                <Table.Cell>
                  <DeleteTeaModal teaid={tea.id} />
                </Table.Cell>
              </Table.Row>
            ))}
        </Table.Body>
      </Table.Root>
    </>
  )
}
