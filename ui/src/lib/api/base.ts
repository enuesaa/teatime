import { useQuery } from 'react-query'

const apiBaseUrl = import.meta.env.API_BASE

export const query = <T>(path: string) =>
  useQuery<T>(path, async () => {
    const res = await fetch(`${apiBaseUrl}/${path}`)
    const body = await res.json()
    return body?.data ?? {}
  })

export const postfn = async (path: string, data: any) => {
  const res = await fetch(`${apiBaseUrl}/${path}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  })
  const body = await res.json()
  return body
}

export const deletefn = async (path: string) => {
  const res = await fetch(`${apiBaseUrl}/${path}`, {
    method: 'DELETE',
  })
  const body = await res.json()
  return body
}
