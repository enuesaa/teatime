import { ReactNode } from 'react'
import { css, useTheme } from '@emotion/react'

type Props = {
  children: ReactNode
}
export const Main = ({ children }: Props) => {
  const theme = useTheme()

  const styles = {
    main: css({
      margin: '20px',
      padding: '0 10px 10px 10px',
      color: theme.color.main,
    }),
  }

  return <section css={styles.main}>{children}</section>
}
