import { Heading, Text, Table } from '@radix-ui/themes'
import { AddPluginModal } from '@/components/setting/AddPluginModal'
import { EditPluginModal } from '@/components/setting/EditPluginModal'
import { DeletePluginModal } from '@/components/setting/DeletePluginModal'
import { useListProviders } from '@/lib/api'
import { PluginInfo } from './PluginInfo'

export const ListProviders = () => {
  const { data: providers } = useListProviders()

  return (
    <>
      <Heading as='h3'>Plugins <AddPluginModal /></Heading>
      <Text as='p' size='4' mt='2' mb='6' color='gray'></Text>

      <Table.Root variant='surface'>
        <Table.Header>
          <Table.Row>
            <Table.ColumnHeaderCell width='10%'>Name</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell width='20%'>Command</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell>Info</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell width='5%'></Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell width='5%'></Table.ColumnHeaderCell>
          </Table.Row>
        </Table.Header>
        <Table.Body>
          {providers !== undefined && providers.items.map(p => (
            <Table.Row key={p.name}>
              <Table.RowHeaderCell>{p.name}</Table.RowHeaderCell>
              <Table.Cell>{p.command}</Table.Cell>
              <Table.Cell><PluginInfo name={p.name} /></Table.Cell>
              <Table.Cell><EditPluginModal /></Table.Cell>
              <Table.Cell><DeletePluginModal /></Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table.Root>
    </>
  )
}