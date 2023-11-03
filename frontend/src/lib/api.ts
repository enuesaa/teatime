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
