import { useGetProviderInfo, useUpdateProvider } from '@/lib/api'
import { ProviderSchema } from '@/lib/schema'
import { css } from '@emotion/react'
import { Dialog, Button, Flex, Text, TextField } from '@radix-ui/themes'
import { useForm } from 'react-hook-form'

type Props = {
  id: string
}
export const EditPluginModal = ({ id }: Props) => {
  const {data: info} = useGetProviderInfo(id)
  const { mutateAsync: updateProvider } =  useUpdateProvider(id)
  const { register, handleSubmit } = useForm<ProviderSchema>()

  const styles = {
    trigger: css({
      cursor: 'pointer',
      height: '28px',
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

        <form onSubmit={handleSubmit(data => updateProvider(data))}>
          <Flex direction='column' gap='3'>
            <label>
              <Text as='div' size='2' mb='1' weight='bold'>Name</Text>
              <TextField.Input defaultValue={info?.data?.name} data-1p-ignore {...register('name')} />
            </label>
            <label>
              <Text as='div' size='2' mb='1' weight='bold'>Command</Text>
              <TextField.Input defaultValue={info?.data?.command} {...register('command')} />
            </label>
          </Flex>

          <Flex gap='3' mt='4' justify='end'>
            <Dialog.Close>
              <Button variant='soft' color='gray'>Cancel</Button>
            </Dialog.Close>
            <Dialog.Close>
              <Button type='submit'>Update</Button>
            </Dialog.Close>
          </Flex>
        </form>

      </Dialog.Content>
    </Dialog.Root>
  )
}
