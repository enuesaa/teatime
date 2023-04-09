import '@emotion/react'
import { CSSObject } from '@emotion/react'

export type SurfDef = {
  main: CSSObject,
  sub: CSSObject,
  reverse: CSSObject,
  transparent: CSSObject
}
export type SizeDef = {
  x1: CSSObject,
  x2: CSSObject,
  x3: CSSObject,
}
export type DecorateDef = {
  a: CSSObject,
  b: CSSObject,
}

declare module '@emotion/react' {
  export interface Theme {
    surf: SurfDef;
    size: SizeDef;
    decorate: DecorateDef;
  }
}