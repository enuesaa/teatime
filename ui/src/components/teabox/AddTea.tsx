import { Textarea } from '../common/Textarea'
import styles from './AddTea.css'
import { useGetTeapodInfo } from '@/lib/api/teapods'
import { useAddTea } from '@/lib/api/teas'
import { useGetTeasFilter } from '@/states/teasfilter'
import { useAddToast } from '@/states/toast'
import { Dialog, Button, Callout } from '@radix-ui/themes'
import { FormEventHandler, useState } from 'react'
import { FaPlus } from 'react-icons/fa'

const useAddTeaForm = (teapod: string, teabox: string) => {
  const info = useGetTeapodInfo(teapod)
  const addTea = useAddTea(teapod, teabox)
  const addToast = useAddToast()

  const submit: FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault()
    const formdata = new FormData(e.currentTarget)
    const data = Object.fromEntries(formdata.entries())
    addTea.mutate(data)

    addToast({ title: 'OK', description: 'tea created!'})
  }

  const reset = () => {
    addTea.reset()
  }

  const error = addTea.data?.error
  const hasError = error !== undefined

  if (addTea.isSuccess && !hasError) {
    reset()
  }
  const inputs = info.data?.teaboxes.find((b) => b.name === teabox)?.inputs ?? []

  return { submit, reset, error, hasError, inputs }
}

export const AddTea = () => {
  const filter = useGetTeasFilter()
  const [open, setOpen] = useState(false)
  const form = useAddTeaForm(filter.teapod, filter.teabox)

  return (
    <Dialog.Root open={open || form.hasError} onOpenChange={setOpen}>
      <Dialog.Trigger>
        <Button variant='outline' radius='full' mx='2' style={{ padding: '10px' }}>
          <FaPlus />
        </Button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth='450px'>
        <Dialog.Title>Add Tea</Dialog.Title>

        {form.hasError && (
          <Callout.Root>
            <Callout.Text style={{whiteSpace: 'pre'}}>{form.error}</Callout.Text>
          </Callout.Root>
        )}

        <form onSubmit={form.submit}>
          {form.inputs.map(v => (
            <Textarea label={v.name} className={styles.texts} name={v.name} key={v.name} />
          ))}

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
