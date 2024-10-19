import { useGetTeasFilter } from '@/states/teasfilter'
import { SegmentedControl } from '@radix-ui/themes'
import { useNavigate } from 'react-router-dom'

export const ListTeasCtl = () => {
  const { teapod, teabox, teaboxes } = useGetTeasFilter()
  const navigate = useNavigate()

  const gotoTeaboxPage = (teabox: string) => {
    navigate(`/teapods/${teapod}/teas?teabox=${teabox}`)
  }

  if (teabox === undefined) {
    return (<>a</>)
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
