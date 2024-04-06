import { useDeleteTea } from '@/lib/api/teas'
import { Button, Flex, AlertDialog } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'

type Props = {
  teaid: string
}
export const DeleteTeaModal = ({ teaid }: Props) => {
  const { mutateAsync: deleteProvider } = useDeleteTea(teaid)
  const handleDelete: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    await deleteProvider()
  }

  return (
    <AlertDialog.Root>
      <AlertDialog.Trigger>
        <Button color='red' radius='full' style={{ height: '28px' }}>
          delete
        </Button>
      </AlertDialog.Trigger>
      <AlertDialog.Content>
        <AlertDialog.Title>Delete Tea</AlertDialog.Title>
        <AlertDialog.Description size='2'>Do you proceed?</AlertDialog.Description>

        <Flex gap='3' mt='4' justify='end'>
          <AlertDialog.Cancel>
            <Button variant='soft' color='gray'>
              Cancel
            </Button>
          </AlertDialog.Cancel>
          <AlertDialog.Action>
            <Button variant='solid' color='red' onClick={handleDelete}>
              Delete
            </Button>
          </AlertDialog.Action>
        </Flex>
      </AlertDialog.Content>
    </AlertDialog.Root>
  )
}
