import { useGetTeapodInfo } from '@/lib/api/teapods'

type Props = {
  name: string
}
export const TeapodDescription = ({ name }: Props) => {
  const { data: info } = useGetTeapodInfo(name)

  return <>{info?.data.description}</>
}
