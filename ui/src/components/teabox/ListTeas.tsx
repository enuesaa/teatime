import { useGetTeapodInfo } from '@/lib/api/teapods'
import { AddTea } from './AddTea'
import { ListTeasCtl } from './ListTeasCtl'
import { ListTeasTable } from './ListTeasTable'
import { Heading, Text, Flex, Box } from '@radix-ui/themes'

type Props = {
  teapod: string
  teabox?: string
}
export const ListTeas = ({ teapod, teabox }: Props) => {
  const info = useGetTeapodInfo(teapod)
  const teaboxes = info.data?.teaboxes.map((v) => v.name) ?? []
  const selectedTeabox = teabox ?? teaboxes[0]

  return (
    <>
      <Heading as='h3'>
        <Flex>
          <Box width='500px'>{teapod} <AddTea teapod={teapod} teabox={selectedTeabox} /></Box>
          <Box flexGrow='1' flexShrink='1'>
            {teaboxes.length > 0 && <ListTeasCtl teapod={teapod} teabox={selectedTeabox} teaboxes={teaboxes} />}
          </Box>
        </Flex>
      </Heading>
      <Text as='p' size='4' mt='2' mb='2' color='gray'></Text>
      <ListTeasTable teapod={teapod} teabox={selectedTeabox} />
    </>
  )
}
