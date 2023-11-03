import { Header } from '@/components/common/Header'
import { Container } from '@radix-ui/themes'
import { useFetchCards } from '@/lib/api'
import { MainCard } from '@/components/card/MainCard'

export default function TopPage() {
  const { data, error, isLoading } = useFetchCards()

  return (
    <>
      <Header />
      {!isLoading && data?.items.map(v => (<MainCard rid={v.rid} />))}
    </>
  )
}
