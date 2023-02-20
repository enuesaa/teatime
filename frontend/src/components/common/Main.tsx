import { ReactNode, useEffect } from 'react'
import { css } from '@emotion/react'
import { createConnectTransport, createPromiseClient } from '@bufbuild/connect-web'
import { SettingsService } from '../../../gen/v1/settings_connectweb'

type Props = {
  children: ReactNode
}
export const Main = ({ children }: Props) => {
  const styles = {
    main: css({
      width: '90%',
      margin: '0 auto',
      height: '100vh',
      padding: '20px 0',
    }),
  }

  useEffect(() => {
    (async() => {
      const transport = createConnectTransport({
        baseUrl: '/api/',
      })
      const client = createPromiseClient(SettingsService, transport)
      const res = await client.getAppearance({})
      console.log(res)
    })()
  }, []);

  return <section css={styles.main}>{children}</section>
}
