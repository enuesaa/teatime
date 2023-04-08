import { Board } from '@/gen/v1/board_connectweb'
import { queriesInit } from './use'
import { AddBoardRequest, ListBoardsResponse, CheckinRequest, ArchiveBoardRequest, UnArchiveBoardRequest, ListTimelineRequest } from '@/gen/v1/board_pb'

/**
 * @todo list board
 */
export const {
  useQuery: useListBoardsQuery,
  useLazy: useListBoardsLazy,
} = queriesInit<{}, ListBoardsResponse>(Board, async(client, arg) => await client.listBoards(arg))

export const {
  useQuery: useGetBookmarkQuery,
  useLazy: useGetBookmarkLazy,
} = queriesInit<AddBoardRequest, {}>(Board, async (client, arg) => await client.addBoard(arg))

export const {
  useQuery: useCheckinQuery,
  useLazy: useCheckinLazy,
} = queriesInit<CheckinRequest, {}>(Board, async (client, arg) => await client.checkin(arg))

export const {
  useQuery: useListTimelineQuery,
  useLazy: useListTimelineLazy,
} = queriesInit<ListTimelineRequest, ListBoardsResponse>(Board, async (client, arg) => await client.listTimeline(arg))

export const {
  useQuery: useArchiveBoardQuery,
  useLazy: useArchiveBoardLazy,
} = queriesInit<ArchiveBoardRequest, {}>(Board, async (client, arg) => await client.archive(arg))

export const {
  useQuery: useUnArchiveBoardQuery,
  useLazy: useUnArchiveBoardLazy,
} = queriesInit<UnArchiveBoardRequest, {}>(Board, async (client, arg) => await client.unArchive(arg))
