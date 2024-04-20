import { TeapodCard } from './TeapodCard'
import { useGetTeapodInfo } from '@/lib/api/teapods'

type Props = {
  teapod: string
}
export const ListCards = ({ teapod }: Props) => {
  const info = useGetTeapodInfo(teapod)

  return <>{info.data && info.data?.cards.map((v, i) => <TeapodCard teapod={teapod} name={v} key={i} />)}</>
}
