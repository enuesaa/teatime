import { useGetProviderCard } from '@/lib/api';
import { MainCard } from './MainCard'
import { DisabledCard } from './DisabledCard'

type Props = {
  name: string;
}
export const Card = ({ name }: Props) => {
  const {data, isLoading} = useGetProviderCard('pinit', name)
  console.log(data)

  if (isLoading || data?.enable === false) {
    return (<DisabledCard />)
  }

  if (data?.layout === 'textCard') {
    return (<MainCard name={name} data={data} />)
  }

  return (<DisabledCard />)
}

