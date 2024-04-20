import { useSearchParams } from 'react-router-dom'

export const useQueryStr = (name: string): string | null => {
  const [searchParams] = useSearchParams()
  return searchParams.get(name)
}
