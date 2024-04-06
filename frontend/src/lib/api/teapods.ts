import { useQuery } from 'react-query'
import { ApiBase, ApiListBase } from './schema'

const apiBaseUrl = import.meta.env.VITE_API_BASE_URL
console.log(import.meta.env)

type TeapodSchema = {
  name: string
  command: string
}
export const useListTeapods = () =>
  useQuery<ApiListBase<TeapodSchema>>('listTeapods', async () => {
    const res = await fetch(`${apiBaseUrl}/api/teapods`)
    const body = await res.json()
    return body
  })

type TeapodInfoSchema = {
  name: string
  command: string
  description: string
  cards: string[]
  schemas: any[]
}
export const useGetTeapodInfo = (name: string) =>
  useQuery<ApiBase<TeapodInfoSchema>>(`getProviderInfo-${name}`, async () => {
    const res = await fetch(`${apiBaseUrl}/api/teapods/${name}`)
    const body = await res.json()
    return body
  })
