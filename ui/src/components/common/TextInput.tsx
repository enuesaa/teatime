import styles from './TextInput.css'
import { TextField } from '@radix-ui/themes'
import { forwardRef } from 'react'

type Props = {
  label: string
} & React.ComponentPropsWithoutRef<typeof TextField.Root>

export const TextInput = forwardRef<HTMLInputElement, Props>(({ label, ...props }, ref) => {
  return (
    <label className={styles.label}>
      {label}
      <TextField.Root data-1p-ignore ref={ref} {...props} />
    </label>
  )
})
