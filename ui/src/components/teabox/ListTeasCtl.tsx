import { SegmentedControl } from '@radix-ui/themes'
import { useNavigate } from 'react-router-dom'

type Props = {
  teapod: string
  teaboxes: string[]
}
export const ListTeasCtl = ({ teapod, teaboxes }: Props) => {
  const navigate = useNavigate()

  const gotoTeaboxPage = (teabox: string) => {
    navigate(`/teapods/${teapod}/teas?teabox=${teabox}`)
  }

  return (
    <SegmentedControl.Root
      defaultValue={teaboxes[0]}
      onValueChange={gotoTeaboxPage}
      radius='full'
      variant='classic'
      size='3'
    >
      {teaboxes.map((v, i) => (
        <SegmentedControl.Item value={v} key={i}>
          {v}
        </SegmentedControl.Item>
      ))}
    </SegmentedControl.Root>
  )
}
