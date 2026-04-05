export type Character = {
  id: string
  user_id: string
  current_level: number
  current_exp: number
  exp_to_next_level: number
  image_url: string
}

export type Expense = {
  id: string
  name: string
  amount: number
  occured_at: string
  category_id: string
  user_id: string
}

export type Category = {
  id: string
  name: string
}

export type User = {
  id: string
  name: string
  email: string
  role: string
}
