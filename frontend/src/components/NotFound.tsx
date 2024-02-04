import { Link } from 'react-router-dom'
import './NotFound.css'

export function NotFound() {
  return (
    <div className='notFoundContainer'>
      <div className='notFound'>
        <h1 className='notFound__title'>404</h1>
        <p className='notFound__text'>
          The link you selected may not work or the page may have been removed.
          <Link to="/"> Return to Robotgram</Link>
        </p>
      </div>
    </div>
  )
}
