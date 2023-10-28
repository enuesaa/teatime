import { queriesInit } from './use'

const Feed = {}

export const {
  useQuery: useListFeedsQuery,
  useLazy: useListFeedsLazy,
} = queriesInit<{}, {}>(Feed, async (client, arg) => await client.listFeeds(arg))

export const {
  useQuery: useGetFeedQuery,
  useLazy: useGetFeedLazy,
} = queriesInit<{}, {}>(Feed, async (client, arg) => await client.getFeed(arg))

export const {
  useQuery: useAddFeedQuery,
  useLazy: useAddFeedLazy,
} = queriesInit<{}, {}>(Feed, async (client, arg) => await client.addFeed(arg))

export const {
  useQuery: useDeleteFeedQuery,
  useLazy: useDeleteFeedLazy,
} = queriesInit<{}, {}>(Feed, async (client, arg) => await client.deleteFeed(arg))

export const {
  useQuery: useFetchQuery,
  useLazy: useFetchLazy,
} = queriesInit<{}, {}>(Feed, async (client, arg) => await client.fetch(arg))

export const {
  useQuery: useListAllItemsQuery,
  useLazy: useListAllItemsLazy,
} = queriesInit<{}, {}>(Feed, async (client, arg) => await client.listAllItems(arg))

export const {
  useQuery: useListItemsQuery,
  useLazy: useListItemsLazy,
} = queriesInit<{}, {}>(Feed, async (client, arg) => await client.listItems(arg))

export const {
  useQuery: useRemoveAllItemsQuery,
  useLazy: useRemoveAllItemsLazy,
} = queriesInit<{}, {}>(Feed, async (client, arg) => await client.removeAllItems(arg))
