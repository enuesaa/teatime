import { Text, Strong } from '@radix-ui/themes'
import { BaseCard } from './BaseCard'

type Props = {
  rid: string
}
export const MainCard = ({ rid }: Props) => {
  return (
    <BaseCard>
      <Text size='3' weight='bold' as='div'>
        <Strong>Rid</Strong> {rid}
      </Text>
    </BaseCard>
  )
}
