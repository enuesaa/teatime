import { Header } from '@/components/common/Header'
import { Dashboard as FeedDashboard } from '@/components/feed/Dashboard'
import { Dashboard as BookmarkDashboard } from '@/components/bookmark/Dashboard'
import { Dashboard as BoardDashboard } from '@/components/board/Dashboard'
import { Main } from '@/components/common/Main'

export default function TopPage() {
  return (
    <>
      <Header />
      <Main>
        <BoardDashboard />
        <BookmarkDashboard />
        <FeedDashboard />
      </Main>
    </>
  )
}
