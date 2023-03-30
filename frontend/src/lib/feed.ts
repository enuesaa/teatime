import { Feed } from '@/gen/v1/feed_connectweb'
import { queriesInit } from './use'
import { AddFeedRequest, AddFeedResponse, DeleteFeedRequest, DeleteFeedResponse, FetchRequest, FetchResponse, GetFeedRequest, GetFeedResponse, ListFeedsResponse, ListItemsRequest, ListItemsResponse} from '@/gen/v1/feed_pb'

export const {
  useQuery: useListFeedsQuery,
  useLazy: useListFeedsLazy,
} = queriesInit<{}, ListFeedsResponse>(Feed, async (client, arg) => await client.listFeeds(arg))

export const {
  useQuery: useGetFeedQuery,
  useLazy: useGetFeedLazy,
} = queriesInit<GetFeedRequest, GetFeedResponse>(Feed, async (client, arg) => await client.getFeed(arg))

export const {
  useQuery: useAddFeedQuery,
  useLazy: useAddFeedLazy,
} = queriesInit<AddFeedRequest, AddFeedResponse>(Feed, async (client, arg) => await client.addFeed(arg))

export const {
  useQuery: useDeleteFeedQuery,
  useLazy: useDeleteFeedLazy,
} = queriesInit<DeleteFeedRequest, DeleteFeedResponse>(Feed, async (client, arg) => await client.deleteFeed(arg))

export const {
  useQuery: useFetchQuery,
  useLazy: useFetchLazy,
} = queriesInit<FetchRequest, FetchResponse>(Feed, async (client, arg) => await client.fetch(arg))

export const {
  useQuery: useListItemsQuery,
  useLazy: useListItemsLazy,
} = queriesInit<ListItemsRequest, ListItemsResponse>(Feed, async (client, arg) => await client.listItems(arg))
