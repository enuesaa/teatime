import { AlertDialog, Button } from '@radix-ui/themes'
import { useForm } from 'react-hook-form'
import styles from './DeleteAllLogs.css'
import { useDeleteAllLogs } from '@/lib/api/log'

export const DeleteAllLogs = () => {
  const deleteAllLogs = useDeleteAllLogs()
  const { handleSubmit } = useForm<{}>()
  const submit = handleSubmit(() => deleteAllLogs.mutate({}))

  return (
    <AlertDialog.Root>
      <AlertDialog.Trigger>
        <Button variant='outline'>Delete All</Button>
      </AlertDialog.Trigger>

      <AlertDialog.Content>
        <AlertDialog.Title>Delete App</AlertDialog.Title>
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
