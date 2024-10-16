import { AlertDialog, Button } from '@radix-ui/themes'
import { useForm } from 'react-hook-form'
import styles from './AddTea.css'
import { useDeleteTea } from '@/lib/api/teas'

type Props = {
  teapod: string
  teabox: string
  teaId: string
}
export const DeleteTea = ({ teapod, teabox, teaId }: Props) => {
  const delteTeapod = useDeleteTea(teapod, teabox, teaId)
  const { handleSubmit } = useForm<{}>()
  const submit = handleSubmit(() => delteTeapod.mutate({}))

  return (
    <AlertDialog.Root>
      <AlertDialog.Trigger>
        <Button variant='outline'>Delete</Button>
      </AlertDialog.Trigger>

      <AlertDialog.Content>
        <AlertDialog.Title>Delete Tea</AlertDialog.Title>
        <AlertDialog.Description>Are you sure?</AlertDialog.Description>

        <form onSubmit={submit} className={styles.actions}>
          <AlertDialog.Cancel>
              <Button variant='soft' color='gray'>Cancel</Button>
            </AlertDialog.Cancel>
            <AlertDialog.Action>
              <Button variant='solid' color='red' type='submit'>Delete</Button>
            </AlertDialog.Action>
        </form>
      </AlertDialog.Content>
    </AlertDialog.Root>
  )
}