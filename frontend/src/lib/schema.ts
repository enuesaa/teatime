export type ApiListBase<T> = {
  data: T[]
}
export type ApiBase<T> = {
  data: T
}

export type ProviderSchema = {
  id: string
  name: string
  command: string
}

export type TeaSchema = {
  id: string
  value: string
}

// this is for dev.
export type LinksTeaSchema = {
  title: string
  link: string
}

export type SchemaSchema = {
  name: string
  vals: Record<string, 'str' | 'bool' | 'int'>
}

export type InfoSchema = {
  name: string
  command: string
  description: string
  cards: string[]
  panelMap: Record<string, string>
}

export type CardSchema = {
  enable: boolean
  layout: string
  textCard: TextCardConfig
}

export type TextCardConfig = {
  heading: string
  center: boolean
  text: string
}

export type PanelSchema = {
  enable: boolean
  layout: string
  tablePanel: TablePanelConfig
}

export type TablePanelConfig = {
  title: string
  description: string
  head: string[]
  records: TablePanelRecord[]
}

export type TablePanelRecord = {
  model: string
  name: string
  values: TablePanelRecordValue[]
}

export type TablePanelRecordValue = {
  readonly: boolean
  key: string
}

export type RecordSchema = {
  enable: boolean
  name: string
  values: Record<string, RecordValue>
}

export type RecordValue = {
  type: string
  strVal: string
  intVal: number
  boolVal: string
}
