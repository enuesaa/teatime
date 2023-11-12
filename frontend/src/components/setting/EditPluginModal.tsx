import { css } from '@emotion/react'
import { Dialog, Button, Flex, Text, TextField } from '@radix-ui/themes'

export const EditPluginModal = () => {
  const styles = {
    trigger: css({
      cursor: 'pointer',
      height: '26px',
    })
  }

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <Button radius='full' color='gray' css={styles.trigger}>edit</Button>
      </Dialog.Trigger>

      <Dialog.Content>
        <Dialog.Title>Edit Plugin</Dialog.Title>
        <Dialog.Description mb='4'></Dialog.Description>

        <Flex direction='column' gap='3'>
          <label>
            <Text as='div' size='2' mb='1' weight='bold'>Name</Text>
            <TextField.Input defaultValue='name' />
          </label>
          <label>
            <Text as='div' size='2' mb='1' weight='bold'>Command</Text>
            <TextField.Input defaultValue='command' />
          </label>
        </Flex>

        <Flex gap='3' mt='4' justify='end'>
          <Dialog.Close>
            <Button variant='soft' color='gray'>Cancel</Button>
          </Dialog.Close>
          <Dialog.Close>
            <Button>Update</Button>
          </Dialog.Close>
        </Flex>

      </Dialog.Content>
    </Dialog.Root>
  )
}
