import { AddTeapod } from '@/components/setting/AddTeapod'
import { ListTeapods } from '@/components/setting/ListTeapods'
import { Heading, Text } from '@radix-ui/themes'

export default function Page() {
  return (
    <>
      <Heading as='h2' size='8'>
        Settings <AddTeapod />
      </Heading>
      <Text as='p' size='4' mt='2' mb='6' color='gray'>
        teatime settings
      </Text>

      <ListTeapods />
    </>
  )
}
