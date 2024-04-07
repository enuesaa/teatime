import { LinksTeaSchema, useAddTea } from '@/lib/api/teas'
import { Dialog, Button, Flex, Text, TextField, IconButton, Kbd } from '@radix-ui/themes'
import { BiPlus } from 'react-icons/bi'
import { TeapodInfoTeabox, useGetTeapodInfo } from '@/lib/api/teapods'
import { FormEventHandler } from 'react'

type Props = {
  teapod: string
  teabox: string
}
export const AddTeaModal = ({ teapod, teabox: teaboxName }: Props) => {
  const addTea = useAddTea(teapod)
  const info = useGetTeapodInfo(teapod)
  const teabox = info.data?.teaboxes.find(v => v.name === teaboxName) ?? {vals: {}, name: ''} as TeapodInfoTeabox

  const handleSubmit: FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault()
    const data: LinksTeaSchema = {
      teabox: teabox.name,
      vals: {},
    }
    Object.keys(teabox.vals).map(key => {
      if (e.target.hasOwnProperty(key)) {
        data.vals[key] = (e.target as any)[key].value ?? '' 
      } else {
        data.vals[key] = ''
      }
    })
    addTea.mutate(data)
  }

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <IconButton radius='full' style={{ margin: '0 10px', fontSize: '20px' }}>
          <BiPlus />
        </IconButton>
      </Dialog.Trigger>

      <Dialog.Content>
        <Dialog.Title>Add <Kbd>{teabox.name.toUpperCase()}</Kbd></Dialog.Title>
        <Dialog.Description mb='4'>{teabox.comment}</Dialog.Description>

        <form onSubmit={handleSubmit}>
          <Flex direction='column' gap='3'>
            {teabox?.vals && Object.keys(teabox.vals).map((v, i) => (
              <label key={i}>
                <Text as='div' size='2' mb='1' weight='bold'>
                  {v}
                </Text>
                <TextField.Root name={v} />
              </label>
            ))}
          </Flex>

          <Flex gap='3' mt='4' justify='end'>
            <Dialog.Close>
              <Button variant='soft' color='gray'>
                Cancel
              </Button>
            </Dialog.Close>
            <Dialog.Close>
              <Button type='submit'>Save</Button>
            </Dialog.Close>
          </Flex>
        </form>
      </Dialog.Content>
    </Dialog.Root>
  )
}
