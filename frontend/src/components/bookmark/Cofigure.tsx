import { PageTitle } from '@/components/common/PageTitle'
import { css, useTheme } from '@emotion/react'
import { useListBookmarksQuery, useAddBookmarkLazy } from '@/lib/bookmark'
import Link from 'next/link'
import { FormEventHandler } from 'react'
import { AddBookmarkRequest } from '@/gen/v1/bookmark_pb'

export const Configure = () => {
  const theme = useTheme()

  const bookmarks = useListBookmarksQuery({})
  const { invoke: invokeAddBookmark } = useAddBookmarkLazy()
  const styles = {
    item: css({
      padding: '10px',
      margin: '10px',
      borderRadius: '10px',
      border: 'solid 1px rgba(255,255,255,0.2)',
    }),
    form: css({
      margin: '20px',
      'input': { 
        ...theme.input,
        background: 'rgba(255,255,255,0.1)',
        padding: '5px 7px',
        borderRadius: '5px',
        color: theme.color.main,
        margin: '5px 0 20px 0',
        fontSize: theme.fontSize.large,
      },
      'button': {
        ...theme.input,
        background: 'rgna(0,0,0,0.1)',
        padding: '5px',
        borderRadius: '5px',
      },
    }),
  }

  const handleAddBookmark: FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault()
    const formelms = e.currentTarget.elements
    const name = (formelms.namedItem('name') as HTMLInputElement).value ?? ''
    const url = (formelms.namedItem('url') as HTMLInputElement).value ?? ''
    invokeAddBookmark({ name, url } as AddBookmarkRequest)
  }

  return (
    <>
      <PageTitle title='Bookmark' />
      {bookmarks?.items.map((v,i) => (
        <Link key={i} href={v.url} css={styles.item} target='_blank'>
          {v?.name}
        </Link>
      ))}
      <form onSubmit={handleAddBookmark} css={styles.form}>
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
    </>
  )
}