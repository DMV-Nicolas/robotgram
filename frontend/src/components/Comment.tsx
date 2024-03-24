import { useLikes } from '../hooks/useLikes'
import { useUserByID } from '../hooks/useUserByID'
import { getTimeElapsed } from '../services/time'
import { EmptyHeart, Heart } from './Icons'
import { type CommentType } from '../types'
import './Comment.css'

interface Props {
  comment: CommentType
  withLike: boolean
}

export function Comment({ comment, withLike }: Props) {
  const { user } = useUserByID({ userID: comment.userID })
  const { toggleLike, liked, likes } = useLikes({ targetID: comment.id })
  const elapsedTime = getTimeElapsed(comment.createdAt)

  const handleClick = () => {
    toggleLike()
  }

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
          {likes !== 0 &&
            <small className='comment__small'>
              {likes} {likes === 1 ? 'Like' : 'Likes'}
            </small>
          }
          <button className='comment__small comment__small--button'>Responder</button>
        </footer>
      </main>
      {withLike &&
        <button
          className='comment__likeButton'
          onClick={handleClick}
        >
          {liked
            ? <Heart size={12} />
            : <EmptyHeart size={12} />
          }
        </button>
      }
    </div>
  )
}
