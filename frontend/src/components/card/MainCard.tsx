import { css } from '@emotion/react'
import { Card, Inset, Text, Strong } from '@radix-ui/themes'

type Props = {
  rid: string
}
export const MainCard = ({ rid }: Props) => {
  const styles = {
    main: css({
      maxWidth: '300px',
      margin: '20px',
    }),
    img: css({
      display: 'block',
      objectFit: 'cover',
      width: '100%',
      height: 140,
    }),
  }

  return (
    <Card size='4' css={styles.main}>
      <Inset clip="padding-box" side="top" pb="current">
        <img css={styles.img} src="https://images.unsplash.com/photo-1617050318658-a9a3175e34cb?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=600&q=80" />
      </Inset>
      <Text size='3'>
        <Strong>Rid</Strong> {rid}
      </Text>
    </Card>
  )
}
