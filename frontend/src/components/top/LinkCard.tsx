import { css } from '@emotion/react'
import { useStyles } from '@/styles/use';
import Link from 'next/link'

type Props = {
  href: string;
  text: string;
}
export const LinkCard = ({ href, text }: Props) => {
  const styles = useStyles(theme => ({
    link: theme({}).css({
      display: 'inline-block',
      padding: '10px 20px',
      color: '#cccccc',
      fontSize: '20px',
      background: '#333333',
      margin: '10px',
      borderRadius: '10px',
    })
  }))

  return (
    <Link href={href} css={styles.link}>{text}</Link>
  )
}
