import { Link } from 'react-router-dom'
import './NotFound.css'

export function NotFound() {
  return (
    <div className='notFound'>
      <div className='content'>
        <h1 className='title'>404</h1>
        <p className='text'>
          The link you selected may not work or the page may have been removed.
          <Link to="/"> Return to Robotgram</Link>
        </p>
      </div>
    </div>
  )
}
