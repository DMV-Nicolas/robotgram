import { useId } from 'react'
import { Link } from 'react-router-dom'
import { useLogin } from '../hooks/useLogin'
import { Lock, User } from './Icons'
import './Login.css'

export function Login() {
  const { login } = useLogin()
  const inputUsernameID = useId()
  const inputPasswordID = useId()

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
