import { Button, DropdownMenu } from '@radix-ui/themes'
import styles from './ActionMenu.css'
import { useGetTeapodInfo, useActTeapod } from '@/lib/api/teapods'
import { MouseEventHandler } from 'react'
import { useAddToast } from '@/states/toast'

type Props = {
  teapod: string
}
export const ActionMenu = ({ teapod }: Props) => {
  const info = useGetTeapodInfo(teapod)

  return (
    <DropdownMenu.Root>
      <DropdownMenu.Trigger disabled={info.data?.actions.length === 0}>
        <Button variant='outline'>
          Action
          <DropdownMenu.TriggerIcon />
        </Button>
      </DropdownMenu.Trigger>

      <DropdownMenu.Content className={styles.content}>
        {info.data?.actions.map((v, i) => (
          <ActionMenuItem key={i} name={v.name} teapod={teapod} />
        ))}
      </DropdownMenu.Content>

    </DropdownMenu.Root>
  )
}

const ActionMenuItem = ({ name, teapod }: { name: string, teapod: string }) => {
  const actTeapod = useActTeapod(teapod)
  const addToast = useAddToast()

  const handleClick: MouseEventHandler<HTMLDivElement> = async (e) => {
    e.preventDefault()
    const res = await actTeapod.mutateAsync({
      action: name,
    })
    addToast({ title: 'OK', description: res.message })
  }

  return (
    <DropdownMenu.Item onClick={handleClick}>
      {name}
    </DropdownMenu.Item>
  )
}
