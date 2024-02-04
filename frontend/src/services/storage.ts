export function store(key: string, value: string) {
  localStorage.setItem(key, value)
}

export function read(key: string) {
  return localStorage.getItem(key)
}
