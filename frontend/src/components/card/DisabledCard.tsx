import { Strong, Text } from '@radix-ui/themes'
import { BaseCard } from './BaseCard'

export const DisabledCard = () => {
  return (
    <BaseCard>
      <Text size='3' weight='bold' as='div'>
        <Strong>disabled</Strong>
      </Text>
    </BaseCard>
  )
}