import { createConnectTransport } from '@bufbuild/connect-web'
import { createPromiseClient } from '@bufbuild/connect'
import { ServiceType } from '@bufbuild/protobuf'

export const createClient = (service: ServiceType) => {
  const transport = createConnectTransport({
    baseUrl: '/api/',
  })
  return createPromiseClient(service, transport)
}
