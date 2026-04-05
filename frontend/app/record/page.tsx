import Link from 'next/link'
import { recordExpense } from '@/app/actions/expense'
import { getCategories } from '@/lib/server/api'
import { Title } from '@/components/ui/Title'

export default async function RecordPage() {
  const categories = await getCategories()

  return (
    <main>
      <Title>支出を記録</Title>
      <form action={recordExpense}>
        <div>
          <label htmlFor="name">支出名</label>
          <input id="name" name="name" required />
        </div>
        <div>
          <label htmlFor="amount">金額</label>
          <input id="amount" name="amount" type="number" min="1" required />
        </div>
        <div>
          <label htmlFor="occured_at">日付</label>
          <input id="occured_at" name="occured_at" type="date" required />
        </div>
        <div>
          <label htmlFor="category_id">カテゴリ</label>
          <select id="category_id" name="category_id" required>
            {categories.map((c) => (
              <option key={c.id} value={c.id}>
                {c.name}
              </option>
            ))}
          </select>
        </div>
        <button type="submit">記録する</button>
      </form>
      <Link href="/">トップへ戻る</Link>
    </main>
  )
}
