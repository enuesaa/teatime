import { Heading, Text } from '@radix-ui/themes'
import { ListTeapods } from '@/components/setting/ListTeapods'

export default function Page() {
  return (
    <>
      <Heading as='h2' size='8'>
        Settings
      </Heading>
      <Text as='p' size='4' mt='2' mb='6' color='gray'>
        teatime settings
      </Text>

      <ListTeapods />
    </>
  )
}
