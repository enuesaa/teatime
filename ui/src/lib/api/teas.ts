import { useMutation, useQueryClient } from 'react-query'
import { query, postfn, deletefn } from './base'

const apiBaseUrl = import.meta.env.API_BASE

export type TeaSchema = {
  teaid: string
  teabox: string
  vals: Record<string, string>
}
export const useListTeas = (teapod: string, teabox: string) => query<TeaSchema[]>(`teapods/${teapod}/teas?teabox=${teabox}`)

export const useGetTea = (teapod: string, teaid: string) => query<TeaSchema>(`teapods/${teapod}/teas/${teaid}`)

export type CreateTeaReq = {
  teabox: string
  vals: Record<string, string>
}
export const useAddTea = (teapod: string, teabox: string) => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (value: CreateTeaReq) => {
      value.teabox = teabox
      return await postfn(`teapods/${teapod}/teas`, value)
    },
    onSuccess: () => queryClient.invalidateQueries(`listTeas-${teapod}`),
  })
}

export const useDeleteTea = (teapod: string, teaid: string) => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async () => await deletefn(`teapods/${teapod}/teas/${teaid}`),
    onSuccess: () => queryClient.invalidateQueries(`listTeas-${teapod}`),
  })
}
