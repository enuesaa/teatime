import { useQuery, useMutation, useQueryClient } from 'react-query'
import { ApiBase, ApiListBase, CardSchema, InfoSchema, PanelSchema, ProviderSchema, RecordSchema } from './schema'
const backendApiHost = import.meta.env.VITE_BACKEND_API_HOST as string

export const useListProviders = () => useQuery('listProviders', async (): Promise<ApiListBase<ProviderSchema>> => {
  const res = await fetch(`http://${backendApiHost}/providers`)
  const body = await res.json()
  return body as ApiListBase<ProviderSchema>
})

export const useAddProvider = () => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (reqbody: ProviderSchema) => {
      const res = await fetch(`http://${backendApiHost}/providers`, {
        method: 'POST',
        body: JSON.stringify(reqbody),
      })
      const body = await res.json()
      console.log(body)
    },
    onSuccess: () => queryClient.invalidateQueries('listProviders')
  })
}

export const useUpdateProvider = (id: string) => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (reqbody: ProviderSchema) => {
      const res = await fetch(`http://${backendApiHost}/providers/${id}`, {
        method: 'PUT',
        body: JSON.stringify(reqbody),
      })
      const body = await res.json()
      console.log(body)
    },
    onSuccess: () => queryClient.invalidateQueries('listProviders')
  })
}

export const useDeleteProvider = (id: string) => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (): Promise<void> => {
      const res = await fetch(`http://${backendApiHost}/providers/${id}`, {
        method: 'DELETE',
      })
      const body = await res.json()
      console.log(body)
    },
    onSuccess: () => queryClient.invalidateQueries('listProviders')
  })
}

export const useGetProviderInfo = (name: string) => useQuery('getProviderInfo', async (): Promise<ApiBase<InfoSchema>> => {
  const res = await fetch(`http://${backendApiHost}/providers/${name}`)
  const body = await res.json()
  return body as ApiBase<InfoSchema>
})

export const useGetProviderCard = (name: string, cardName: string) => useQuery('getProviderCard', async (): Promise<CardSchema> => {
  const res = await fetch(`http://${backendApiHost}/providers/${name}/cards/${cardName}`)
  const body = await res.json()
  return body.data as CardSchema
})

export const useGetProviderPanel = (name: string, panelName: string) => useQuery('getProviderPanel', async (): Promise<PanelSchema> => {
  const res = await fetch(`http://${backendApiHost}/providers/${name}/panels/${panelName}`)
  const body = await res.json()
  return body.data as PanelSchema
})

export const useGetProviderRecord = (name: string, modelName: string, recordName: string) => useQuery('getProviderRecord', async (): Promise<RecordSchema> => {
  const res = await fetch(`http://${backendApiHost}/providers/${name}/models/${modelName}/records/${recordName}`)
  const body = await res.json()
  return body.data as RecordSchema
})

export const useRegisterRecord = (name: string, modelName: string) => useMutation(
  async (recordName: string): Promise<string> => {
    const res = await fetch(`http://${backendApiHost}/providers/${name}/models/${modelName}/records/${recordName}`, {
      method: 'POST',
    })
    const body = await res.json()
    console.log(body)
    return recordName
  },
)

export const useUpdateRecord = (name: string, modelName: string) => useMutation(
  async ({recordName, record}: {recordName: string, record: RecordSchema}): Promise<string> => {
    const res = await fetch(`http://${backendApiHost}/providers/${name}/models/${modelName}/records/${recordName}`, {
      method: 'PUT',
      body: JSON.stringify(record),
    })
    const body = await res.json()
    console.log(body)
    return recordName
  },
)

export const useDeleteRecord = (name: string, modelName: string) => useMutation(
  async (recordName: string): Promise<string> => {
    const res = await fetch(`http://${backendApiHost}/providers/${name}/models/${modelName}/records/${recordName}`, {
      method: 'DELETE',
    })
    const body = await res.json()
    console.log(body)
    return recordName
  },
)
