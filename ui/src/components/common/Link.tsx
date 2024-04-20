import { Link as RadixLink, type LinkProps } from '@radix-ui/themes'
import { Link as RouterLink } from 'react-router-dom'

export const Link = ({ href, children, ...props }: LinkProps) => {
  return (
    <RadixLink asChild {...props}>
      <RouterLink to={href ?? '/'}>{children}</RouterLink>
    </RadixLink>
  )
}
