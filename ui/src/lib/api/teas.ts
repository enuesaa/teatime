import { useQuery, useMutation, useQueryClient } from 'react-query'
import { ApiBase, ApiListBase } from './schema'

export type TeaSchema = {
  id: string
  value: string
}
export const useListTeas = () =>
  useQuery<ApiListBase<TeaSchema>>('listTeas', async () => {
    const res = await fetch(`${API_BASE_URL}/teapods/links/teas`)
    return await res.json()
  })

export const useGetTea = (rid: string) =>
  useQuery<ApiBase<TeaSchema>>(`getTeaInfo-${rid}`, async () => {
    const res = await fetch(`${API_BASE_URL}/teapods/links/teas/${rid}`)
    return await res.json()
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
      const res = await fetch(`${API_BASE_URL}/teapods/links/teas`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(value),
      })
      const body = await res.json()
    },
    onSuccess: () => queryClient.invalidateQueries('listTeas'),
  })
}

export const useDeleteTea = (rid: string) => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (): Promise<void> => {
      const res = await fetch(`${API_BASE_URL}/teapods/links/teas/${rid}`, {
        method: 'DELETE',
      })
      const body = await res.json()
    },
    onSuccess: () => queryClient.invalidateQueries('listTeas'),
  })
}
