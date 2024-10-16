import { ListLogs } from '@/components/logs/ListLogs'
import { Heading, Text } from '@radix-ui/themes'

export default function Page() {
  return (
    <>
      <Heading as='h2' size='8'>
        Logs
      </Heading>
      <Text as='p' size='4' mt='2' mb='6' color='gray'>
        teatime app logs
      </Text>

      <ListLogs />
    </>
  )
}
