'use server'

import { cookies } from 'next/headers'
import { redirect } from 'next/navigation'
import { LoginRequest, LoginResponse, SignUpRequest } from '@/types/api'
import { BASE_URL } from './util'

export async function login(formData: FormData): Promise<void> {
  const body: LoginRequest = {
    email: formData.get('email') as string,
    password: formData.get('password') as string,
  }

  const res = await fetch(`${BASE_URL}/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body),
  })

  if (!res.ok) return

  const { token }: LoginResponse = await res.json()
  const cookieStore = await cookies()
  cookieStore.set('token', token, { httpOnly: true, path: '/' })
  redirect('/')
}

export async function signup(formData: FormData): Promise<void> {
  const body: SignUpRequest = {
    name: formData.get('name') as string,
    email: formData.get('email') as string,
    password: formData.get('password') as string,
  }

  const res = await fetch(`${BASE_URL}/signup`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body),
  })

  if (!res.ok) return
  redirect('/login')
}

export async function logout() {
  const cookieStore = await cookies()
  cookieStore.delete('token')
  redirect('/login')
}
