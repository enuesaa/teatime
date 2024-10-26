import { mutatePost } from './base'

export const useDeleteAllLogs = () =>
  mutatePost<{}, {}>(`api/log/delete-all`, {
    invalidate: [],
  })
