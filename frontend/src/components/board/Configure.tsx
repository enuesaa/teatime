import { useStyles } from '@/styles/use'
import Link from 'next/link'
import { FaPlus } from 'react-icons/fa'

export const Configure = () => {
  const styles = useStyles(theme => ({
    h2: theme({ size: 'x3' }).css({
      display: 'inline-block',
    }),
    addLink: theme().css({
      margin: '10px',
    }),
  }))

  return (
    <>
      <section>
        <h2 css={styles.h2}>Board</h2>
        <Link href='/board/configure/add' css={styles.addLink}>
          <FaPlus />
        </Link>
      </section>
    </>
  )
}