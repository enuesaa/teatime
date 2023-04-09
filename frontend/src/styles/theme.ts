import { type Theme } from '@emotion/react'

export const baseTheme: Theme = {
  surf: {
    main: {
      color: '#fafafa',
      background: '#222222',
    },
    sub: {
      color: '#fafafa',
      background: '#333333',
    },
    reverse: {
      color: '#111111',
      background: '#fafafa',
    },
    transparent: {
      color: 'none',
      background: 'transparent',
    },
  },
  size: {
    x1: {
      fontSize: '16px',
      fontWeight: '400',
      // padding: '2px 5px',
    },
    x2: {
      fontSize: '20px',
      fontWeight: '600',
      // padding: '5px 10px',
    },
    x3: {
      fontSize: '25px',
      fontWeight: '800',
      // padding: '10px 15px',
    },
  },
  decorate: {
    a: {
      border: 'none',
      borderRadius: '0',
      boxShadow: 'none',
    },
    b: {
      border: 'none',
      borderRadius: '10px',
      boxShadow: 'none',
    },
  }
}
