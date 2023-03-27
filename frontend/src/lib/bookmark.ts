import { createClient } from '@/lib/client'
import { Bookmark } from '../../gen/v1/bookmark_connectweb'
import { useState, useEffect } from 'react'

export const useGetBookmarksQuery = () => {
  const [data, setData] = useState<any>(null)

  useEffect(() => {
    (async () => {
      const client = createClient(Bookmark)
      const res = await client.getBookmark({})
      setData(res)
    })()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  return data
}
