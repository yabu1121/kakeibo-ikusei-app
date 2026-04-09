'use server'

import { redirect } from 'next/navigation'
import { ExpenseRequest, ExpenseResponse } from '@/types/api'
import { BASE_URL, getToken } from './util'


export async function recordExpense(formData: FormData): Promise<void> {
  const token = await getToken()

  const body: ExpenseRequest = {
    name: formData.get('name') as string,
    amount: Number(formData.get('amount')),
    category_id: formData.get('category_id') as string,
    occured_at: new Date(formData.get('occured_at') as string).toISOString(),
  }

  const res = await fetch(`${BASE_URL}/user/expense`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify(body),
  })

  if (!res.ok) return
  redirect('/')
}


export async function GETExpenseById(id: string): Promise<ExpenseResponse | undefined> {
  const token = await getToken()

  const res = await fetch(`${BASE_URL}/user/expense/${id}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    },
  })

  if (!res.ok) return
  return await res.json()
}


export async function UpdateExpense(formData: FormData, id: string): Promise<ExpenseResponse | undefined> {
  const token = await getToken()

  const body: ExpenseRequest = {
    name: formData.get('name') as string,
    amount: Number(formData.get('amount')),
    category_id: formData.get('category_id') as string,
    occured_at: new Date(formData.get('occured_at') as string).toISOString(),
  }

  const res = await fetch(`${BASE_URL}/user/expense/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify(body),
  })
  if (!res.ok) return
  return await res.json()
}

export async function DeleteExpense(id: string): Promise<Response | undefined> {
  const token = await getToken()
  const res = await fetch(`${BASE_URL}/user/expense/${id}`, {
    method: 'DELETE',
    headers: {
      Authorization: `Bearer ${token}`,
    }
  })

  if (!res.ok) return
  return res
}