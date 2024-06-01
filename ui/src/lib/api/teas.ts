import { query, postfn } from './base'
import { useMutation, useQueryClient } from 'react-query'

export type TeaSchema = {
  teaid: string
  teabox: string
  vals: Record<string, string>
}
export const useListTeas = (teapod: string, teabox: string) =>
  query<TeaSchema[]>(`teapods/${teapod}/teas?teabox=${teabox}`)

// export type CreateTeaReq = {
//   teabox: string
//   vals: Record<string, string>
// }
// export const useAct = (teapod: string, teabox: string) => {
//   const queryClient = useQueryClient()
//   return useMutation({
//     mutationFn: async (value: CreateTeaReq) => {
//       value.teabox = teabox
//       return await postfn(`teapods/${teapod}/teas`, value)
//     },
//     onSuccess: () => queryClient.invalidateQueries(`listTeas-${teapod}`),
//   })
// }
