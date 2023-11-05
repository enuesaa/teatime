import { Text, Strong, Badge, Inset } from '@radix-ui/themes'
import { BaseCard } from './BaseCard'
import { CardSchema } from '@/lib/schema';
import { css } from '@emotion/react';

type Props = {
  name: string;
  data: CardSchema;
}
export const TextCard = ({ name, data }: Props) => {
  const styles = {
    text: css({
      textAlign: data.textCard.center ? 'center' : 'left',
      height: '70%',
      display: 'flex',
      justifyContent: 'center',
      alignItems: 'center',
    })
  }

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
