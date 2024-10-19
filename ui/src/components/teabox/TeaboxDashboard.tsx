import { useGetTeapodInfo } from '@/lib/api/teapods'
import { ListTeas } from './ListTeas'

type Props = {
  teapod: string
  teabox: null | string
}
export const TeaboxDashboard = ({ teapod, teabox }: Props) => {
  const info = useGetTeapodInfo(teapod)
  const teaboxes = info.data?.teaboxes.map((v) => v.name) ?? []
  const selectedTeabox = teabox ?? teaboxes[0]

  return <ListTeas teapod={teapod} teabox={selectedTeabox} teaboxes={teaboxes} />
}
