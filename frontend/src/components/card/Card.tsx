import { useGetProviderCard } from '@/lib/api'
import { TextCard } from './TextCard'
import { DisabledCard } from './DisabledCard'

type Props = {
  name: string;
}
export const Card = ({ name }: Props) => {
  const {data, isLoading} = useGetProviderCard('pinit', name)

  if (isLoading || data?.enable === false) {
    return (<DisabledCard />)
  }

  if (data?.layout === 'textCard') {
    return (<TextCard name={name} data={data} />)
  }

  return (<DisabledCard />)
}
