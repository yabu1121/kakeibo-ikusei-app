import { login } from '@/app/actions/auth'
import { LoginField } from '@/components/features/LoginFileld'
import { Title } from '@/components/ui/Title'

export default function LoginPage() {
  return (
    <main>
      <Title>ログイン</Title>
      <form action={login}>
        <LoginField />
      </form>
      <a href="/signup">アカウント登録はこちら</a>
    </main>
  )
}
