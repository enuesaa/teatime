import { SegmentedControl } from '@radix-ui/themes'

type Props = {
  handleTeaboxChange: (value: string) => void
  teaboxes: string[]
}
export const ListTeasCtl = ({ handleTeaboxChange, teaboxes }: Props) => {
  return (
    <SegmentedControl.Root
      defaultValue={teaboxes[0]}
      onValueChange={handleTeaboxChange}
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
