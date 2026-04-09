	// user.GET("/category", categoryHandler.GetAll)
	// user.POST("/category", categoryHandler.Create)

import { BASE_URL, getToken } from "./util";

export async function GETAllCategory () {
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

export async function CreateCategory (formData: FormData) {
  const token = await getToken()

  const res = await fetch(`${BASE_URL}/user/category`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify({
      name: formData.get('name'),
    })
  })
  if (!res.ok) return
  return await res.json()
}