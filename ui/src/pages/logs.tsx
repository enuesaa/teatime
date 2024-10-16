import { DeleteAllLogs } from '@/components/logs/DeleteAllLogs'
import { ListLogs } from '@/components/logs/ListLogs'
import { Heading, Text } from '@radix-ui/themes'

export default function Page() {
  return (
    <>
      <Heading as='h2' size='8' style={{ position: 'relative' }}>
        Logs
        <div style={{ position: 'absolute', top: '0', right: '0' }}>
          <DeleteAllLogs />
        </div>
      </Heading>
      <Text as='p' size='4' mt='2' mb='6' color='gray'>
        teatime app logs
      </Text>

      <ListLogs />
    </>
  )
}
