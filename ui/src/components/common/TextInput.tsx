import { TextField } from '@radix-ui/themes'
import styles from './TextInput.css'

type Props = {
  label: string
} & React.RefAttributes<HTMLInputElement>
export const TextInput = ({ label, ...props }: Props) => {
  return (
    <label className={styles.label}>
      {label}
      <TextField.Root data-1p-ignore {...props} />
    </label>
  )
}
