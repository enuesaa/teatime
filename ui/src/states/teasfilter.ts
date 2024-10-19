import { useGetTeapodInfo } from '@/lib/api/teapods'
import { atom, useAtom } from 'jotai'
import { useEffect } from 'react'

type TeasFilter = {
  teapod: string
  teabox: string
  teaboxes: string[]
}
const state = atom<Partial<TeasFilter>>({
  teapod: undefined,
  teabox: undefined,
  teaboxes: [],
})

export const useTeasFilter = (): [TeasFilter, (newone: Partial<TeasFilter>) => void] => {
  const [value, setValue] = useAtom(state)
  const setter = (newone: Partial<TeasFilter>) => setValue({...value, ...newone})
  const data = {
    teapod: value.teapod ?? '',
    teabox: value.teabox ?? '',
    teaboxes: value.teaboxes ?? [],
  }
  return [data, setter]
}

export const useGetTeasFilter = () => {
  const [data, _] = useTeasFilter()
  return data
}

export const useInitTeasFilter = (teapod: string, teabox?: string) => {
  const info = useGetTeapodInfo(teapod)
  const [data, setData] = useTeasFilter()
  const ok = data.teabox !== ''

  useEffect(() => {
    const teaboxes = info.data?.teaboxes.map((v) => v.name) ?? []
    const selected = teabox ?? (teaboxes.length > 0 ? teaboxes[0] : '')
    setData({ teapod, teabox: selected, teaboxes })
  }, [info.dataUpdatedAt, teabox, teapod])

  return {...data, ok}
}
