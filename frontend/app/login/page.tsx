import { login } from '@/app/actions/auth'

export default function LoginPage() {
  return (
    <main>
      <h1>ログイン</h1>
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
