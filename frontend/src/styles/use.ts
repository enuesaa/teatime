import { css, useTheme, type CSSObject, SerializedStyles, Theme } from '@emotion/react'
import { SurfDef, SizeDef, DecorateDef } from './emotion'

type ThemedStyleArgs = {
  surf: keyof SurfDef;
  size: keyof SizeDef;
  decorate: keyof DecorateDef;
}
class ThemedStyle {
  public cssobject: CSSObject = {};

  constructor(
    theme: Theme,
    {surf, size, decorate}: Partial<ThemedStyleArgs>
  ) {
    this.cssobject = {
      ...(surf === undefined ? {} : theme.surf[surf]),
      ...(size === undefined ? {} : theme.size[size]),
      ...(decorate === undefined ? {} : theme.decorate[decorate]),
    }
  }

  css(css: CSSObject) {
    this.cssobject = { ...this.cssobject, ...css }
    return this
  }
}

export const useStyles = <A extends string>(
  createStyles: (theme: (args: Partial<ThemedStyleArgs>) => ThemedStyle) => Record<A, ThemedStyle>
): Record<A, SerializedStyles> => {
  const theme = useTheme()
  const styles = createStyles((args) => new ThemedStyle(theme, args))

  return Object.fromEntries(
    Object.entries<ThemedStyle>(styles).map(([k, v]) => [k, css(v.cssobject)])
  ) as Record<A, SerializedStyles>
}
