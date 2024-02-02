import { useId } from 'react'
import './Login.css'
import { Lock, User } from './Icons'
import { Link } from 'react-router-dom'

export function Login() {
  const inputUsernameID = useId()
  const inputPasswordID = useId()

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()

    const form = e.target as HTMLFormElement
    const formData = new FormData(form)

    const username = formData.get('username')
    const password = formData.get('password')
  }

  return (
    <div className='container'>
      <div className='login'>
        <form className='form' onSubmit={handleSubmit}>
          <h1 className='title'>Log In</h1>
          <div className='inputField'>
            <label htmlFor={inputUsernameID}>
              <User />
            </label>
            <input id={inputUsernameID} name='username' type="text" placeholder='Username or email' />
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
