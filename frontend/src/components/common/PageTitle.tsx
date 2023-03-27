import { css, useTheme } from "@emotion/react"

type Props = {
  title: string,
}
export const PageTitle = ({ title }: Props) => {
  const theme = useTheme()
  const styles = {
    h2: css(theme.heading, {
      color: theme.color.main,
      padding: '0 0 0 10px',
    }),
  }

  return (
    <h2 css={styles.h2}>{title}</h2>
  )
}