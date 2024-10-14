import { mutatePost, queryGet } from './base'

type TeapodSchema = {
  name: string
  command: string
}
export const useListTeapods = () => queryGet<TeapodSchema[]>('api/teapods')

type TeapodInfoSchema = {
  name: string
  description: string
  teaboxes: TeapodInfoTeabox[]
}
export type TeapodInfoTeabox = {
  name: string
}
export const useGetTeapodInfo = (name: string) => queryGet<TeapodInfoSchema>(`api/teapods/${name}/info`)

export type AddReqSchema = {
  name: string
}
export const useAddTeapod = () =>
  mutatePost<AddReqSchema, {}>('api/teapods', {
    invalidate: [],
  })
