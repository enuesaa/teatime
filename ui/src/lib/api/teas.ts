import { useQuery, useMutation, useQueryClient } from 'react-query'
import { ApiBase, ApiListBase } from './schema'

const apiBaseUrl = import.meta.env.VITE_API_BASE_URL

export type TeaSchema = {
  id: string
  value: string
}
export const useListTeas = (teapod: string) =>
  useQuery<ApiListBase<TeaSchema>>(`listTeas-${teapod}`, async () => {
    const res = await fetch(`${apiBaseUrl}/teapods/${teapod}/teas`)
    const body = await res.json()
    return body
  })

export const useGetTea = (teapod: string, teaid: string) =>
  useQuery<ApiBase<TeaSchema>>(`getTeaInfo-${teapod}-${teaid}`, async () => {
    const res = await fetch(`${apiBaseUrl}/teapods/${teapod}/teas/${teaid}`)
    const body = await res.json()
    return body
  })

// this is for dev.
export type LinksTeaSchema = {
  teabox: string
  vals: Record<string, string>
}
export const useAddTea = (teapod: string) => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (value: LinksTeaSchema) => {
      value.teabox = 'links'
      const res = await fetch(`${apiBaseUrl}/teapods/${teapod}/teas`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(value),
      })
      const body = await res.json()
      return body
    },
    onSuccess: () => queryClient.invalidateQueries('listTeas'),
  })
}

export const useDeleteTea = (teapod: string, teaid: string) => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async () => {
      const res = await fetch(`${apiBaseUrl}/teapods/${teapod}/teas/${teaid}`, {
        method: 'DELETE',
      })
      const body = await res.json()
    },
    onSuccess: () => queryClient.invalidateQueries('listTeas'),
  })
}
