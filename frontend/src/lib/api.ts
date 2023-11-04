import { useQuery } from 'react-query'

export type CardItem = {
  layout: string
  rid: string
}
export type Cards = {
  items: CardItem[]
}
const fetchCards = async (): Promise<Cards> => {
  const res = await fetch('http://localhost:3000/cards')
  const body = await res.json()
  return body as Cards
}
export const useFetchCards = () => useQuery('getCards', fetchCards)

export const useGetProviderInfo = (name: string) => useQuery('getProviderInfo', async (): Promise<InfoSchema> => {
  const res = await fetch(`http://localhost:3000/providers/${name}`)
  const body = await res.json()
  return body.data as InfoSchema
})

export const useGetProviderCard = (name: string, cardName: string) => useQuery('getProviderCard', async (): Promise<CardSchema> => {
  const res = await fetch(`http://localhost:3000/providers/${name}/cards/${cardName}`)
  const body = await res.json()
  return body.data as CardSchema
})

export const useGetProviderPanel = (name: string, panelName: string) => useQuery('getProviderPanel', async (): Promise<PanelSchema> => {
  const res = await fetch(`http://localhost:3000/providers/${name}/panels/${panelName}`)
  const body = await res.json()
  return body.data as PanelSchema
})

export const useGetProviderRecord = (name: string, modelName: string, recordName: string) => useQuery('getProviderRecord', async (): Promise<RecordSchema> => {
  const res = await fetch(`http://localhost:3000/providers/${name}/models/${modelName}/records/${recordName}`)
  const body = await res.json()
  return body.data as RecordSchema
})

// 	app.POST("/providers/:name/models/:model/records/:recordName", controller.RegisterRecord)
// export const useGetProviderRecord = (name: string, modelName: string, recordName: string) => useQuery('getProviderRecord', async (): Promise<RecordSchema> => {
//   const res = await fetch(`http://localhost:3000/providers/${name}/models/${modelName}/records/${recordName}`)
//   const body = await res.json()
//   return body.data as RecordSchema
// })
// 	app.PUT("/providers/:name/models/:model/records/:recordName", controller.SetRecord)
// 	app.DELETE("/providers/:name/models/:model/records/:recordName", controller.DelRecord)
