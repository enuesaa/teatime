import { useListSchemas } from '@/lib/api'

export const Schemas = () => {
  const schemas = useListSchemas()

  if (schemas.isLoading) {
    return (<></>)
  }

  return (
    <>
      {schemas.data?.data?.map((d, i) => (
        <div key={i}>{d.name}</div>
      ))}
    </>
  )
}
