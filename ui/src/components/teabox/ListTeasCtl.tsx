import { SegmentedControl } from '@radix-ui/themes'
import { useNavigate } from 'react-router-dom'

type Props = {
  teapod: string
  teabox: string
  teaboxes: string[]
}
export const ListTeasCtl = ({ teapod, teabox, teaboxes }: Props) => {
  const navigate = useNavigate()

  const gotoTeaboxPage = (teabox: string) => {
    navigate(`/teapods/${teapod}/teas?teabox=${teabox}`)
  }

  return (
    <SegmentedControl.Root
      defaultValue={teabox}
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
