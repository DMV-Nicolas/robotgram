import { useId, useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { Lock, User } from './Icons'
import { store } from '../services/storage'
import { type UsersLoginResponse } from '../types'
import { useToken } from '../hooks/useToken'
import './Login.css'

export function Login() {
  const { updateTokens } = useToken()
  const [error, setError] = useState('')
  const navigate = useNavigate()
  const inputUsernameID = useId()
  const inputPasswordID = useId()

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

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()

    const form = e.target as HTMLFormElement
    const formData = new FormData(form)

    const usernameOrEmail = formData.get('usernameOrEmail') as string
    const password = formData.get('password') as string

    login(usernameOrEmail, password)
  }

  return (
    <div className='container'>
      <div className='login'>
        <h1 className='title'>Log In</h1>
        <span style={{ color: 'red' }}>{error}</span>
        <form className='form' onSubmit={handleSubmit}>
          <div className='inputField'>
            <label htmlFor={inputUsernameID}>
              <User />
            </label>
            <input id={inputUsernameID} name='usernameOrEmail' type="text" placeholder='Username or email' />
          </div>
          <div className='inputField'>
            <label htmlFor={inputPasswordID}>
              <Lock />
            </label>
            <input id={inputPasswordID} name='password' type="text" placeholder='Password' />
          </div>
          <button className='submit'>Log in</button>
        </form>
        <div className='notForm'>
          <p>{"Don't"} have an account?</p>
          <Link className='notForm' to="/signup"> Sign-up</Link>
        </div>
      </div>
    </div>
  )
}
