import { Board } from '@/gen/v1/board_connectweb'
import { queriesInit } from './use'
import { AddBoardRequest } from '@/gen/v1/board_pb'

/**
 * @todo list board
 */

export const {
  useQuery: useGetBookmarkQuery,
  useLazy: useGetBookmarkLazy,
} = queriesInit<AddBoardRequest, {}>(Board, async (client, arg) => await client.addBoard(arg))
