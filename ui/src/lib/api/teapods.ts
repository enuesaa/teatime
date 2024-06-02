import { query } from './base'

type TeapodSchema = {
  name: string
  command: string
}
export const useListTeapods = () => query<TeapodSchema[]>('teapods')

type TeapodInfoSchema = {
  name: string
  command: string
  description: string
  cards: string[]
  teaboxes: TeapodInfoTeabox[]
}
export type TeapodInfoTeabox = {
  name: string
  comment: string
  valdefs: TeapodInfoTeaboxValdef[]
}
export type TeapodInfoTeaboxValdef = {
  name: string
  cast: string
  nullable: boolean
}
export const useGetTeapodInfo = (name: string) => query<TeapodInfoSchema>(`teapods/${name}`)
