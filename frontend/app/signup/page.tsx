import { signup } from '@/app/actions/auth'
import { SignUpField } from '@/components/features/SignUpField'
import { Title } from '@/components/ui/Title'

export default function SignupPage() {
  return (
    <main>
      <Title>アカウント登録</Title>
      <form action={signup}>
        <SignUpField />
      </form>
      <a href="/login">ログインはこちら</a>
    </main>
  )
}
