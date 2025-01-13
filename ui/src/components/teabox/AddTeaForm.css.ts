import { style, globalStyle } from '@vanilla-extract/css'

const form = style({})

globalStyle(`${form} fieldset`, {
  border: 'none',
})
globalStyle(`${form} textarea`, {
  display: 'block',
  width: '100%',
  padding: '5px',
  outline: 'none',
  fontSize: '16px',
  minHeight: '20px',
})

export default {
  form,
}
