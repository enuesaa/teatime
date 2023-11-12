import { useGetProviderInfo } from '@/lib/api'
import { css } from '@emotion/react'
import { Code } from '@radix-ui/themes'

type Props = {
  id: string
}
export const PluginInfo = ({ id }: Props) => {
  const {data: info} = useGetProviderInfo(id)

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
