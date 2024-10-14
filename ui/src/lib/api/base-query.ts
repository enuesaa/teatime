import { useQuery } from 'react-query'

const apiBaseUrl = import.meta.env.API_BASE

type QueryConfig = {}

export const query = <T>(method: string, path: string, config: QueryConfig) =>
  useQuery(path, async (): Promise<T> => {
    const res = await fetch(`${apiBaseUrl}${path}`, {
      method,
      headers: {
        Accept: 'application/json',
      },
    })
    const body = await res.json()
    return body?.data ?? {}
  })
