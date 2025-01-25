import { atom, useAtom } from 'jotai'
import { nanoid } from 'nanoid'

type Toast = {
  id: string
  title: string
  description: string
}
const state = atom<Toast[]>([])

export const useAddToast = () => {
  const [_, setValue] = useAtom(state)
  const hide = useHideToast()

  const setter = ({ title, description }: {title: string, description: string}) => {
    const value = {
      id: nanoid(),
      title,
      description,
    }
    setValue(l => [...l, value])
    setTimeout(() => hide({ id: value.id }), 18000)
  }

  return setter
}

export const useListToast = () => {
  const [value, _] = useAtom(state)
  return value
}

export const useHideToast = () => {
  const [_, setValue] = useAtom(state)
  const setter = ({ id }: {id: string}) => {
    setValue(l => l.filter(i => i.id !== id))
  }

  return setter
}
