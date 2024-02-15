import { createContext, useState } from 'react'
import { read, store } from '../services/storage'
import { type RefreshTokenResponse, type TokenContextType } from '../types'
import { toast } from 'sonner'

export const TokenContext = createContext<TokenContextType | null>(null)

export function TokenProvider({ children }: { children?: React.ReactNode }) {
  const [accessToken, setAccessToken] = useState(() => {
    const accessTokenStorage = read('access_token')
    if (accessTokenStorage === null) return ''
    return accessTokenStorage
  })

  const [refreshToken, setRefreshToken] = useState(() => {
    const refreshTokenStorage = read('refresh_token')
    if (refreshTokenStorage === null) return ''
    return refreshTokenStorage
  })

  const updateAccessToken = (newToken: string) => {
    setAccessToken(newToken)
    store('access_token', newToken)
  }

  const updateRefreshToken = (newToken: string) => {
    setRefreshToken(newToken)
    store('refresh_token', newToken)
  }

  const refreshAccessToken = async () => {
    const res = await fetch('http://localhost:5000/v1/token/refresh', {
      method: 'POST',
      body: JSON.stringify({ refresh_token: refreshToken }),
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json'
      }
    })

    if (!res.ok) {
      toast.error('cannot refresh token')
      return Error('cannot refresh token')
    }

    const data: RefreshTokenResponse = await res.json()
    updateAccessToken(data.access_token)
  }

  return (
    <TokenContext.Provider value={{ accessToken, refreshToken, updateAccessToken, updateRefreshToken, refreshAccessToken }}>
      {children}
    </TokenContext.Provider>
  )
}
