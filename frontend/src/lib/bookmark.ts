import { Bookmark } from '@/gen/v1/bookmark_connectweb'
import { queriesInit } from './use'
import { GetBookmarkResponse, ListBookmarksResponse } from '@/gen/v1/bookmark_pb'

export const {
  useQuery: useGetBookmarkQuery,
  useLazy: useGetBookmarkLazy,
} = queriesInit<GetBookmarkResponse>(Bookmark, async (client) => await client.getBookmark({ id: 'aaa' }))

export const {
  useQuery: useListBookmarksQuery,
  useLazy: useListBookmarksLazy,
} = queriesInit<ListBookmarksResponse>(Bookmark, async (client) => await client.listBookmarks())