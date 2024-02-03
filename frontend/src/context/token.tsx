import { createContext } from 'react'
import { read } from '../services/storage'
import { type TokenContextType } from '../types'

export const TokenContext = createContext<TokenContextType | null>(null)

export function TokenProvider({ children }: { children?: React.ReactNode }) {
  const accessToken = (() => {
    const accessTokenStorage = read('access_token')
    if (accessTokenStorage === null) return ''
    return accessTokenStorage
  })()

  const refreshToken = (() => {
    const refreshTokenStorage = read('refresh_token')
    if (refreshTokenStorage === null) return ''
    return refreshTokenStorage
  })()

  return (
    <TokenContext.Provider value={{ accessToken, refreshToken }}>
      {children}
    </TokenContext.Provider>
  )
}
