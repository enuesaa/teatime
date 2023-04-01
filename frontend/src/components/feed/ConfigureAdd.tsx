import { PageTitle } from '@/components/common/PageTitle'
import { useAddFeedLazy } from '@/lib/feed'
import { AddFeedRequest } from '@/gen/v1/feed_pb'
import { useForm } from 'react-hook-form';
import { TextInput } from '@/components/common/TextInput'

type FormData = {
  name: string;
  url: string;
}
export const ConfigureAdd = () => {
  const { invoke: invokeAddFeed } = useAddFeedLazy()

  const { register, handleSubmit } = useForm<FormData>()
  const handleAddFeed = handleSubmit((data) => {
    invokeAddFeed(data as AddFeedRequest)
  })

  return (
    <section>
      <PageTitle title='Feed Add' />
      <form onSubmit={handleAddFeed}>
        <TextInput label='name' regist={register('name')} />
        <TextInput label='url' regist={register('url')} />
        <button type='submit'>add</button>
      </form>
    </section>
  )
}