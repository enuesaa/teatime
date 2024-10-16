import styles from './AddTeapod.css'
import { TextInput } from '@/components/common/TextInput'
import { useAddTeapod, type AddReqSchema } from '@/lib/api/teapods'
import { Dialog, Button, Callout } from '@radix-ui/themes'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { FaPlus } from 'react-icons/fa'

export const AddTeapod = () => {
  const [open, setOpen] = useState(false)
  const addTeapod = useAddTeapod()
  const form = useForm<AddReqSchema>()

  const hasError = addTeapod.data?.error !== undefined
  const submit = form.handleSubmit((data) => {
    addTeapod.mutate(data)
    setOpen(false)
  })
  const reset = () => {
    addTeapod.reset()
    form.reset()
  }

  return (
    <Dialog.Root open={open || hasError} onOpenChange={setOpen}>
      <Dialog.Trigger>
        <Button variant='outline' radius='full' mx='2' mt='1' style={{ padding: '10px' }}>
          <FaPlus />
        </Button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth='450px'>
        <Dialog.Title>Add Teapod</Dialog.Title>
        <Dialog.Description />

        {hasError && 
          <Callout.Root>
            <Callout.Text>{addTeapod.data?.error}</Callout.Text>
          </Callout.Root>
        }

        <form onSubmit={submit}>
          <TextInput label='name' {...form.register('name')} />

          <div className={styles.actions}>
            <Dialog.Close>
              <Button variant='soft' color='gray' onClick={reset}>
                Cancel
              </Button>
            </Dialog.Close>

            <Button type='submit'>Save</Button>
          </div>
        </form>
      </Dialog.Content>
    </Dialog.Root>
  )
}
