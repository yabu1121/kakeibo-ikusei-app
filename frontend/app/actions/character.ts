import { CharacterResponse } from "@/types/api";
import { BASE_URL, getToken } from "./util";

export async function GetCharacterInformation(): Promise<CharacterResponse | undefined> {
  const token = await getToken()

  const res = await fetch(`${BASE_URL}/user/character`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`
    }
  })
  if (!res.ok) return
  return await res.json()
}

export async function LoginBonus(): Promise<CharacterResponse | undefined> {
  const token = await getToken()

  const res = await fetch(`${BASE_URL}/user/character/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`
    }
  })
  if (!res.ok) return
  return await res.json()
}