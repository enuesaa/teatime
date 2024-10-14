import styles from './AddTeapod.css'
import { TextInput } from '@/components/common/TextInput'
import { useAddTeapod, type AddReqSchema } from '@/lib/api/teapods'
import { Dialog, Button } from '@radix-ui/themes'
import { useForm } from 'react-hook-form'
import { FaPlus } from 'react-icons/fa'

export const AddTeapod = () => {
  const addTeapod = useAddTeapod()
  const { register, handleSubmit, formState } = useForm<AddReqSchema>()
  const submit = handleSubmit((data) => addTeapod.mutate(data))

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <Button variant='outline'>
          <FaPlus />
        </Button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth='450px'>
        <Dialog.Title>Add Teapod</Dialog.Title>
        <Dialog.Description />

        <form onSubmit={submit}>
          <TextInput label='name' {...register('name')} />

          <div className={styles.actions}>
            <Dialog.Close>
              <Button variant='soft' color='gray'>
                Cancel
              </Button>
            </Dialog.Close>

            <Dialog.Close>
              <Button type='submit'>Save</Button>
            </Dialog.Close>
          </div>
        </form>
      </Dialog.Content>
    </Dialog.Root>
  )
}
