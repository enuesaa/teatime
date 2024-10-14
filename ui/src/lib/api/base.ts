import { query } from './base-query'
import { mutate, type MutateConfig } from './base-mutate'

export const queryGet = <T>(path: string) => query<T>('GET', path, {})

export const mutatePost = <R, T>(path: string, config: MutateConfig) => mutate<R, T>('POST', path, config)
