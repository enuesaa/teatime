import styles from './AddTea.css'
import { useAddTea } from '@/lib/api/teas'
import { Dialog, Button } from '@radix-ui/themes'
import { useForm } from 'react-hook-form'
import { FaPlus } from 'react-icons/fa'
import { Textarea } from '../common/Textarea'

type Props = {
  teapod: string
  teabox: string
}
export const AddTea = ({ teapod, teabox }: Props) => {
  const addTea = useAddTea(teapod, teabox)
  const { register, handleSubmit, formState } = useForm<{data: string}>()
  const submit = handleSubmit(req => {
    const data = JSON.parse(req.data)
    addTea.mutate(data)
  })

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <Button variant='outline' radius='full'>
          <FaPlus />
        </Button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth='450px'>
        <Dialog.Title>Add Tea</Dialog.Title>
        <Dialog.Description />

        <form onSubmit={submit}>
          <Textarea label='data' {...register('data')} />

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
