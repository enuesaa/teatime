import { css } from '@emotion/react'
import { Button, Flex, AlertDialog } from '@radix-ui/themes'

export const DeletePluginModal = () => {
  const styles = {
    trigger: css({
      cursor: 'pointer',
      height: '26px',
    })
  }

  return (
    <AlertDialog.Root>
      <AlertDialog.Trigger>
        <Button color='red' radius='full' css={styles.trigger}>delete</Button>
      </AlertDialog.Trigger>
      <AlertDialog.Content>
        <AlertDialog.Title>Delete Plugin</AlertDialog.Title>
        <AlertDialog.Description size='2'>
          Do you proceed?
        </AlertDialog.Description>

        <Flex gap='3' mt='4' justify='end'>
          <AlertDialog.Cancel>
            <Button variant='soft' color='gray'>Cancel</Button>
          </AlertDialog.Cancel>
          <AlertDialog.Action>
            <Button variant='solid' color='red'>Delete</Button>
          </AlertDialog.Action>
        </Flex>
      </AlertDialog.Content>
    </AlertDialog.Root>
  )
}
