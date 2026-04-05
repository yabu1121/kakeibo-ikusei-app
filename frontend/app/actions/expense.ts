'use server'

import { cookies } from 'next/headers'
import { redirect } from 'next/navigation'

const BASE_URL = process.env.BACKEND_URL || 'http://localhost:8080'

async function getToken() {
  const cookieStore = await cookies()
  return cookieStore.get('token')?.value
}

export async function recordExpense(formData: FormData) {
  const token = await getToken()

  const res = await fetch(`${BASE_URL}/user/expense`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify({
      name: formData.get('name'),
      amount: Number(formData.get('amount')),
      category_id: formData.get('category_id'),
      occured_at: new Date(formData.get('occured_at') as string).toISOString(),
    }),
  })

  if (!res.ok) return
  redirect('/')
}
