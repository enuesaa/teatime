import { PageTitle } from '@/components/common/PageTitle'
import { useAddFeedLazy } from '@/lib/feed'
import { AddFeedRequest } from '@/gen/v1/feed_pb'
import { useForm } from 'react-hook-form';
import { TextInput } from '@/components/common/TextInput'
import { useStyles } from '@/styles/use';

type FormData = {
  name: string;
  url: string;
}
export const Add = () => {
  const { invoke: invokeAddFeed } = useAddFeedLazy()
  const styles = useStyles((theme) => ({
    form: theme({ around: 'x1' }).css({
      input: theme({ surf: 'sub', size: 'x1', around: 'x1' }).to(),
      button: theme({ surf: 'reverse', size: 'x1', decorate: 'shadow', around: 'x2' })
        .css({
          cursor: 'pointer',
        })
        .to(),
    }),
  }))

  const { register, handleSubmit } = useForm<FormData>()
  const handleAddFeed = handleSubmit((data) => {
    invokeAddFeed(data as AddFeedRequest)
  })

  return (
    <section>
      <PageTitle title='Feed Add' />
      <form onSubmit={handleAddFeed} css={styles.form}>
        <TextInput label='name' regist={register('name')} />
        <TextInput label='url' regist={register('url')} />
        <button type='submit'>add</button>
      </form>
    </section>
  )
}