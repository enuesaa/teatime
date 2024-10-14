import styles from './AddTea.css'
import { useAddTea } from '@/lib/api/teas'
import { Dialog, Button, Callout } from '@radix-ui/themes'
import { useForm } from 'react-hook-form'
import { FaPlus } from 'react-icons/fa'
import { Textarea } from '../common/Textarea'
import { useState } from 'react'

type Props = {
  teapod: string
  teabox: string
}
type Form = {
  data: string
}
export const AddTea = ({ teapod, teabox }: Props) => {
  const [open, setOpen] = useState(false)
  const addTea = useAddTea(teapod, teabox)
  const form = useForm<Form>()

  const submit = form.handleSubmit(req => {
    const data = JSON.parse(req.data)
    addTea.mutate(data)
  })
  const hasError = addTea.data?.error !== undefined
  const reset = () => {
    addTea.reset()
    form.reset()
  }

  return (
    <Dialog.Root open={open || hasError} onOpenChange={setOpen}>
      <Dialog.Trigger>
        <Button variant='outline' radius='full' mx='2' style={{ padding: '10px' }}>
          <FaPlus />
        </Button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth='450px'>
        <Dialog.Title>Add Tea</Dialog.Title>
        <Dialog.Description />

        {hasError && 
          <Callout.Root>
            <Callout.Text>{addTea.data?.error}</Callout.Text>
          </Callout.Root>
        }

        <form onSubmit={submit}>
          <Textarea label='data' {...form.register('data')} />

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
