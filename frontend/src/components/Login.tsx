import { useId, useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { Lock, User } from './Icons'
import { store } from '../services/storage'
import { type UsersLoginResponse } from '../types'
import { useToken } from '../hooks/useToken'
import './Login.css'

export function Login() {
  const navigate = useNavigate()
  const [error, setError] = useState('')
  const inputUsernameID = useId()
  const inputPasswordID = useId()
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

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()

    const form = e.target as HTMLFormElement
    const formData = new FormData(form)

    const usernameOrEmail = formData.get('usernameOrEmail') as string
    const password = formData.get('password') as string

    login(usernameOrEmail, password)
  }

  return (
    <div className='loginContainer'>
      <div className='login'>
        <h1 className='login__title'>Log In</h1>
        <span className='login_error'>{error}</span>
        <form className='login__form' onSubmit={handleSubmit}>
          <div className='login__inputField'>
            <label className='login__label' htmlFor={inputUsernameID}>
              <User />
            </label>
            <input className='login__input' id={inputUsernameID} name='usernameOrEmail' type="text" placeholder='Username or email' />
          </div>
          <div className='login__inputField'>
            <label className='login__label' htmlFor={inputPasswordID}>
              <Lock />
            </label>
            <input className='login__input' id={inputPasswordID} name='password' type="text" placeholder='Password' />
          </div>
          <button className='login__submit'>Log in</button>
        </form>
        <div className='login__dontHaveAnAccount'>
          <p>{"Don't"} have an account?</p>
          <Link className='notForm' to="/signup"> Sign-up</Link>
        </div>
      </div>
    </div>
  )
}
