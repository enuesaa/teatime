import { useSearchParams } from 'react-router-dom'

export const useQueryStr = (name: string): undefined|string => {
  const [searchParams] = useSearchParams()
  const value = searchParams.get(name)
  return value ?? undefined
}
