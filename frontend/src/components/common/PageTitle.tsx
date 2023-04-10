import { useStyles } from '@/styles/use'

type Props = {
  title: string,
}
export const PageTitle = ({ title }: Props) => {
  const styles = useStyles(theme => ({
    h2: theme({}).css({
      // color: theme.color.main,
      padding: '0 0 0 10px',
    }),
  }))

  return (
    <h2 css={styles.h2}>{title}</h2>
  )
}