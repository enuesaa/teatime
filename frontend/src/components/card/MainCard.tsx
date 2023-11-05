import { Text, Strong } from '@radix-ui/themes'
import { BaseCard } from './BaseCard'
import { useGetProviderCard } from '@/lib/api';
import { CardSchema } from '@/lib/schema';

type Props = {
  name: string;
  data: CardSchema;
}
export const MainCard = ({ name, data }: Props) => {
  return (
    <BaseCard>
      <Text size='3' weight='bold' as='div'>
        <Strong>{data.textCard.heading}</Strong>
        {/* <div>layout {!isLoading && data?.layout} </div>
        <div>textCard.heading {!isLoading && data?.textCard.heading} </div>
        <div>textCard.center {!isLoading && data?.textCard.center} </div> */}
      </Text>
    </BaseCard>
  )
}
