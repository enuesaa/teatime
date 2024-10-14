import { Dialog, Button } from '@radix-ui/themes'
import { FaPlus } from 'react-icons/fa'
import { useAddTeapod, type AddReqSchema } from '@/lib/api/teapods'
import { useForm } from 'react-hook-form'
import { TextInput } from '@/components/common/TextInput'
import styles from './AddTeapod.css'

export const AddTeapod = () => {
  const addTeapod = useAddTeapod()
  const { register, handleSubmit, formState } = useForm<AddReqSchema>()
  const submit = handleSubmit(data => addTeapod.mutate(data))

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <Button className={styles.trigger}><FaPlus /></Button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth='450px'>
        <Dialog.Title>Add Teapod</Dialog.Title>
        <Dialog.Description />

        <form onSubmit={submit}>
          <TextInput label='name' {...register('name')} />

          <div className={styles.actions}>
            <Dialog.Close>
              <Button variant='soft' color='gray'>Cancel</Button>
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
