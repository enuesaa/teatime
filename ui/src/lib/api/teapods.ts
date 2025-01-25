import { mutateDelete, mutatePost, queryGet } from './base'

type TeapodSchema = {
  name: string
  command: string
}
export const useListTeapods = () => queryGet<{ items: TeapodSchema[] }>('api/teapods')

type TeapodInfoSchema = {
  name: string
  description: string
  teaboxes: TeapodInfoTeabox[]
  actions: TeapodInfoAction[]
}
export type TeapodInfoTeabox = {
  name: string
  inputs: {name: string; type: 'text'}[]
}
export type TeapodInfoAction = {
  name: string
  comment: string
}
export const useGetTeapodInfo = (name: string) => queryGet<TeapodInfoSchema>(`api/teapods/${name}/info`)

export type AddReqSchema = {
  name: string
}
export type AddResSchema = {
  error?: string
}
export const useAddTeapod = () =>
  mutatePost<AddReqSchema, AddResSchema>('api/teapods', {
    invalidate: [],
  })

export const useDeleteTeapod = (name: string) =>
  mutateDelete<{}, {}>(`api/teapods/${name}`, {
    invalidate: [],
  })

export type ActReqSchema = {
  action: string
}
export type ActResSchema = {
  message: string
}
export const useActTeapod = (name: string) => 
  mutatePost<ActReqSchema, ActResSchema>(`api/teapods/${name}/act`, {
    invalidate: true,
  })
