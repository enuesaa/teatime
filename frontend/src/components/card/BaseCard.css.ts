import { style } from '@vanilla-extract/css'

export default {
  main: style({
    display: 'inline-block',
    width: '300px',
    height: '300px',
    margin: '20px',
    cursor: 'pointer',
  }),
  dialog: style({
    maxWidth: '1000px',
    position: 'relative',
  }),
  dialogClose: style({
    position: 'absolute',
    top: '10px',
    right: '10px',
    cursor: 'pointer',
  }),
}
