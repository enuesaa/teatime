import { type Theme } from '@emotion/react'

export const baseTheme: Theme = {
  input: {
    outline: 'none',
    appearance: 'none',
    border: 'none',
    display: 'block',
  },
  button: {
    outline: 'none',
    appearance: 'none',
    border: 'none',
    display: 'block',
  },
  box: {
    padding: '10px',
    width: '100%',
  },
  outerbox: {
    padding: '10px',
    margin: '10px',
  },
  linkCard: {
    display: 'inline-block',
    padding: '10px 20px',
    color: '#cccccc',
    fontSize: '20px',
    background: '#333333',
    margin: '10px',
    borderRadius: '10px',
  },
  heading: {
    width: '100%',
    display: 'block',
    fontWeight: '600',
    padding: '10px',
    fontSize: '30px',
    userSelect: 'none',
  },

  color: {
    main: '#e6e6e6',
    sub: 'rgba(0, 0, 0, 0.15)',
    subHover: 'rgba(0, 0, 0, 0.25)',
    contrast: '#1a1a1a',
    highlight: '#ff6633',
    highlightHover: '#ff9933',
    danger: '#ff1111',
  },
  fontSize: {
    small: '14px',
    normal: '16px',
    large: '20px',
    xlarge: '25px',
    x2large: '30px',
    x3large: '35px',
  },
}
