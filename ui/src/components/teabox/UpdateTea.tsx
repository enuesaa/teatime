import { useGetTeapodInfo } from '@/lib/api/teapods'
import { TextInput } from '../common/TextInput'
import { Textarea } from '../common/Textarea'
import styles from './AddTea.css'
import { useUpdateTea } from '@/lib/api/teas'
import { useGetTea } from '@/lib/api/teas'
import { useGetTeasFilter } from '@/states/teasfilter'
import { Dialog, Button, Callout } from '@radix-ui/themes'
import { FormEventHandler, useState } from 'react'

const useUpdateTeaForm = (teapod: string, teabox: string, teaId: string) => {
  const info = useGetTeapodInfo(teapod)
  const tea = useGetTea(teapod, teabox, teaId)
  const updateTea = useUpdateTea(teapod, teabox, teaId)

  const submit: FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault()
    const formdata = new FormData(e.currentTarget)
    const data = Object.fromEntries(formdata.entries())
    updateTea.mutate(data)
  }

  const reset = () => {
    updateTea.reset()
  }

  const inputs = info.data?.teaboxes.find((b) => b.name === teabox)?.inputs ?? []
  const error = updateTea.data?.error
  const hasError = error !== undefined

  return { submit, reset, error, hasError, inputs, current: tea.data }
}

type Props = {
  teaId: string
}
export const UpdateTea = ({ teaId }: Props) => {
  const filter = useGetTeasFilter()
  const [open, setOpen] = useState(false)
  const form = useUpdateTeaForm(filter.teapod, filter.teabox, teaId)

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
          {form.inputs.map(v => (
            <Textarea
              key={v.name}
              label={v.name}
              className={styles.texts}
              name={v.name}
              defaultValue={form.current?.data[v.name] ?? ''}
            />
          ))}
          <TextInput label='created' value={form.current?.created} readOnly />
          <TextInput label='updated' value={form.current?.updated} readOnly />

          <div className={styles.actions}>
            <Dialog.Close>
              <Button variant='soft' color='gray' onClick={form.reset}>
                Cancel
              </Button>
            </Dialog.Close>
            <Dialog.Close>
              <Button type='submit'>
                Save
              </Button>
            </Dialog.Close>
          </div>
        </form>
      </Dialog.Content>
    </Dialog.Root>
  )
}
