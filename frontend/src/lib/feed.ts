import { Feed } from '@/gen/v1/feed_connectweb'
import { queriesInit } from './use'
import { ListFeedsResponse} from '@/gen/v1/feed_pb'

export const {
  useQuery: useListFeedsQuery,
  useLazy: useListFeedsLazy,
} = queriesInit<{}, ListFeedsResponse>(Feed, async (client, arg) => await client.listFeeds(arg))