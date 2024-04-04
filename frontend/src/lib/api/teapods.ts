import { useQuery } from 'react-query'
import { ApiBase, ApiListBase } from './schema'
const backendApiHost = 'localhost:3000/api'

export type ProviderSchema = {
  id: string
  name: string
  command: string
}

export type InfoSchema = {
  name: string
  command: string
  description: string
  cards: string[]
}

export const useListTeapods = () =>
  useQuery('listProviders', async (): Promise<ApiListBase<ProviderSchema>> => {
    const res = await fetch(`http://${backendApiHost}/teapods`)
    const body = await res.json()
    return body as ApiListBase<ProviderSchema>
  })

export const useGetProviderInfo = (name: string) =>
  useQuery('getProviderInfo', async (): Promise<ApiBase<InfoSchema>> => {
    const res = await fetch(`http://${backendApiHost}/providers/${name}`)
    const body = await res.json()
    return body as ApiBase<InfoSchema>
  })
