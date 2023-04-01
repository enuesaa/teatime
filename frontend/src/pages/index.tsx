import { Header } from '@/components/common/Header'
import { Timeline } from '@/components/feed/Timeline'
import { TopDashboard as BookmarkDashboard } from '@/components/bookmark/TopDashboard'
import { TopDashboard as BoardDashboard } from '@/components/board/TopDashboard'
import { Main } from '@/components/common/Main'

export default function TopPage() {
  return (
    <>
      <Header />
      <Main>
        <BoardDashboard />
        <BookmarkDashboard />
        <Timeline />
      </Main>
    </>
  )
}
