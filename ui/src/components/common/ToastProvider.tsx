import { useHideToast, useListToast } from '@/states/toast'
import * as Toast from '@radix-ui/react-toast'
import styles from './ToastProvider.css'
import { MouseEventHandler } from 'react'
import { Button } from '@radix-ui/themes'

export const ToastProvider = () => {
  const toasts = useListToast()

  return (
    <Toast.Provider>
      {toasts.map((t, i) => (
        <Item id={t.id} title={t.title} description={t.description} key={i} />
      ))}
      <Toast.Viewport className={styles.viewport} />
    </Toast.Provider>
  )
}

const Item = ({ id, title, description }: { id: string, title: string, description: string }) => {
  const hide = useHideToast()

  const handleClick: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    hide({ id })
  }

  return (
    <Toast.Root asChild>
      <Button className={styles.toast} variant='ghost' onClick={handleClick}>
        <Toast.Title>{title}</Toast.Title>
        <Toast.Description>{description}</Toast.Description>
      </Button>
    </Toast.Root>
  )
}