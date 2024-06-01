// import { AddTeaModal } from './AddTeaModal'
import { ListTeasCtl } from './ListTeasCtl'
import { ListTeasTable } from './ListTeasTable'
import { useGetTeapodInfo } from '@/lib/api/teapods'
import { Heading, Text, Flex, Box } from '@radix-ui/themes'

type Props = {
  teapod: string
  teabox: null | string
}
export const ListTeas = ({ teapod, teabox }: Props) => {
  const info = useGetTeapodInfo(teapod)
  const teaboxes = info.data?.teaboxes.map((v) => v.name) ?? []
  if (teabox === null) {
    teabox = teaboxes[0]
  }

  return (
    <>
      <Heading as='h3'>
        <Flex>
          <Box width='500px'>
            Teas {/* <AddTeaModal teapod={teapod} teabox={teabox} /> */}
          </Box>
          <Box flexGrow='1' flexShrink='1'>
            {teaboxes.length > 0 && <ListTeasCtl teapod={teapod} teaboxes={teaboxes} />}
          </Box>
        </Flex>
      </Heading>
      <Text as='p' size='4' mt='2' mb='2' color='gray'></Text>
      <ListTeasTable teapod={teapod} teabox={teabox} />
    </>
  )
}
