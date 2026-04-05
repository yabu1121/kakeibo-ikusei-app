'use server'

import { cookies } from 'next/headers'
import { redirect } from 'next/navigation'

const BASE_URL = process.env.BACKEND_URL || 'http://localhost:8080'

export async function login(formData: FormData) {
  const res = await fetch(`${BASE_URL}/user/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      email: formData.get('email'),
      password: formData.get('password'),
    }),
  })

  if (!res.ok) return

  const { token } = await res.json()
  const cookieStore = await cookies()
  cookieStore.set('token', token, { httpOnly: true, path: '/' })
  redirect('/')
}

export async function signup(formData: FormData) {
  const res = await fetch(`${BASE_URL}/user/signup`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      name: formData.get('name'),
      email: formData.get('email'),
      password: formData.get('password'),
    }),
  })

  if (!res.ok) return
  redirect('/login')
}

export async function logout() {
  const cookieStore = await cookies()
  cookieStore.delete('token')
  redirect('/login')
}
