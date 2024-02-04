import { useContext } from 'react'
import { TokenContext } from '../context/token'

export function useToken() {
  const context = useContext(TokenContext)

  if (context === null) {
    throw new Error('Cannot load token context')
  }

  const { accessToken, refreshToken, updateTokens } = context
  return { accessToken, refreshToken, updateTokens }
}
