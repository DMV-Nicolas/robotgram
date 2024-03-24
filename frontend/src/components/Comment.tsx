import { useUserByID } from '../hooks/useUserByID'
import { type CommentType } from '../types'
import './Comment.css'

interface Props {
  comment: CommentType
}

export function Comment({ comment }: Props) {
  const { user } = useUserByID({ userID: comment.userID })

  return (
    <div className='comment'>
      <img className="comment__avatar" src={user.avatar} alt={`Avatar image of ${user.username}`} />
      <p className='comment__content'>
        <strong className="comment__username">{user.username} </strong>
        {comment.content}
      </p>
    </div>
  )
}
