import { Header } from '@/components/common/Header'
import { Timeline } from '@/components/feed/Timeline'
import { TopDashboard as BookmarkDashboard } from '@/components/bookmark/TopDashboard'
import { TopDashboard as BoardDashboard } from '@/components/board/TopDashboard'

export default function TopPage() {
  return (
    <>
      <Header />
      <BoardDashboard />
      <BookmarkDashboard />
      <Timeline />
    </>
  )
}
