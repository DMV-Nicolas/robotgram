export function validateUrl(url: string) {
  return URL.canParse(url)
}
