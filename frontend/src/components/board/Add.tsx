import { useStyles } from '@/styles/use'
import Link from 'next/link'
import { FaPlus } from 'react-icons/fa'

export const Add = () => {
  const styles = useStyles(theme => ({
    h2: theme({ size: 'x3' }).css({
      display: 'inline-block',
    }),
  }))

  return (
    <>
      <section>
        <h2 css={styles.h2}>Board Add</h2>
      </section>
    </>
  )
}