import { queriesInit } from './use'

const Board = {}

/**
 * @todo list board
 */
export const {
  useQuery: useListBoardsQuery,
  useLazy: useListBoardsLazy,
} = queriesInit<{}, {}>(Board, async(client, arg) => await client.listBoards(arg))

export const {
  useQuery: useGetBookmarkQuery,
  useLazy: useGetBookmarkLazy,
} = queriesInit<{}, {}>(Board, async (client, arg) => await client.addBoard(arg))

export const {
  useQuery: useCheckinQuery,
  useLazy: useCheckinLazy,
} = queriesInit<{}, {}>(Board, async (client, arg) => await client.checkin(arg))

export const {
  useQuery: useListTimelineQuery,
  useLazy: useListTimelineLazy,
} = queriesInit<{}, {}>(Board, async (client, arg) => await client.listTimeline(arg))

export const {
  useQuery: useArchiveBoardQuery,
  useLazy: useArchiveBoardLazy,
} = queriesInit<{}, {}>(Board, async (client, arg) => await client.archive(arg))

export const {
  useQuery: useUnArchiveBoardQuery,
  useLazy: useUnArchiveBoardLazy,
} = queriesInit<{}, {}>(Board, async (client, arg) => await client.unArchive(arg))
