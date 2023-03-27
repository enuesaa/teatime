import { Bookmark } from '@/gen/v1/bookmark_connectweb'
import { queriesInit } from './use'
import { GetBookmarkResponse } from '@/gen/v1/bookmark_pb'

export const {
  useQuery: useGetBookmarksQuery,
  useLazy: useGetBookmarksLazy,
} = queriesInit<GetBookmarkResponse>(Bookmark, async (client) => await client.getBookmark())