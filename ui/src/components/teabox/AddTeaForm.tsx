import Form from '@rjsf/core'
import { RJSFSchema, BaseInputTemplateProps } from '@rjsf/utils'
import validator from '@rjsf/validator-ajv8'
import styles from './AddTeaForm.css'

const schema: RJSFSchema = {
  type: 'object',
  properties: {
    name: {
      type: 'string',
    },
    age: {
      type: 'string',
    },
  },
}

export const AddTeaForm = () => {
  return (
    <Form
      schema={schema}
      validator={validator}
      className={styles.form}
      templates={{ ButtonTemplates: { SubmitButton }, BaseInputTemplate }}
    />
  )
}

const SubmitButton = () => {
  return (<></>)
}

const BaseInputTemplate = (props: BaseInputTemplateProps) => {
  return <textarea id={props.id} name={props.name} value={props.value} />
}
