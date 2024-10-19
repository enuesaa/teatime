import { useGetTeapodInfo } from '@/lib/api/teapods'
import { atom, useAtom } from 'jotai'
import { useEffect } from 'react'

type TeasFilter = {
  teapod: string
  teabox: string
  teaboxes: string[]
}

const state = atom<TeasFilter>({
  teapod: '',
  teabox: '',
  teaboxes: [],
})

export const useTeasFilter = (): [TeasFilter, (newone: Partial<TeasFilter>) => void] => {
  const [value, setValue] = useAtom(state)
  const setter = (newone: Partial<TeasFilter>) => setValue({...value, ...newone})

  return [value, setter]
}

export const useGetTeasFilter = () => {
  const [value, _] = useAtom(state)
  return value
}

export const useInitTeasFilter = (teapod: string, teabox?: string) => {
  const info = useGetTeapodInfo(teapod)
  const [value, setValue] = useTeasFilter()

  useEffect(() => {
    const teaboxes = info.data?.teaboxes.map((v) => v.name) ?? []
    const selected = teabox ?? (teaboxes.length > 0 ? teaboxes[0] : '')
    setValue({ teapod, teabox: selected, teaboxes })
  }, [info.dataUpdatedAt, teabox, teapod])

  return value
}
