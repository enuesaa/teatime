import { Text, Strong } from '@radix-ui/themes'
import { BaseCard } from './BaseCard'
import { useGetProviderCard } from '@/lib/api';

type Props = {
  cardName: string;
}
export const MainCard = ({ cardName }: Props) => {
  const {data, isLoading} = useGetProviderCard('pinit', cardName)

  return (
    <BaseCard>
      <Text size='3' weight='bold' as='div'>
        <Strong>mainCard</Strong>
        <div>enable {!isLoading && data?.enable} </div>
        <div>layout {!isLoading && data?.layout} </div>
        <div>textCard.heading {!isLoading && data?.textCard.heading} </div>
        <div>textCard.center {!isLoading && data?.textCard.center} </div>
      </Text>
    </BaseCard>
  )
}
