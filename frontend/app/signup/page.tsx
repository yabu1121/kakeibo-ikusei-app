import { signup } from '@/app/actions/auth'
import { SignUpField } from '@/components/features/SignUpField'
import { Title } from '@/components/ui/Title'

export default function SignupPage() {
  return (
    <main className="flex-1 flex items-center justify-center py-8">
      <div className="w-full bg-white rounded-2xl shadow-md p-8 flex flex-col gap-6">
        <Title>アカウント登録</Title>
        <form action={signup} className="flex flex-col gap-4">
          <SignUpField />
        </form>
        <a href="/login" className="text-sm text-center text-gray-400 hover:text-gray-600 transition-colors">
          ログインはこちら
        </a>
      </div>
    </main>
  )
}
