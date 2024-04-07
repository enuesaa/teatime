import { useQuery, useMutation, useQueryClient } from 'react-query'
import { ApiBase, ApiListBase } from './schema'

const apiBaseUrl = import.meta.env.VITE_API_BASE_URL

export type TeaSchema = {
  id: string
  value: string
}
export const useListTeas = (teapod: string, teabox: string) =>
  useQuery<ApiListBase<TeaSchema>>(`listTeas-${teapod}-${teabox}`, async () => {
    const res = await fetch(`${apiBaseUrl}/teapods/${teapod}/teabox/${teabox}/teas`)
    const body = await res.json()
    return body
  })

export const useGetTea = (teapod: string, teabox: string, teaid: string) =>
  useQuery<ApiBase<TeaSchema>>(`getTeaInfo-${teapod}-${teabox}-${teaid}`, async () => {
    const res = await fetch(`${apiBaseUrl}/teapods/${teapod}/teabox/${teabox}/teas/${teaid}`)
    const body = await res.json()
    return body
  })

// this is for dev.
export type LinksTeaSchema = {
  title: string
  link: string
}
export const useAddTea = (teapod: string, teabox: string) => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (value: LinksTeaSchema) => {
      const res = await fetch(`${apiBaseUrl}/teapods/${teapod}/teabox/${teabox}/teas`, {
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

export const useDeleteTea = (teapod: string, teabox: string, teaid: string) => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async () => {
      const res = await fetch(`${apiBaseUrl}/teapods/${teapod}/teabox/${teabox}/teas/${teaid}`, {
        method: 'DELETE',
      })
      const body = await res.json()
    },
    onSuccess: () => queryClient.invalidateQueries('listTeas'),
  })
}
