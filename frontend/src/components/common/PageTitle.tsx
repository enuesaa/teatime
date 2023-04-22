import { useStyles } from '@/styles/use'
import { ReactNode } from 'react'

type Props = {
  title: string
  children?: ReactNode
}
export const PageTitle = ({ title, children }: Props) => {
  const styles = useStyles((theme) => ({
    main: theme({ around: 'x2tb' }),
    h2: theme({ size: 'x3' }).css({
      display: 'inline-block',
      margin: '0 10px 0 0',
    }),
  }))

  return (
    <div css={styles.main}>
      <h2 css={styles.h2}>{title}</h2>
      {children}
    </div>
  )
}