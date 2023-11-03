import { css } from '@emotion/react'
import { Card, Inset, Dialog, DialogClose, Button } from '@radix-ui/themes'
import { ReactNode } from 'react'

type Props = {
  children: ReactNode,
}
export const BaseCard = ({ children }: Props) => {
  const styles = {
    main: css({
      display: 'inline-block',
      width: '300px',
      height: '300px',
      margin: '20px',
      cursor: 'pointer',
    }),
    dialog: css({
      maxWidth: '1000px',
      position: 'relative',
    }),
    dialogClose: css({
      position: 'absolute',
      top: '10px',
      right: '10px',
      cursor: 'pointer',
    })
  }

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
