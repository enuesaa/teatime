import { createConnectTransport } from '@bufbuild/connect-web'
import { createPromiseClient } from '@bufbuild/connect'

export const createClient = (service: any) => {
  const transport = createConnectTransport({
    baseUrl: '/api/',
  })
  return createPromiseClient(service, transport)
}
