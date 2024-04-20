import { useQuery } from 'react-query'

const apiBaseUrl = import.meta.env.API_BASE

export const query = <T>(path: string) => useQuery<T>(path, async () => {
  const res = await fetch(`${apiBaseUrl}/${path}`)
  const body = await res.json()
  return body?.data ?? {}
})
