import { Dialog, Button, TextField } from '@radix-ui/themes'
import { FaPlus } from 'react-icons/fa'
import { useAddTeapod, type AddReqSchema } from '@/lib/api/teapods'
import { useForm } from 'react-hook-form'

export const AddTeapod = () => {
  const addTeapod = useAddTeapod()
  const { register, handleSubmit, formState } = useForm<AddReqSchema>()
  const submit = handleSubmit(data => addTeapod.mutate(data))

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <Button><FaPlus /></Button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth='450px'>
        <Dialog.Title>Add Teapod</Dialog.Title>
        <Dialog.Description />

        <form onSubmit={submit}>
          <TextField.Root data-1p-ignore {...register('name')} />

          <Dialog.Close>
            <Button variant='soft' color='gray'>Cancel</Button>
          </Dialog.Close>

          <Dialog.Close>
            <Button type='submit'>Save</Button>
          </Dialog.Close>
        </form>

      </Dialog.Content>
    </Dialog.Root>
  )
}
