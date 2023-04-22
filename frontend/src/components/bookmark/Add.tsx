import { PageTitle } from '@/components/common/PageTitle'
import { useStyles } from '@/styles/use'
import { useAddBookmarkLazy } from '@/lib/bookmark'
import { AddBookmarkRequest } from '@/gen/v1/bookmark_pb'
import { useForm } from 'react-hook-form';
import { TextInput } from '@/components/common/TextInput'

type FormData = {
  name: string;
  url: string;
}
export const Add = () => {
  const { invoke: invokeAddBookmark } = useAddBookmarkLazy()
  const styles = useStyles(theme => ({
    form: theme().css({
      margin: '20px',
      'input': { 
        background: 'rgba(255,255,255,0.1)',
        padding: '5px 7px',
        borderRadius: '5px',
        margin: '5px 0 20px 0',
      },
      'button': {
        background: 'rgna(0,0,0,0.1)',
        padding: '5px',
        borderRadius: '5px',
      },
    }),
  }))
  const { register, handleSubmit } = useForm<FormData>()
  const handleAddBookmark = handleSubmit((data) => {
    invokeAddBookmark(data as AddBookmarkRequest)
  })

  return (
    <>
      <PageTitle title='Bookmark Add' />
      <form onSubmit={handleAddBookmark} css={styles.form}>
        <TextInput label='name' regist={register('name')} />
        <TextInput label='url' regist={register('url')} />
        <button type='submit'>add</button>
      </form>
    </>
  )
}