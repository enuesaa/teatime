import { AddTea } from './AddTea'
import { ListTeasCtl } from './ListTeasCtl'
import { ListTeasTable } from './ListTeasTable'
import { Heading, Text, Flex, Box } from '@radix-ui/themes'

type Props = {
  teapod: string
  teabox: string
  teaboxes: string[]
}
export const ListTeas = ({ teapod, teabox, teaboxes }: Props) => {
  return (
    <>
      <Heading as='h3'>
        <Flex>
          <Box width='500px'>{teapod} <AddTea teapod={teapod} teabox={teabox} /></Box>
          <Box flexGrow='1' flexShrink='1'>
            {teaboxes.length > 0 && <ListTeasCtl teapod={teapod} teabox={teabox} teaboxes={teaboxes} />}
          </Box>
        </Flex>
      </Heading>
      <Text as='p' size='4' mt='2' mb='2' color='gray'></Text>
      <ListTeasTable teapod={teapod} teabox={teabox} />
    </>
  )
}
