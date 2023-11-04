import { Header } from '@/components/common/Header'
import { useGetProviderInfo } from '@/lib/api'
import { MainCard } from '@/components/card/MainCard'
import { TableCard } from '@/components/card/TableCard'
import { MemoCard } from '@/components/card/MemoCard'

export default function TopPage() {
  const { data, error, isLoading } = useGetProviderInfo('pinit')

  return (
    <>
      <Header />
      {!isLoading && data?.cards.map(v => (<MainCard cardName={v} key={v} />))}
      <TableCard />
      <MemoCard />
    </>
  )
}
