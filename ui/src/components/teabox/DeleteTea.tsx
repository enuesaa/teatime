import styles from './AddTea.css'
import { useDeleteTea } from '@/lib/api/teas'
import { useGetTeasFilter } from '@/states/teasfilter'
import { AlertDialog, Button } from '@radix-ui/themes'
import { useForm } from 'react-hook-form'

type Props = {
  teaId: string
}
export const DeleteTea = ({ teaId }: Props) => {
  const filter = useGetTeasFilter()
  const delteTeapod = useDeleteTea(filter.teapod, filter.teabox, teaId)
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
