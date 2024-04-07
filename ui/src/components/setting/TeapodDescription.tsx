import { useGetTeapodInfo } from '@/lib/api/teapods'

type Props = {
  name: string
}
export const TeapodDescription = ({ name }: Props) => {
  const info = useGetTeapodInfo(name)

  return <>{info.data?.description}</>
}
