import { useMutation, useQueryClient } from 'react-query'

const apiBaseUrl = import.meta.env.API_BASE

export type MutateConfig = {
  invalidate: string[]
}
export const mutate = <R, T>(method: string, path: string, { invalidate }: MutateConfig) => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: async (data: R): Promise<T> => {
      const res = await fetch(`${apiBaseUrl}${path}`, {
        method,
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      })
      const body = await res.json()
      return body?.data ?? {}
    },
    onSuccess: () => queryClient.invalidateQueries(invalidate),
  })
}
