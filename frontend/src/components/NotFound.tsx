import { Link } from 'react-router-dom'
import './NotFound.css'

export function NotFound() {
  return (
    <div className='notFound'>
      <h1>Not found (404)</h1>
      <p>
        The link you selected may not work or the page may have been removed.
        <Link to="/"> Return to Robotgram</Link>
      </p>
      <img src="" alt="" />
    </div>
  )
}
