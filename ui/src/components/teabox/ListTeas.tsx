import { Heading, Text, Flex, Box } from '@radix-ui/themes'
import { AddTeaModal } from './AddTeaModal'
import { useState } from 'react'
import { ListTeasTable } from './ListTeasTable'
import { ListTeasCtl } from './ListTeasCtl'
import { useGetTeapodInfo } from '@/lib/api/teapods'

type Props = {
  teapod: string
}
export const ListTeas = ({ teapod }: Props) => {
  const info = useGetTeapodInfo(teapod)
  const [teabox, setTeabox] = useState<string>('links')
  const handleTeaboxChange = (value: string) => setTeabox(value)
  const teaboxes = info.data?.teaboxes.map(v => v.name) ?? []

  return (
    <>
      <Heading as='h3'>
        <Flex>
          <Box width='500px'>
            Teas <AddTeaModal teapod={teapod} teabox={teabox} />
          </Box>
          <Box flexGrow='1' flexShrink='1'>
            {teaboxes.length > 0 && <ListTeasCtl handleTeaboxChange={handleTeaboxChange} teaboxes={teaboxes} />}
          </Box>
        </Flex>
      </Heading>
      <Text as='p' size='4' mt='2' mb='2' color='gray'></Text>
      <ListTeasTable teapod={teapod} teabox={teabox} />
    </>
  )
}
