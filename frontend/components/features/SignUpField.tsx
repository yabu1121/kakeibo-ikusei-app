import { FormLine } from '../ui/FormLine'
import { Button } from '../ui/Button'

export const SignUpField = () => {
  return (
    <div>
      <FormLine type='name'/>
      <FormLine type='email'/>
      <FormLine type='password'/>
      <Button type='submit'>登録</Button>
    </div>
  )
}
