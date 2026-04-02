const BaseURL = process.env.BACKEND_URL || 'http://localhost:8080';

type ApiRequestOptions = {
  method: 'GET' | 'POST' | 'PUT' | 'DELETE';
  path: string;
  body?: unknown;
  errorMessage?: string;
};

async function request<T>({ method, path, body, errorMessage }: ApiRequestOptions): Promise<T> {
  const urlPath = path.startsWith('/') ? path : `/${path}`;
  const url = `${BaseURL}${urlPath}`;

  const res = await fetch(url, {
    method,
    headers: { 'Content-Type': 'application/json' },
    body: body ? JSON.stringify(body) : undefined,
  });

  if (!res.ok) {
    console.error(`API ERROR ${res.status} : ${method} ${url}`);
    throw new Error(`${errorMessage || '通信に失敗しました'} (Status: ${res.status})`);
  }

  if (res.status === 204) return {} as T;
  
  return res.json() as Promise<T>;
}

export const api = {
  get: <T>(path: string, error?: string) => 
    request<T>({ method: 'GET', path, errorMessage: error }),

  post: <T>(path: string, body: unknown, error?: string) => 
    request<T>({ method: 'POST', path, body, errorMessage: error }),

  update: <T>(path: string, body: unknown, error?: string) => 
    request<T>({ method: 'PUT', path, body, errorMessage: error }),

  delete: <T>(path: string, error?: string) => 
    request<T>({ method: 'DELETE', path, errorMessage: error }),
};