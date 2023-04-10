import { ReactNode } from 'react'
import { useStyles } from '@/styles/use'

type Props = {
  children: ReactNode
}
export const Main = ({ children }: Props) => {
  const styles = useStyles(theme => ({
    main: theme({surf: 'main'}).css({
      margin: '20px',
      padding: '0 10px 10px 10px',
    }),
  }))

  return <section css={styles.main}>{children}</section>
}
