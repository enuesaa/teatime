import { useGetTeasFilter } from '@/states/teasfilter'
import { SegmentedControl } from '@radix-ui/themes'
import { useNavigate } from 'react-router-dom'

export const ListTeasCtl = () => {
  const filter = useGetTeasFilter()
  const navigate = useNavigate()

  const gotoTeaboxPage = (teabox: string) => {
    navigate(`/teapods/${filter.teapod}/teas?teabox=${teabox}`)
  }

  return (
    <SegmentedControl.Root
      defaultValue={filter.teabox}
      onValueChange={gotoTeaboxPage}
      radius='full'
      variant='classic'
      size='3'
    >
      {filter.teaboxes.map((v, i) => (
        <SegmentedControl.Item value={v} key={i}>
          {v}
        </SegmentedControl.Item>
      ))}
    </SegmentedControl.Root>
  )
}
