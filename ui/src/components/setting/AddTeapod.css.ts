import { style, globalStyle } from '@vanilla-extract/css'

const actions = style({
  padding: '5px 0',
})

globalStyle(`${actions} button`, {
  margin: '0 5px',
})

export default {
  trigger: style({
    margin: '0 5px',
    padding: '10px 10px',
    borderRadius: '10px',
    verticalAlign: 'middle',
    background: '#ff9933'
  }),
  actions,
}
