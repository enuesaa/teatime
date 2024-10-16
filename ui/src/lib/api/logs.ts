import { queryGet } from './base'

export type LogSchema = {
  created: string
  message: string
}
export const useListLogs = () =>
  queryGet<LogSchema[]>(`api/logs`)
