import { useQuery } from 'react-query'
import { ApiBase } from './schema'

const apiBaseUrl = import.meta.env.VITE_API_BASE_URL

type TeapodSchema = {
  name: string
  command: string
}
export const useListTeapods = () =>
  useQuery<TeapodSchema[]>('listTeapods', async () => {
    const res = await fetch(`${apiBaseUrl}/teapods`)
    const body = await res.json()
    return body?.data ?? []
  })

type TeapodInfoSchema = {
  name: string
  command: string
  description: string
  cards: string[]
  teaboxes: TeapodInfoTeabox[]
}
type TeapodInfoTeabox = {
  name: string
  vals: Record<string, string>
}
export const useGetTeapodInfo = (name: string) =>
  useQuery<TeapodInfoSchema>(`getProviderInfo-${name}`, async () => {
    const res = await fetch(`${apiBaseUrl}/teapods/${name}`)
    const body = await res.json()
    return body?.data ?? {}
  })
