import { useGetTeapodInfo } from '@/lib/api/teapods'
import { TeapodCard } from './TeapodCard'

type Props = {
  teapod: string
}
export const ListCards = ({ teapod }: Props) => {
  const info = useGetTeapodInfo(teapod)

  return (
    <>
      {info.data && info.data?.cards.map((v, i) => (
        <TeapodCard teapod={teapod} name={v} key={i} />
      ))}
    </>
  )
}
