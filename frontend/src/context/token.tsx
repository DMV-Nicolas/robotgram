import { createContext, useRef } from 'react'
import { read, store } from '../services/storage'
import { type RefreshTokenResponse, type TokenContextType } from '../types'
import { toast } from 'sonner'

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

  const refreshAccessToken = async () => {
    const res = await fetch('http://localhost:5000/v1/token/refresh', {
      method: 'POST',
      body: JSON.stringify({ refresh_token: refreshToken.current }),
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
        Authorization: `Bearer ${accessToken.current}`
      },
      credentials: 'include'
    })

    if (!res.ok) {
      toast.error('cannot refresh token')
    }

    const data: RefreshTokenResponse = await res.json()
    store('access_token', data.access_token)
  }

  const updateTokens = () => {
    accessToken.current = read('access_token') ?? ''
    refreshToken.current = read('refresh_token') ?? ''
  }

  return (
    <TokenContext.Provider value={{ accessToken, refreshToken, updateTokens, refreshAccessToken }}>
      {children}
    </TokenContext.Provider>
  )
}
