import { logout } from '@/app/actions/auth'
import { getCharacter } from '@/lib/server/api'
import Image from 'next/image'
import Link from 'next/link'

export default async function TopPage() {
  const character = await getCharacter()
  const frontURL = process.env.NEXT_PUBLIC_FRONTEND_URL || "http://localhost:3000"
  return (
    <main>
      <h1>かけぼん</h1>
      {character ? (
        <div>
          <Image src={`${frontURL}${character.image_url}`} alt={`レベル${character.current_level}`} width={200} height={200}/>
          <p>Lv.{character.current_level}</p>
          <p>
            EXP: {character.current_exp} / {character.exp_to_next_level}
          </p>
          <nav>
            <Link href="/record">支出を記録</Link>
            <form action={logout}>
              <button type="submit">ログアウト</button>
            </form>
          </nav>
        </div>
      ) : (
        <div>
          <p>ログインしてください</p>
          <a href="/login">ログイン</a>
          <a href="/signup">アカウント登録</a>
        </div>
      )}
    </main>
  )
}
