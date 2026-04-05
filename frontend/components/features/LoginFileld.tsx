import { Button } from "../ui/Button"
import { FormLine } from "../ui/FormLine"

export const LoginField = () => {
  return (
    <div>
      <FormLine type="email"/>
      <FormLine type="password"/>
      <Button type="submit">ログイン</Button>
    </div>
  )
}