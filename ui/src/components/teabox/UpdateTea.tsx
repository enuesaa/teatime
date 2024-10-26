import { TextInput } from '../common/TextInput'
import { Textarea } from '../common/Textarea'
import styles from './AddTea.css'
import { useUpdateTea } from '@/lib/api/teas'
import { useGetTea } from '@/lib/api/teas'
import { format } from '@/lib/utility/json'
import { tryunwrap } from '@/lib/utility/try'
import { useGetTeasFilter } from '@/states/teasfilter'
import { Dialog, Button, Callout } from '@radix-ui/themes'
import { KeyboardEventHandler, useEffect, useState } from 'react'
import { useForm } from 'react-hook-form'

type Form = {
  data: string
}
const useUpdateTeaForm = (teapod: string, teabox: string, teaId: string) => {
  const tea = useGetTea(teapod, teabox, teaId)
  const updateTea = useUpdateTea(teapod, teabox, teaId)
  const form = useForm<Form>()

  const submit = form.handleSubmit((req) => updateTea.mutate(JSON.parse(req.data)))
  const reset = () => {
    updateTea.reset()
    form.reset()
  }

  const error = updateTea.data?.error
  const hasError = error !== undefined

  useEffect(() => {
    form.setValue('data', format(JSON.stringify(tea.data?.data)))
  }, [tea.dataUpdatedAt])
  const prepared = tryunwrap(() => JSON.parse(form.watch('data')), true, false)

  return { ...form, submit, reset, error, hasError, current: tea.data, prepared }
}

type Props = {
  teaId: string
}
export const UpdateTea = ({ teaId }: Props) => {
  const filter = useGetTeasFilter()
  const [open, setOpen] = useState(false)
  const form = useUpdateTeaForm(filter.teapod, filter.teabox, teaId)

  const handleKeyUp: KeyboardEventHandler<HTMLTextAreaElement> = (e) => {
    e.currentTarget.value = format(e.currentTarget.value)
  }

  return (
    <Dialog.Root open={open || form.hasError} onOpenChange={setOpen}>
      <Dialog.Trigger>
        <Button variant='outline'>Update</Button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth='450px'>
        <Dialog.Title>Update Tea</Dialog.Title>

        {form.hasError && (
          <Callout.Root>
            <Callout.Text>{form.error}</Callout.Text>
          </Callout.Root>
        )}

        <form onSubmit={form.submit}>
          <TextInput label='id' value={form.current?.id} readOnly />
          <Textarea label='data' onKeyUp={handleKeyUp} className={styles.texts} {...form.register('data')} />
          <TextInput label='created' value={form.current?.created} readOnly />
          <TextInput label='updated' value={form.current?.updated} readOnly />

          {!form.prepared && (
            <Callout.Root>
              <Callout.Text>JSON Format Error</Callout.Text>
            </Callout.Root>
          )}

          <div className={styles.actions}>
            <Dialog.Close>
              <Button variant='soft' color='gray' onClick={form.reset}>
                Cancel
              </Button>
            </Dialog.Close>
            <Dialog.Close>
              <Button type='submit' disabled={!form.prepared}>
                Save
              </Button>
            </Dialog.Close>
          </div>
        </form>
      </Dialog.Content>
    </Dialog.Root>
  )
}
