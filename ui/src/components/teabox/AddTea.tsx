import styles from './AddTea.css'
import { useAddTea } from '@/lib/api/teas'
import { Dialog, Button, Callout } from '@radix-ui/themes'
import { useForm } from 'react-hook-form'
import { FaPlus } from 'react-icons/fa'
import { Textarea } from '../common/Textarea'
import { KeyboardEventHandler, useState } from 'react'
import { useGetTeapodInfo } from '@/lib/api/teapods'

type Props = {
  teapod: string
  teabox: string
}
type Form = {
  data: string
}
export const AddTea = ({ teapod, teabox }: Props) => {
  const [open, setOpen] = useState(false)
  const teapodinfo = useGetTeapodInfo(teapod)
  const defaultValue = teapodinfo.data?.teaboxes.filter(b => b.name === teabox).at(0)?.placeholder ?? '{}'
  const addTea = useAddTea(teapod, teabox)
  const form = useForm<Form>()

  const hasError = addTea.data?.error !== undefined
  const submit = form.handleSubmit(req => {
    const data = JSON.parse(req.data)
    addTea.mutate(data)
    setOpen(false)
  })
  const reset = () => {
    addTea.reset()
    form.reset()
  }

  const format: KeyboardEventHandler<HTMLTextAreaElement> = (e) => {
    try {
      const data = JSON.parse(e.currentTarget.value)
      e.currentTarget.value = JSON.stringify(data, null, '  ')
    } catch (e) {}
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
          <Textarea label='data' onKeyUp={format} className={styles.texts} defaultValue={defaultValue}
            {...form.register('data')}
          />

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
