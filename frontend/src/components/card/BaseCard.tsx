import { css } from '@emotion/react'
import { Card, Inset } from '@radix-ui/themes'
import { ReactNode } from 'react'

type Props = {
  children: ReactNode,
}
export const BaseCard = ({ children }: Props) => {
  const styles = {
    main: css({
      display: 'inline-block',
      width: '300px',
      height: '300px',
      margin: '20px',
    }),
  }

  return (
    <Card size='4' css={styles.main}>
      <Inset clip='padding-box'>
        {children}
      </Inset>
    </Card>
  )
}