'use server'

import { SlackRequest } from "@/types/api";
import { BASE_URL, getToken } from "./util";

export async function NotifySlack(formData: FormData): Promise<void> {
  const token = await getToken()

  const body: SlackRequest = {
    message: formData.get('message') as string,
  }

  const res = await fetch(`${BASE_URL}/user/slack/notify`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify(body),
  })

  if (!res.ok) return
}
