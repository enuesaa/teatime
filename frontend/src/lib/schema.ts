
type InfoSchema = {
  name: string
  description: string
  cards: string[]
  panelMap: Record<string, string>
}

type CardSchema = {
  enable: boolean
  layout: string
  textCard: TextCardConfig
}

type TextCardConfig = {
  heading: string
  center: boolean
}

type PanelSchema = {
  enable: boolean
  layout: string
  tablePanel: TablePanelConfig
}

type TablePanelConfig = {
  title: string
  description: string
  head: string[]
  records: TablePanelRecord[]
}

type TablePanelRecord = {
  model: string
  name: string
  values: TablePanelRecordValue[];
}

type TablePanelRecordValue = {
  readonly: boolean
  key: string
}

type RecordSchema = {
  enable: boolean
  name: string
  values: Record<string, RecordValue>
}

type RecordValue = {
  type: string
  strVal: string
  intVal: number
  boolVal: string
}
