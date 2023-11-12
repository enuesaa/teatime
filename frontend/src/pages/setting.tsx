import { Header } from '@/components/common/Header'
import { Container, Heading, Text, Table } from '@radix-ui/themes'
import { AddPluginModal } from '@/components/setting/AddPluginModal'
import { EditPluginModal } from '@/components/setting/EditPluginModal'
import { DeletePluginModal } from '@/components/setting/DeletePluginModal'

export default function SettingPage() {

  return (
    <>
      <Header />
      <Container size='4'>
        <Heading as='h2' size='8'>Setting</Heading>
        <Text as='p' size='4' mt='2' mb='6' color='gray'>teatime settings</Text>

        <Heading as='h3'>Plugin <AddPluginModal /></Heading>
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
            <Table.Row>
              <Table.RowHeaderCell>pinit</Table.RowHeaderCell>
              <Table.Cell>./plugins/teatime-plugin-pinit</Table.Cell>
              <Table.Cell>aa</Table.Cell>
              <Table.Cell><EditPluginModal /></Table.Cell>
              <Table.Cell><DeletePluginModal /></Table.Cell>
            </Table.Row>

            <Table.Row>
              <Table.RowHeaderCell>pinit</Table.RowHeaderCell>
              <Table.Cell>./plugins/teatime-plugin-pinit</Table.Cell>
              <Table.Cell>aa</Table.Cell>
              <Table.Cell><EditPluginModal /></Table.Cell>
              <Table.Cell><DeletePluginModal /></Table.Cell>
            </Table.Row>
          </Table.Body>
        </Table.Root>

      </Container>
    </>
  )
}
