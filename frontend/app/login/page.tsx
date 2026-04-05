import { login } from '@/app/actions/auth'
import { LoginField } from '@/components/features/LoginFileld'
import { Title } from '@/components/ui/Title'

export default function LoginPage() {
  return (
    <main className="flex-1 flex items-center justify-center py-8">
      <div className="w-full bg-white rounded-2xl shadow-md p-8 flex flex-col gap-6">
        <Title>ログイン</Title>
        <form action={login} className="flex flex-col gap-4">
          <LoginField />
        </form>
        <a href="/signup" className="text-sm text-center text-gray-400 hover:text-gray-600 transition-colors">
          アカウント登録はこちら
        </a>
      </div>
    </main>
  )
}
