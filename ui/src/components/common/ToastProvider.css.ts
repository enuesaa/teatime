import { style } from '@vanilla-extract/css'

export default {
  viewport: style({
    position: 'fixed',
    bottom: '0',
    right: '0',
  }),
  toast: style({
    display: 'block',
    border: 'solid 1px rgba(255,255,255,0.3)',
    color: '#fafafa',
    padding: '10px 10px',
    margin: '10px',
    borderRadius: '10px',
    width: '300px',
    wordBreak: 'break-all',
    textAlign: 'left',
    background: '#111111',
  }),
}
