import { useAddTea } from '@/lib/api'
import { Dialog, Button, Flex, Text, TextField, IconButton } from '@radix-ui/themes'
import { BiPlus } from 'react-icons/bi'
import { useForm } from 'react-hook-form'
import { LinksTeaSchema } from '@/lib/schema'

export const AddTeaModal = () => {
  const { mutate: addTea } = useAddTea()
  const { register, handleSubmit, reset } = useForm<LinksTeaSchema>()

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <IconButton radius='full' style={{ margin: '0 10px', fontSize: '20px' }}>
          <BiPlus />
        </IconButton>
      </Dialog.Trigger>

      <Dialog.Content>
        <Dialog.Title>Add Tea</Dialog.Title>
        <Dialog.Description mb='4'></Dialog.Description>

        <form
          onSubmit={handleSubmit((data) => {
            addTea(data)
            reset()
          })}
        >
          <Flex direction='column' gap='3'>
            <label>
              <Text as='div' size='2' mb='1' weight='bold'>
                Name
              </Text>
              <TextField.Root {...register('title')} />
            </label>
            <label>
              <Text as='div' size='2' mb='1' weight='bold'>
                Links
              </Text>
              <TextField.Root {...register('link')} />
            </label>
          </Flex>

          <Flex gap='3' mt='4' justify='end'>
            <Dialog.Close>
              <Button variant='soft' color='gray'>
                Cancel
              </Button>
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
