import { Card, Inset, Dialog, DialogClose, Button } from '@radix-ui/themes'
import { ReactNode } from 'react'
import styles from './BaseCard.css'

type Props = {
  children: ReactNode,
}
export const BaseCard = ({ children }: Props) => {
  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <Card size='4' css={styles.main}>
          {children}
        </Card>
      </Dialog.Trigger>
      <Dialog.Content css={styles.dialog}>
        <Dialog.Title>aa</Dialog.Title>
        a
        <DialogClose css={styles.dialogClose}>
          <Button variant='soft' color='gray'>Close</Button>
        </DialogClose>
      </Dialog.Content>
    </Dialog.Root>
  )
}
