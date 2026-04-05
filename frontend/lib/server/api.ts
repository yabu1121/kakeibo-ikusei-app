import { Category, Character } from '@/types/type'
import { cookies } from 'next/headers'

const BASE_URL = process.env.BACKEND_URL || 'http://localhost:8080'

export async function apiFetch(path: string, options?: RequestInit) {
  const cookieStore = await cookies()
  const token = cookieStore.get('token')?.value

  return fetch(`${BASE_URL}${path}`, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...options?.headers,
    },
  })
}


export async function getToken() {
  const cookieStore = await cookies()
  return cookieStore.get('token')?.value
}

export async function getCategories(): Promise<Category[]> {
  const cookieStore = await cookies()
  const token = cookieStore.get('token')?.value
  if (!token) return []

  const res = await fetch(`${BASE_URL}/user/category`, {
    headers: { Authorization: `Bearer ${token}` },
  })
  if (!res.ok) return []
  return res.json()
}

export async function getCharacter(): Promise<Character | null> {
  const cookieStore = await cookies()
  const token = cookieStore.get('token')?.value
  if (!token) return null

  const res = await fetch(`${BASE_URL}/user/character`, {
    headers: { Authorization: `Bearer ${token}` },
  })
  if (!res.ok) return null
  return res.json()
}
