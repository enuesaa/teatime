import { useStyles } from '@/styles/use'

export const Dashboard = () => {
  const styles = useStyles(theme => ({
    h2: theme({ size: 'x3' }),
  }))

  return (
    <>
      <h2 css={styles.h2}>Setting</h2>
    </>
  )
}