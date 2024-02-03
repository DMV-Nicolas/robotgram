import { useContext } from 'react'
import { TokenContext } from '../context/token'

export function useToken() {
  const context = useContext(TokenContext)

  if (context === null) {
    console.log('Cannot load token context')
    return { accessToken: '', refreshToken: '' }
  }

  const { accessToken, refreshToken } = context
  return { accessToken, refreshToken }
}
