import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useToken } from './useToken'
import { store } from '../services/storage'
import { type UsersLoginResponse } from '../types'

export function useLogin() {
  const navigate = useNavigate()
  const [error, setError] = useState('')
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
      setError('Invalid credentials'); return
    }
    setError('')

    const data: UsersLoginResponse = await res.json()
    store('access_token', data.access_token)
    store('refresh_token', data.refresh_token)
    updateTokens()
    navigate('/')
  }

  return { login, error }
}
