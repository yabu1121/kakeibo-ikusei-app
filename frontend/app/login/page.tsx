import { login } from '@/app/actions/auth'
import { Title } from '@/components/ui/Title'

export default function LoginPage() {
  return (
    <main>
      <Title>ログイン</Title>
      <form action={login}>
        <div>
          <label htmlFor="email">メールアドレス</label>
          <input id="email" name="email" type="email" required />
        </div>
        <div>
          <label htmlFor="password">パスワード</label>
          <input id="password" name="password" type="password" required />
        </div>
        <button type="submit">ログイン</button>
      </form>
      <a href="/signup">アカウント登録はこちら</a>
    </main>
  )
}
