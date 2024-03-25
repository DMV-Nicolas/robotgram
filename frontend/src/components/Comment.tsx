import { useLikes } from '../hooks/useLikes'
import { useUserByID } from '../hooks/useUserByID'
import { useComments } from '../hooks/useComments'
import { getTimeElapsed } from '../services/time'
import { EmptyHeart, Heart } from './Icons'
import { type CommentType } from '../types'
import './Comment.css'

interface Props {
  comment: CommentType
  withLike: boolean
  updateTransform: ({ newTransform }: { newTransform: ({ content }: { content: string }) => Promise<void> }) => void
  focusInput: () => void
  updateInputValue: ({ newValue }: { newValue: string }) => void
}

export function Comment({ comment, withLike, updateTransform, focusInput, updateInputValue }: Props) {
  const { user } = useUserByID({ userID: comment.userID })
  const { comments, createComment } = useComments({ targetID: comment.id })
  const { toggleLike, liked, likes } = useLikes({ targetID: comment.id })
  const elapsedTime = getTimeElapsed(comment.createdAt)

  const handleToggleLike = () => {
    toggleLike()
  }

  const handleReply = () => {
    updateTransform({ newTransform: createComment })
    focusInput()
    updateInputValue({ newValue: `@${user.username} ` })
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
          <button
            className='comment__small comment__small--button'
            onClick={handleReply}>Reply</button>
        </footer>
      </main>
      {withLike &&
        <button
          className='comment__likeButton'
          onClick={handleToggleLike}
        >
          {liked
            ? <Heart size={12} />
            : <EmptyHeart size={12} />
          }
        </button>
      }
      <aside>
        {comments.map((comment) => (
          <small key={comment.id}>{comment.content}</small>
        ))}
      </aside>
    </div>
  )
}
