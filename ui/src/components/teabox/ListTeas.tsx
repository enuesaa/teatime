import { AddTea } from './AddTea'
import { ListTeasCtl } from './ListTeasCtl'
import { ListTeasTable } from './ListTeasTable'
import { Heading, Text, Flex, Box } from '@radix-ui/themes'
import { useInitTeasFilter } from '@/states/teasfilter'

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
          <Box width='500px'>{teapod} <AddTea /></Box>
          <Box flexGrow='1' flexShrink='1'><ListTeasCtl /></Box>
        </Flex>
      </Heading>
      <Text as='p' size='4' mt='2' mb='2' color='gray'></Text>
      <ListTeasTable />
    </>
  )
}
