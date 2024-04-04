import { Heading, Text, Table } from '@radix-ui/themes'
import { useListTeapods } from '@/lib/api/teapods'
import { ProviderInfo } from './ProviderInfo'

export const ListProviders = () => {
  const { data: teapods } = useListTeapods()

  return (
    <>
      <Heading as='h3'>Plugins</Heading>
      <Text as='p' size='4' mt='2' mb='6' color='gray'></Text>

      <Table.Root variant='surface'>
        <Table.Header>
          <Table.Row>
            <Table.ColumnHeaderCell width='10%'>Name</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell width='20%'>Command</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell>Info</Table.ColumnHeaderCell>
          </Table.Row>
        </Table.Header>
        <Table.Body>
          {teapods !== undefined &&
            teapods?.data?.map((p) => (
              <Table.Row key={p.id}>
                <Table.RowHeaderCell>{p.name}</Table.RowHeaderCell>
                <Table.Cell>{p.command}</Table.Cell>
                <Table.Cell>
                  <ProviderInfo id={p.id} />
                </Table.Cell>
              </Table.Row>
            ))}
        </Table.Body>
      </Table.Root>
    </>
  )
}
