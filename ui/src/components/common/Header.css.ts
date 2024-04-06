import { style, globalStyle } from '@vanilla-extract/css'

const main = style({
  fontSize: '25px',
  fontWeight: 'bold',
})
globalStyle(`${main} a`, {
  textDecoration: 'none',
})

const heading = style({
  color: 'white',
  margin: '0 10px',
})
globalStyle(`${heading} svg`, {
  verticalAlign: 'middle',
  margin: '0 10px',
})

const setting = style({
  color: 'white',
})
globalStyle(`${setting} svg`, {
  verticalAlign: 'middle',
})

export default {
  main,
  heading,
  setting,
}
