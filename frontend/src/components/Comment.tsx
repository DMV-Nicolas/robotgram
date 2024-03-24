import { useUserByID } from '../hooks/useUserByID'
import { getTimeElapsed } from '../services/time'
import { type CommentType } from '../types'
import './Comment.css'

interface Props {
  comment: CommentType
}

export function Comment({ comment }: Props) {
  const { user } = useUserByID({ userID: comment.userID })
  const elapsedTime = getTimeElapsed(comment.createdAt)

  return (
    <div className='comment'>
      <div className='comment__left'>
        <img className="comment__avatar" src={user.avatar} alt={`Avatar image of ${user.username}`} />
      </div>
      <main className='comment__main'>
        <p className='comment__content'>
          <strong className="comment__username">{user.username} </strong>
          {comment.content}
        </p>
        <footer className='comment__footer'>
          <small className='comment__small'>{elapsedTime}</small>
          <small className='comment__small'>1 Me gusta</small>
          <button className='comment__small comment__small--button'>Responder</button>
        </footer>
      </main>
    </div>
  )
}
