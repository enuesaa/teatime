import { Button, DropdownMenu } from '@radix-ui/themes'
import styles from './ActionMenu.css'
import { useGetTeapodInfo } from '@/lib/api/teapods'

type Props = {
  teapod: string
}
export const ActionMenu = ({ teapod }: Props) => {
  const info = useGetTeapodInfo(teapod)

  return (
    <DropdownMenu.Root>
      <DropdownMenu.Trigger>
        <Button variant='outline'>
          Action
          <DropdownMenu.TriggerIcon />
        </Button>
      </DropdownMenu.Trigger>

      <DropdownMenu.Content className={styles.content}>
        {info.data?.actions.map((v, i) => (
          <DropdownMenu.Item key={i}>{v.name}</DropdownMenu.Item>
        ))}
      </DropdownMenu.Content>

    </DropdownMenu.Root>
  )
}
