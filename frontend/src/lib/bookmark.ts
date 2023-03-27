import { Bookmark } from '../../gen/v1/bookmark_connectweb'
import { queriesInit } from './use'

export const {
  useQuery: useGetBookmarksQuery,
  useLazy: useGetBookmarksLazy,
} = queriesInit(Bookmark, async (client) => await client.getBookmark())