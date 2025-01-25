import { ActionMenu } from './ActionMenu'
import { AddTea } from './AddTea'
import { ListTeasCtl } from './ListTeasCtl'
import { ListTeasTable } from './ListTeasTable'
import { useInitTeasFilter } from '@/states/teasfilter'
import { Heading, Text, Flex, Box } from '@radix-ui/themes'

type Props = {
  teapod: string
  teabox?: string
}
export const ListTeas = ({ teapod, teabox }: Props) => {
  const init = useInitTeasFilter(teapod, teabox)
  if (!init.ok) {
    return <></>
  }

  return (
    <>
      <Heading as='h3'>
        <Flex>
          <Box width='500px'>
            {teapod} <AddTea /> <ActionMenu teapod={teapod} />
          </Box>
          <Box flexGrow='1' flexShrink='1'>
            <ListTeasCtl />
          </Box>
        </Flex>
      </Heading>
      <Text as='p' size='4' mt='2' mb='2' color='gray'></Text>
      <ListTeasTable />
    </>
  )
}
