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
  vals: Record<string, string>
}
export const useGetTeapodInfo = (name: string) => query<TeapodInfoSchema>(`teapods/${name}`)

export type TeapodCard = {
  name: string
  title: string
  text: string
}
export const useGetTeapodCard = (teapod: string, name: string) => query<TeapodCard>(`teapods/${teapod}/cards/${name}`)
