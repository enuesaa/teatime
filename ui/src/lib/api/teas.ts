import { useQuery, useMutation, useQueryClient } from 'react-query'

const apiBaseUrl = import.meta.env.API_BASE

export type TeaSchema = {
  teaid: string
  teabox: string
  vals: Record<string, string>
}
export const useListTeas = (teapod: string, teabox: string) =>
  useQuery<TeaSchema[]>([`listTeas-${teapod}`, teabox], async () => {
    const res = await fetch(`${apiBaseUrl}/teapods/${teapod}/teas?teabox=${teabox}`)
    const body = await res.json()
    return body?.data ?? []
  })

export const useGetTea = (teapod: string, teaid: string) =>
  useQuery<TeaSchema>(`getTeaInfo-${teapod}-${teaid}`, async () => {
    const res = await fetch(`${apiBaseUrl}/teapods/${teapod}/teas/${teaid}`)
    const body = await res.json()
    return body?.data ?? {}
  })

export type CreateTeaReq = {
  teabox: string
  vals: Record<string, string>
}
export const useAddTea = (teapod: string, teabox: string) => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (value: CreateTeaReq) => {
      value.teabox = teabox
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
    onSuccess: () => queryClient.invalidateQueries(`listTeas-${teapod}`),
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
    onSuccess: () => queryClient.invalidateQueries(`listTeas-${teapod}`),
  })
}
