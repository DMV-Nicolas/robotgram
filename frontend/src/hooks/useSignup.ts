import { useState } from 'react'
import { useNavigate } from 'react-router-dom'

export function useSignup() {
  const navigate = useNavigate()
  const [error, setError] = useState('')

  const signup = async (username: string, email: string, password: string, gender: string) => {
    const res = await fetch('http://localhost:5000/v1/users', {
      method: 'POST',
      body: JSON.stringify({
        username,
        email,
        password,
        gender,
        full_name: username,
        avatar: 'https://cdn-icons-png.flaticon.com/512/1068/1068549.png'
      }),
      headers: {
        'Content-Type': 'application/json'
      }
    })

    if (!res.ok) {
      setError('Invalid credentials'); return
    }

    navigate('/login')
  }

  return { signup, error }
}
