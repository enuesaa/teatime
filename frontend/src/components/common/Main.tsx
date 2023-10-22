import { ReactNode } from 'react'

type Props = {
  children: ReactNode
}
export const Main = ({ children }: Props) => {
  // const styles = useStyles((theme) => ({
  //   main: theme({ surf: 'main' }).css({
  //     margin: '0 15px 15px 15px',
  //     padding: '0 10px 10px 10px',
  //   }),
  // }))

  return <section>{children}</section>
}