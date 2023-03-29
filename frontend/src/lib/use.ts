import { useState, useEffect } from 'react'
import { createClient } from '@/lib/client'
import { ServiceType } from '@bufbuild/protobuf'

type CallFn<A> = (client: any, arg: A) => Promise<any>;
const queryInit = <A, R>(service: ServiceType, fn: CallFn<A>) => (arg: A) => {
  const [data, setData] = useState<R|null>(null)
  useEffect(() => {
    (async () => {
      const client = createClient(service)
      const res: R = await fn(client, arg)
      setData(res)
    })()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  return data
}

const lazyInit = <A, R>(service: ServiceType, fn: CallFn<A>) => () => {
  const [data, setData] = useState<R|null>(null)
  const invoke = async (arg: A) => {
    const client = createClient(service)
    const res: R = await fn(client, arg)
    setData(res)
  }

  return { data, invoke }
}

export const queriesInit = <A, R>(cmd: ServiceType, fn: CallFn<A>) => {
  const queryFn = queryInit<A, R>(cmd, fn)
  const lazyFn = lazyInit<A, R>(cmd, fn)

  return {
    [`useQuery`]: queryFn,
    [`useLazy`]: lazyFn,
  }
}