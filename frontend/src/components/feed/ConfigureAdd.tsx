import { PageTitle } from '@/components/common/PageTitle'
import { FormEventHandler } from 'react'
import { useAddFeedLazy } from '@/lib/feed'
import { AddFeedRequest } from '@/gen/v1/feed_pb'

export const ConfigureAdd = () => {
  const { invoke: invokeAddFeed } = useAddFeedLazy()

  const handleAddFeed: FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault()
    const formelms = e.currentTarget.elements
    const name = (formelms.namedItem('name') as HTMLInputElement).value ?? ''
    const url = (formelms.namedItem('url') as HTMLInputElement).value ?? ''
    invokeAddFeed({ name, url } as AddFeedRequest)
  }

  return (
    <section>
      <PageTitle title='Feed Add' />
      <form onSubmit={handleAddFeed}>
        <div>
          <label htmlFor='name'>name</label>
          <input type='text' name='name' />
        </div>
        <div>
          <label htmlFor='url'>url</label>
          <input type='text' name='url' />
        </div>
        <button type='submit'>add</button>
      </form>
    </section>
  )
}