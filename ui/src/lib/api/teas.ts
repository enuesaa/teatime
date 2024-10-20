import { mutateDelete, mutatePost, mutatePut, queryGet } from './base'

export type TeaSchema = {
  id: string
  created: string
  updated: string
  data: Record<string, any>
}
export const useListTeas = (teapod: string, teabox: string) =>
  queryGet<{items: TeaSchema[]}>(`api/teapods/${teapod}/teaboxes/${teabox}/teas`)

export const useGetTea = (teapod: string, teabox: string, teaId: string) =>
  queryGet<TeaSchema>(`api/teapods/${teapod}/teaboxes/${teabox}/teas/${teaId}`)

export type AddReqSchema = Record<string, any>
export type AddResSchema = {
  error?: string
}
export const useAddTea = (teapod: string, teabox: string) =>
  mutatePost<AddReqSchema, AddResSchema>(`api/teapods/${teapod}/teaboxes/${teabox}/teas`, {
    invalidate: [],
  })

export const useUpdateTea = (teapod: string, teabox: string, teaId: string) =>
  mutatePut<AddReqSchema, AddResSchema>(`api/teapods/${teapod}/teaboxes/${teabox}/teas/${teaId}`, {
    invalidate: [],
  })  

export const useDeleteTea = (teapod: string, teabox: string, teaId: string) =>
  mutateDelete<{}, {}>(`api/teapods/${teapod}/teaboxes/${teabox}/teas/${teaId}`, {
    invalidate: [],
  })
