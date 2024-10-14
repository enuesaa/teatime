import { mutatePost, queryGet } from './base'

export type TeaSchema = {
  id: string
  data: Record<string, any>
}
export const useListTeas = (teapod: string, teabox: string) =>
  queryGet<TeaSchema[]>(`api/teapods/${teapod}/teaboxes/${teabox}/teas`)


export type AddReqSchema = Record<string, any>
export const useAddTea = (teapod: string, teabox: string) =>
  mutatePost<AddReqSchema, {}>(`api/teapods/${teapod}/teaboxes/${teabox}/teas`, {
    invalidate: [],
  })
