import { useGetProviderInfo } from '@/lib/api'
import { css } from '@emotion/react'
import { Code } from '@radix-ui/themes'

type Props = {
  name: string
}
export const PluginInfo = ({ name }: Props) => {
  const {data: info} = useGetProviderInfo(name)

  const styles = {
    main: css({
      display: 'block',
      padding: '10px',
    }),
  }

  return (
    <>
      {info !== undefined && (
        <pre>
          <Code color='gray' css={styles.main}>
            {JSON.stringify(info?.data, null, '  ')}
          </Code>
        </pre>
      )}
    </>
  )
}
