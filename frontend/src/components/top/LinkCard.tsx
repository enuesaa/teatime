import { useTheme, css } from '@emotion/react'
import Link from 'next/link'

type Props = {
  href: string;
  text: string;
}
export const LinkCard = ({ href, text }: Props) => {
  const theme = useTheme()
  const styles = {
    link: css(theme.linkCard, {}),
  }

  return (
    <Link href={href} css={styles.link}>{text}</Link>
  )
}
