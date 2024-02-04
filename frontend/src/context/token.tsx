import { createContext, useRef } from 'react'
import { read } from '../services/storage'
import { type TokenContextType } from '../types'

export const TokenContext = createContext<TokenContextType | null>(null)

export function TokenProvider({ children }: { children?: React.ReactNode }) {
  const accessToken = useRef((() => {
    const accessTokenStorage = read('access_token')
    if (accessTokenStorage === null) return ''
    return accessTokenStorage
  })())

  const refreshToken = useRef((() => {
    const refreshTokenStorage = read('refresh_token')
    if (refreshTokenStorage === null) return ''
    return refreshTokenStorage
  })())

  const updateTokens = () => {
    accessToken.current = read('access_token') ?? ''
    refreshToken.current = read('refresh_token') ?? ''
  }

  return (
    <TokenContext.Provider value={{ accessToken, refreshToken, updateTokens }}>
      {children}
    </TokenContext.Provider>
  )
}
