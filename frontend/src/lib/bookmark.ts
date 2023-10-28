import { queriesInit } from './use'

const Bookmark = {}

export const {
  useQuery: useGetBookmarkQuery,
  useLazy: useGetBookmarkLazy,
} = queriesInit<{}, {}>(Bookmark, async (client, arg) => await client.getBookmark(arg))

export const {
  useQuery: useListBookmarksQuery,
  useLazy: useListBookmarksLazy,
} = queriesInit<{}, {}>(Bookmark, async (client, arg) => await client.listBookmarks(arg))

export const {
  useQuery: useAddBookmarkQuery,
  useLazy: useAddBookmarkLazy,
} = queriesInit<{}, {}>(Bookmark, async (client, arg) => await client.addBookmark(arg))

export const {
  useQuery: useDeleteBookmarkQuery,
  useLazy: useDeleteBookmarkLazy,
} = queriesInit<{}, {}>(Bookmark, async (client, arg) => await client.deleteBookmark(arg))
