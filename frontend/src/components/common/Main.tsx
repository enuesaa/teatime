import { ReactNode, useEffect } from 'react'
import { css } from '@emotion/react'
import { createConnectTransport, createPromiseClient } from '@bufbuild/connect-web'
import { Setting } from '../../../gen/v1/setting_connectweb'

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
      const client = createPromiseClient(Setting, transport)
      const res = await client.getAppearance({})
      console.log(res.message)
    })()
  }, []);

  return <section css={styles.main}>{children}</section>
}
