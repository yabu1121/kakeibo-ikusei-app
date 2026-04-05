import Link from 'next/link'
import { recordExpense } from '@/app/actions/expense'
import { getCategories } from '@/lib/server/api'
import { Title } from '@/components/ui/Title'
import { RecordFiled } from '@/components/features/RecordFiled'

export default function RecordPage() {
  return (
    <main className="flex-1 flex items-center justify-center py-8">
      <div className="w-full bg-white rounded-2xl shadow-md p-8 flex flex-col gap-6">
        <Title>支出を記録</Title>
        <form action={recordExpense} className="flex flex-col gap-4">
          <RecordFiled />
        </form>
        <Link href="/" className="text-sm text-center text-gray-400 hover:text-gray-600 transition-colors">
          ← トップへ戻る
        </Link>
      </div>
    </main>
  )
}
