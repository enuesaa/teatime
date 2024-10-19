import styles from './AddTea.css'
import { useUpdateTea } from '@/lib/api/teas'
import { Dialog, Button, Callout } from '@radix-ui/themes'
import { useForm } from 'react-hook-form'
import { Textarea } from '../common/Textarea'
import { KeyboardEventHandler, useState } from 'react'
import { format } from '@/lib/utility/json'
import { useGetTea } from '@/lib/api/teas'

type Form = {
  data: string
}
const useAddTeaForm = (teapod: string, teabox: string, teaId: string) => {
  const tea = useGetTea(teapod, teabox, teaId)
  const updateTea = useUpdateTea(teapod, teabox, teaId)
  const form = useForm<Form>()

  const submit = form.handleSubmit(req => updateTea.mutate(JSON.parse(req.data)))
  const reset = () => {
    updateTea.reset()
    form.reset()
  }

  const error = updateTea.data?.error
  const hasError = error !== undefined

  if (tea.isSuccess) {
    form.setValue('data', format(JSON.stringify(tea.data?.data)))
  }
  if (updateTea.isSuccess && !hasError) {
    reset()
  }

  return { ...form, submit, reset, error, hasError }
}

type Props = {
  teapod: string
  teabox: string
  teaId: string
}
export const UpdateTea = ({ teapod, teabox, teaId }: Props) => {
  const [open, setOpen] = useState(false)
  const form = useAddTeaForm(teapod, teabox, teaId)

  const handleKeyUp: KeyboardEventHandler<HTMLTextAreaElement> = (e) => {
    e.currentTarget.value = format(e.currentTarget.value)
  }

  return (
    <Dialog.Root open={open || form.hasError} onOpenChange={setOpen}>
      <Dialog.Trigger>
        <Button variant='outline'>Update</Button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth='450px'>
        <Dialog.Title>Add Tea</Dialog.Title>

        {form.hasError && 
          <Callout.Root>
            <Callout.Text>{form.error}</Callout.Text>
          </Callout.Root>
        }

        <form onSubmit={form.submit}>
          <Textarea label='data' onKeyUp={handleKeyUp} className={styles.texts} {...form.register('data')} />

          <div className={styles.actions}>
            <Dialog.Close>
              <Button variant='soft' color='gray' onClick={form.reset}>
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
