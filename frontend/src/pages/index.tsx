import { Header } from '@/components/common/Header'
import { useGetProviderInfo } from '@/lib/api'
import { Card } from '@/components/card/Card'
import { TableCard } from '@/components/card/TableCard'
import { MemoCard } from '@/components/card/MemoCard'

export default function TopPage() {
  // const { data, error, isLoading } = useGetProviderInfo('pinit')

  return (
    <>
      <Header />
      {/* {!isLoading && data?.cards.map(v => (<Card name={v} key={v} />))} */}
      <TableCard />
      <MemoCard />
    </>
  )
}
