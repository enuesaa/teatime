import { useQuery, useMutation, useQueryClient } from 'react-query'
import { ApiBase, ApiListBase } from './schema'

const apiBaseUrl = import.meta.env.VITE_API_BASE_URL

export type TeaSchema = {
  id: string
  value: string
}
export const useListTeas = (teabox: string, teapod: string) =>
  useQuery<ApiListBase<TeaSchema>>('listTeas', async () => {
    const res = await fetch(`${apiBaseUrl}/teapods/${teapod}/teabox/${teabox}/teas`)
    const body = await res.json()
    return body
  })

export const useGetTea = (rid: string) =>
  useQuery<ApiBase<TeaSchema>>(`getTeaInfo-${rid}`, async () => {
    const res = await fetch(`${apiBaseUrl}/teapods/links/teabox/links/teas/${rid}`)
    const body = await res.json()
    return body
  })

// this is for dev.
export type LinksTeaSchema = {
  title: string
  link: string
}
export const useAddTea = () => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (value: LinksTeaSchema) => {
      const res = await fetch(`${apiBaseUrl}/teapods/links/teabox/links/teas`, {
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

export const useDeleteTea = (rid: string) => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async () => {
      const res = await fetch(`${apiBaseUrl}/teapods/links/teabox/links/teas/${rid}`, {
        method: 'DELETE',
      })
      const body = await res.json()
    },
    onSuccess: () => queryClient.invalidateQueries('listTeas'),
  })
}
