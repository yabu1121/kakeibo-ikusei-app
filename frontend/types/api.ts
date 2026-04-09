// ---- Auth ----

export type LoginRequest = {
  email: string
  password: string
}

export type LoginResponse = {
  token: string
}

export type SignUpRequest = {
  name: string
  email: string
  password: string
}

// ---- Expense ----

export type ExpenseRequest = {
  name: string
  amount: number
  category_id: string
  occured_at: string // ISO 8601
}

export type ExpenseResponse = {
  id: string
  name: string
  amount: number
  occured_at: string
  category_id: string
  user_id: string
}

// ---- Character ----

export type CharacterResponse = {
  id: string
  user_id: string
  current_level: number
  current_exp: number
  exp_to_next_level: number
  image_url: string
}

// ---- Category ----

export type CreateCategoryRequest = {
  name: string
}

export type CategoryResponse = {
  id: string
  name: string
}
