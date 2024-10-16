import styles from './TextInput.css'
import { TextArea as RadixTextArea } from '@radix-ui/themes'
import { forwardRef } from 'react'

type Props = {
  label: string
} & React.ComponentPropsWithoutRef<typeof RadixTextArea>

export const Textarea = forwardRef<HTMLTextAreaElement, Props>(({ label, ...props }, ref) => {
  return (
    <label className={styles.label}>
      {label}
      <RadixTextArea data-1p-ignore ref={ref} {...props} />
    </label>
  )
})
