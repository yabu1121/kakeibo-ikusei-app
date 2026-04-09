import { CategoryResponse, CreateCategoryRequest } from "@/types/api";
import { BASE_URL, getToken } from "./util";

export async function GETAllCategory(): Promise<CategoryResponse[] | undefined> {
  const token = await getToken()

  const res = await fetch(`${BASE_URL}/user/category`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    }
  })
  if (!res.ok) return
  return await res.json()
}

export async function CreateCategory(formData: FormData): Promise<CategoryResponse | undefined> {
  const token = await getToken()

  const body: CreateCategoryRequest = {
    name: formData.get('name') as string,
  }

  const res = await fetch(`${BASE_URL}/user/category`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify(body),
  })
  if (!res.ok) return
  return await res.json()
}