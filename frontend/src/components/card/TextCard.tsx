import { Text, Strong, Badge, Inset } from '@radix-ui/themes'
import { BaseCard } from './BaseCard'
import { CardSchema } from '@/lib/schema';
import styles from './TextCard.css'

type Props = {
  name: string;
  data: CardSchema;
}
export const TextCard = ({ name, data }: Props) => {
  return (
    <BaseCard>
      <Inset clip='padding-box' side='top' pb='current'>
        <Badge color='gray' size='2'>{data.textCard.heading}</Badge>
      </Inset>
      <Text size='6' as='div' m='-1' css={styles.text}>
        {data.textCard.text}
      </Text>
    </BaseCard>
  )
}
