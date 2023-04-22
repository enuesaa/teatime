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
    shadow: {
      border: 'none',
      borderRadius: '0',
      boxShadow: '2px 2px 2px rgba(0, 0, 0, 0.7)',
    },
  },
  // hover: {
  //   darken: {
  //     background: 'rgba(0,0,0,0.1)'
  //   }
  // }
  around: {
    x1: {
      padding: '5px',
      margin: '5px',
    },
    x1tb: {
      padding: '5px 0',
      margin: '5px 0',
    },
    x2: {
      padding: '10px',
      margin: '10px',
    },
    x2tb: {
      padding: '10px 0',
      margin: '10px 0',
    },
  },
}
