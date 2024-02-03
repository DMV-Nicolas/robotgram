import { useToken } from '../hooks/useToken'
import './Footer.css'

export function Footer() {
  const { accessToken, refreshToken } = useToken()
  return (
    <footer className='footer'>
      {JSON.stringify({ accessToken, refreshToken })}
    </footer>
  )
}
