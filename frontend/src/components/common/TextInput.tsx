import { UseFormRegisterReturn } from 'react-hook-form'

type Props = {
  label: string;
  regist: UseFormRegisterReturn<string>;
}
export const TextInput = ({ label, regist }: Props) => {
  // see https://zenn.dev/dqn/articles/59b4f12ad37b6b

  return (
    <label>
      {label}
      <input type='text' {...regist} />    
    </label>
  )
}
