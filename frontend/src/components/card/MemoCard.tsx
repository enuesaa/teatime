import { TextArea } from '@radix-ui/themes'
import { BaseCard } from './BaseCard'

export const MemoCard = () => {
  return (
    <BaseCard>
      <TextArea placeholder='Reply to comment…' />
    </BaseCard>
  )
}
