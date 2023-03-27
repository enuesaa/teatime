import { useState, useEffect } from 'react'
import { createClient } from '@/lib/client'
import { ServiceType } from '@bufbuild/protobuf'

type CallFn = (client: any) => Promise<any>;
const queryInit = <R>(service: ServiceType, fn: CallFn) => () => {
  const [data, setData] = useState<R|null>(null)
  useEffect(() => {
    (async () => {
      const client = createClient(service)
      const res: R = await fn(client)
      setData(res)
    })()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  return data
}

const lazyInit = <R>(service: ServiceType, fn: CallFn) => () => {
  const [data, setData] = useState<R|null>(null)
  const invoke = async () => {
    const client = createClient(service)
    const res: R = await fn(client)
    setData(res)
  }

  return { data, invoke }
}

export const queriesInit = <R>(cmd: ServiceType, fn: CallFn) => {
  const queryFn = queryInit<R>(cmd, fn)
  const lazyFn = lazyInit<R>(cmd, fn)

  return {
    [`useQuery`]: queryFn,
    [`useLazy`]: lazyFn,
  }
}