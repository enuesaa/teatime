import styles from './AddTeapod.css'
import { TextInput } from '@/components/common/TextInput'
import { useAddTeapod, type AddReqSchema } from '@/lib/api/teapods'
import { Dialog, Button, Callout } from '@radix-ui/themes'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { FaPlus } from 'react-icons/fa'

const useAddTeapodForm = () => {
  const addTeapod = useAddTeapod()
  const form = useForm<AddReqSchema>()

  const submit = form.handleSubmit((data) => addTeapod.mutate(data))
  const reset = () => {
    addTeapod.reset()
    form.reset()
  }
  const error = addTeapod.data?.error
  const hasError = error !== undefined

  if (addTeapod.isSuccess && !hasError) {
    reset()
  }

  return { ...form, submit, reset, error, hasError }
}

export const AddTeapod = () => {
  const [open, setOpen] = useState(false)
  const form = useAddTeapodForm()

  return (
    <Dialog.Root open={open || form.hasError} onOpenChange={setOpen}>
      <Dialog.Trigger>
        <Button variant='outline' radius='full' mx='2' mt='1' style={{ padding: '10px' }}>
          <FaPlus />
        </Button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth='450px'>
        <Dialog.Title>Add Teapod</Dialog.Title>

        {form.hasError && (
          <Callout.Root>
            <Callout.Text>{form.error}</Callout.Text>
          </Callout.Root>
        )}

        <form onSubmit={form.submit}>
          <TextInput label='name' {...form.register('name')} />

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
