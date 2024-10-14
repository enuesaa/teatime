import { useMutation, useQuery, useQueryClient } from 'react-query'

const apiBaseUrl = import.meta.env.API_BASE

export const query = <T>(path: string) =>
  useQuery(path, async (): Promise<T> => {
    const res = await fetch(`${apiBaseUrl}/${path}`)
    const body = await res.json()
    return body?.data ?? {}
  })

type MutateConfig = {
  path: string
  invalidate: string[]
}
const mutate = <R, T>(method: string, { path, invalidate }: MutateConfig) => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: async (data: R): Promise<T> => {
      const res = await fetch(`${apiBaseUrl}/${path}`, {
        method,
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      })
      const body = await res.json()
      return body
    },
    onSuccess: () => queryClient.invalidateQueries(invalidate),
  })
}

export const mutatePost = <R, T>(config: MutateConfig) => mutate<R, T>('POST', config)
