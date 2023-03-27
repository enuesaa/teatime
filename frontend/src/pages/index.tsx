import { Header } from '@/components/common/Header'
import { Timeline } from '@/components/feed/Timeline'
import { TopDashboard } from '@/components/bookmark/TopDashboard'

export default function TopPage() {
  return (
    <>
      <Header />
      <TopDashboard />
      <Timeline />
    </>
  )
}
