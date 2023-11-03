import { Header } from '@/components/common/Header'
import { useFetchCards } from '@/lib/api'
import { MainCard } from '@/components/card/MainCard'
import { TableCard } from '@/components/card/TableCard'

export default function TopPage() {
  const { data, error, isLoading } = useFetchCards()

  return (
    <>
      <Header />
      {!isLoading && data?.items.map(v => (<MainCard rid={v.rid} />))}
      <TableCard />
    </>
  )
}
