import { ReactNode } from 'react'
import { css } from '@emotion/react'

type Props = {
  children: ReactNode
}
export function Main({ children }: Props) {
  const styles = {
    main: css({
      width: '90%',
      margin: '0 auto',
      height: '100vh',
      padding: '20px 0',
    }),
  }

  return <section css={styles.main}>{children}</section>
}