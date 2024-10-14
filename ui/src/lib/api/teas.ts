import { queryGet } from './base'

export type TeaSchema = {
  teaid: string
  teabox: string
  vals: TeaVal[]
}
export type TeaVal = {
  name: string
  value: string
}
export const useListTeas = (teapod: string, teabox: string) =>
  queryGet<TeaSchema[]>(`/api/teapods/${teapod}/teaboxes/${teabox}/teas`)
