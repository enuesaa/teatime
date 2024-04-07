import { Heading, Text, Table, SegmentedControl, Flex, Box } from '@radix-ui/themes'
import { useListTeas } from '@/lib/api/teas'
import { DescribeTea } from './DescribeTea'
import { DeleteTeaModal } from './DeleteTeaModal'
import { AddTeaModal } from './AddTeaModal'
import { FormEventHandler, useState } from 'react'

type Props = {
  teapod: string
}
export const ListTeas = ({ teapod }: Props) => {
  const teas = useListTeas(teapod)
  const [teabox, setTeabox] = useState<string>('links')

  const handleTeaboxChange = (value: string) => setTeabox(value)

  return (
    <>
      <Heading as='h3'>
        <Flex>
          <Box width='500px'>
            Teas <AddTeaModal teapod={teapod} />
          </Box>
          <Box flexGrow='1' flexShrink='1'>
            <SegmentedControl.Root defaultValue='links' onValueChange={handleTeaboxChange} radius='full' variant='classic' size='3'>
              <SegmentedControl.Item value='links'>Links</SegmentedControl.Item>
              <SegmentedControl.Item value='notes'>Notes</SegmentedControl.Item>
            </SegmentedControl.Root>
          </Box>
        </Flex>
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
          {teas.data !== undefined &&
            teas?.data?.map((tea, i) => (
              <Table.Row key={i}>
                <Table.RowHeaderCell>{tea.id}</Table.RowHeaderCell>
                <Table.Cell>
                  <DescribeTea teapod={teapod} teaid={tea.id} />
                </Table.Cell>
                <Table.Cell></Table.Cell>
                <Table.Cell>
                  <DeleteTeaModal teapod={teapod} teaid={tea.id} />
                </Table.Cell>
              </Table.Row>
            ))}
        </Table.Body>
      </Table.Root>
    </>
  )
}
