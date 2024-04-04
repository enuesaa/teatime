import { useQuery, useMutation, useQueryClient } from 'react-query'
import { ApiBase, ApiListBase } from './schema'

const backendApiHost = import.meta.env.BASE_URL

export type TeaSchema = {
  id: string
  value: string
}

// this is for dev.
export type LinksTeaSchema = {
  title: string
  link: string
}

export type SchemaSchema = {
  name: string
  vals: Record<string, 'str' | 'bool' | 'int'>
}

export const useListSchemas = () =>
  useQuery('listSchemas', async (): Promise<ApiListBase<SchemaSchema>> => {
    const res = await fetch(`http://${backendApiHost}/schemas`)
    const body = await res.json()
    return body as ApiListBase<SchemaSchema>
  })

export const useListTeas = () =>
  useQuery('listTeas', async (): Promise<ApiListBase<TeaSchema>> => {
    const res = await fetch(`http://${backendApiHost}/teas`)
    const body = await res.json()
    return body as ApiListBase<TeaSchema>
  })

export const useGetTeaInfo = (rid: string) =>
  useQuery(`getTeaInfo-${rid}`, async (): Promise<ApiBase<TeaSchema>> => {
    const res = await fetch(`http://${backendApiHost}/teas/${rid}`)
    const body = await res.json()
    return body as ApiBase<TeaSchema>
  })

export const useAddTea = () => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (value: LinksTeaSchema) => {
      const res = await fetch(`http://${backendApiHost}/teas`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(value),
      })
      const body = await res.json()
      console.log(body)
    },
    onSuccess: () => queryClient.invalidateQueries('listTeas'),
  })
}

export const useDeleteTea = (rid: string) => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (): Promise<void> => {
      const res = await fetch(`http://${backendApiHost}/teas/${rid}`, {
        method: 'DELETE',
      })
      const body = await res.json()
      console.log(body)
    },
    onSuccess: () => queryClient.invalidateQueries('listTeas'),
  })
}
