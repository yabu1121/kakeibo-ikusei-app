import { CategoryResponse, CreateCategoryRequest } from "@/types/api";
import { BASE_URL, getToken } from "./util";

export async function GETAllCategory(): Promise<CategoryResponse[]> {
  const token = await getToken()

  const res = await fetch(`${BASE_URL}/user/category`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    }
  })
  if (!res.ok) {
    console.error("取得失敗");
    return [];
  }
  return await res.json()
}

export async function CreateCategory(formData: FormData): Promise<CategoryResponse> {
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
  if (!res.ok) {
    const errorMsg = await res.text()
    throw new Error(`カテゴリー作成失敗: ${errorMsg}`)
  }
  return await res.json()
}