import { useNavigate } from 'react-router-dom'
import { Close } from './Icons'
import './CreatePost.css'

export function CreatePost() {
  const navigate = useNavigate()

  const handleGoBack = () => {
    navigate(-1)
  }

  return (
    <div className="createPostContainer">
      <div className="createPost">
        <h2>Create a new post!</h2>
      </div>
      <button className='postModalContainer__close' onClick={handleGoBack}>
        <Close />
      </button>
    </div>
  )
}
