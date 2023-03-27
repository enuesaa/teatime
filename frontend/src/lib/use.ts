import { useState, useEffect } from 'react'
import { createClient } from '@/lib/client'

type CommandName = any
type RetrieveFn = (client: any) => Promise<any>;

const queryInit = <R>(cmd: CommandName, fn: RetrieveFn) => () => {
  const [data, setData] = useState<R|null>(null)
  useEffect(() => {
    (async () => {
      const client = createClient(cmd)
      const res = await fn(client)
      setData(res)
    })()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  return data
}

const lazyInit = <R>(cmd: CommandName, fn: RetrieveFn) => () => {
  const [data, setData] = useState<R|null>(null)
  const invoke = async () => {
    const client = createClient(cmd)
    const res = await fn(client)
    setData(res)
  }

  return { data, invoke }
}

type QueryMap<R> = {
  Query: () => R;
  Lazy: () => { data: null | R, invoke: () => void};
}

type SnakeToCamel<S extends string> = S extends `${infer T}_${infer U}` ? `${Capitalize<T>}${SnakeToCamel<U>}` : `${Capitalize<S>}`
export const queriesInit = <R>(cmd: CommandName, fn: RetrieveFn) => {
  const queryFn = queryInit<R>(cmd, fn)
  const lazyFn = lazyInit<R>(cmd, fn)

  return {
    [`useQuery`]: queryFn,
    [`useLazy`]: lazyFn,
  }
}