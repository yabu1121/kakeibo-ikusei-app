import { cookies } from "next/headers"

export const BASE_URL = process.env.BACKEND_URL || 'http://localhost:8080'

export async function getToken() {
  const cookieStore = await cookies()
  return cookieStore.get('token')?.value
}