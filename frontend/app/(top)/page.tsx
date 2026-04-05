import { logout } from '@/app/actions/auth'
import { Title } from '@/components/ui/Title'
import { getCharacter } from '@/lib/server/api'
import Image from 'next/image'
import Link from 'next/link'

export default async function TopPage() {
  const character = await getCharacter()
  const frontURL = process.env.NEXT_PUBLIC_FRONTEND_URL || "http://localhost:3000"
  return (
    <main className="flex-1 flex items-center justify-center py-8">
      <div className="w-full bg-white rounded-2xl shadow-md p-8 flex flex-col gap-6">
        <Title>かけぼん</Title>
        {character ? (
          <div className="flex flex-col items-center gap-4">
            <Image
              src={`${frontURL}${character.image_url}`}
              alt={`レベル${character.current_level}`}
              width={200}
              height={200}
            />
            <p className="text-lg font-bold text-gray-700">Lv.{character.current_level}</p>
            <p className="text-sm text-gray-500">
              EXP: {character.current_exp} / {character.exp_to_next_level}
            </p>
            <nav className="w-full flex flex-col gap-2 pt-2">
              <Link
                href="/record"
                className="w-full text-center bg-blue-600 hover:bg-blue-700 text-white font-medium py-2.5 rounded-lg transition-colors"
              >
                支出を記録
              </Link>
              <form action={logout}>
                <button
                  type="submit"
                  className="w-full text-center bg-gray-100 hover:bg-gray-200 text-gray-600 font-medium py-2.5 rounded-lg transition-colors"
                >
                  ログアウト
                </button>
              </form>
            </nav>
          </div>
        ) : (
          <div className="flex flex-col gap-3">
            <p className="text-sm text-gray-500 text-center">ログインしてください</p>
            <a
              href="/login"
              className="w-full text-center bg-blue-600 hover:bg-blue-700 text-white font-medium py-2.5 rounded-lg transition-colors"
            >
              ログイン
            </a>
            <a
              href="/signup"
              className="w-full text-center bg-gray-100 hover:bg-gray-200 text-gray-600 font-medium py-2.5 rounded-lg transition-colors"
            >
              アカウント登録
            </a>
          </div>
        )}
      </div>
    </main>
  )
}
