import { style, globalStyle } from '@vanilla-extract/css'

const actions = style({
  padding: '5px 0',
})

globalStyle(`${actions} button`, {
  margin: '0 5px',
})

export default {
  actions,
  texts: style({
    width: '100%',
    height: '300px',
  }),
}
