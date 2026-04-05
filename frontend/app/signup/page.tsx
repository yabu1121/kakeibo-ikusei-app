import { signup } from '@/app/actions/auth'

export default function SignupPage() {
  return (
    <main>
      <h1>アカウント登録</h1>
      <form action={signup}>
        <div>
          <label htmlFor="name">名前</label>
          <input id="name" name="name" required />
        </div>
        <div>
          <label htmlFor="email">メールアドレス</label>
          <input id="email" name="email" type="email" required />
        </div>
        <div>
          <label htmlFor="password">パスワード</label>
          <input id="password" name="password" type="password" required />
        </div>
        <button type="submit">登録</button>
      </form>
      <a href="/login">ログインはこちら</a>
    </main>
  )
}
