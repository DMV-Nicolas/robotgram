import './Login.css'

export function Login() {
  return (
    <div className='login'>
      <div className='left'>
        <img className='picture' src="https://images.wondershare.com/filmora/article-images/2021/best-practices-for-creating-phone-aspect-ratio-vertical-on-your-smartphone8.jpg" />
      </div>
      <div className='right'>
        <form className='form'>
          <h1 className='title'>Robotgram</h1>
          <div>
            <input className='input' type="text" placeholder='Username or email' />
            <input className='input' type="text" placeholder='Password' />
          </div>
          <button className='submit'>Log in</button>
        </form>
        <div className='notForm'>
          <p>You do not have an account?</p>
          <a href="/signup">Sign up</a>
        </div>
        <div className='promotion'>

        </div>
      </div>
    </div>
  )
}
