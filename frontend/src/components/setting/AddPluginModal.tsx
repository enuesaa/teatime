import { useAddProvider } from '@/lib/api'
import { css } from '@emotion/react'
import { Dialog, Button, Flex, Text, TextField, IconButton } from '@radix-ui/themes'
import { BiPlus } from 'react-icons/bi'
import { useForm, SubmitHandler } from 'react-hook-form'
import { ProviderSchema } from '@/lib/schema'

export const AddPluginModal = () => {
  const addProvider = useAddProvider()
  const { register, handleSubmit } = useForm<ProviderSchema>()

  const onSubmit: SubmitHandler<ProviderSchema> = async (data) => {
    await addProvider.mutateAsync(data)
  }

  const styles = {
    trigger: css({
      margin: '0 10px',
      cursor: 'pointer',
      fontSize: '20px',
    })
  }

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <IconButton color='purple' radius='full' css={styles.trigger}><BiPlus /></IconButton>
      </Dialog.Trigger>

      <Dialog.Content>
        <Dialog.Title>Add Plugin</Dialog.Title>
        <Dialog.Description mb='4'></Dialog.Description>

        <form onSubmit={handleSubmit(onSubmit)}>
          <Flex direction='column' gap='3'>
            <label>
              <Text as='div' size='2' mb='1' weight='bold'>Name</Text>
              <TextField.Input {...register('name')} />
            </label>
            <label>
              <Text as='div' size='2' mb='1' weight='bold'>Command</Text>
              <TextField.Input {...register('command')} />
            </label>
          </Flex>

          <Flex gap='3' mt='4' justify='end'>
            <Dialog.Close>
              <Button variant='soft' color='gray'>Cancel</Button>
            </Dialog.Close>
            <Dialog.Close>
              <Button type='submit'>Save</Button>
            </Dialog.Close>
          </Flex>
        </form>

      </Dialog.Content>
    </Dialog.Root>
  )
}
