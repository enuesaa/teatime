import styles from './AddTeapod.css'
import { useDeleteTeapod } from '@/lib/api/teapods'
import { AlertDialog, Button } from '@radix-ui/themes'
import { useForm } from 'react-hook-form'

type Props = {
  name: string
}
export const DeleteTeapod = ({ name }: Props) => {
  const delteTeapod = useDeleteTeapod(name)
  const { handleSubmit } = useForm<{}>()
  const submit = handleSubmit(() => delteTeapod.mutate({}))

  return (
    <AlertDialog.Root>
      <AlertDialog.Trigger>
        <Button variant='outline'>Delete</Button>
      </AlertDialog.Trigger>

      <AlertDialog.Content>
        <AlertDialog.Title>Delete App</AlertDialog.Title>
        <AlertDialog.Description>Are you sure?</AlertDialog.Description>

        <form onSubmit={submit} className={styles.actions}>
          <AlertDialog.Cancel>
            <Button variant='soft' color='gray'>
              Cancel
            </Button>
          </AlertDialog.Cancel>
          <AlertDialog.Action>
            <Button variant='solid' color='red' type='submit'>
              Delete
            </Button>
          </AlertDialog.Action>
        </form>
      </AlertDialog.Content>
    </AlertDialog.Root>
  )
}
