import { Bookmark } from '@/gen/v1/bookmark_connectweb'
import { queriesInit } from './use'
import { AddBookmarkRequest, AddBookmarkResponse, GetBookmarkResponse, ListBookmarksResponse } from '@/gen/v1/bookmark_pb'

export const {
  useQuery: useGetBookmarkQuery,
  useLazy: useGetBookmarkLazy,
} = queriesInit<{}, GetBookmarkResponse>(Bookmark, async (client, arg) => await client.getBookmark(arg))

export const {
  useQuery: useListBookmarksQuery,
  useLazy: useListBookmarksLazy,
} = queriesInit<{}, ListBookmarksResponse>(Bookmark, async (client, arg) => await client.listBookmarks(arg))

export const {
  useQuery: useAddBookmarkQuery,
  useLazy: useAddBookmarkLazy,
} = queriesInit<AddBookmarkRequest, AddBookmarkResponse>(Bookmark, async (client, arg) => await client.addBookmark(arg))