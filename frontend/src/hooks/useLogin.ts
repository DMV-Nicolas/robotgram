import { useNavigate } from 'react-router-dom'
import { useToken } from './useToken'
import { store } from '../services/storage'
import { type UsersLoginResponse } from '../types'
import { toast } from 'sonner'

export function useLogin() {
  const navigate = useNavigate()
  const { updateTokens } = useToken()

  const login = async (usernameOrEmail: string, password: string) => {
    const res = await fetch('http://localhost:5000/v1/users/login', {
      method: 'POST',
      body: JSON.stringify({ username_or_email: usernameOrEmail, password }),
      headers: {
        'Content-Type': 'application/json'
      }
    })

    if (!res.ok) {
      toast.error('Invalid credentials')
      return
    }

    const data: UsersLoginResponse = await res.json()
    store('access_token', data.access_token)
    store('refresh_token', data.refresh_token)
    updateTokens()

    toast.success('Successful login')
    navigate('/')
  }

  return { login }
}
