import { useState, useEffect } from 'react'

type CallFn<A> = (client: any, arg: A) => Promise<any>;
const queryInit = <A, R>(service: any, fn: CallFn<A>) => (arg: A) => {
  const [data, setData] = useState<R|null>(null)
  useEffect(() => {
    (async () => {
      setData({} as any)
    })()
  }, [])

  return data
}

const lazyInit = <A, R>(service: any, fn: CallFn<A>) => () => {
  const [data, setData] = useState<R|null>(null)
  const invoke = async (arg: A) => {
    setData({} as any)
  }

  return { data, invoke }
}

export const queriesInit = <A, R>(cmd: any, fn: CallFn<A>) => {
  const queryFn = queryInit<A, R>(cmd, fn)
  const lazyFn = lazyInit<A, R>(cmd, fn)

  return {
    [`useQuery`]: queryFn,
    [`useLazy`]: lazyFn,
  }
}