import { useLikes } from '../hooks/useLikes'
import { useUserByID } from '../hooks/useUserByID'
import { EmptyHeart, Heart } from './Icons'
import { getTimeElapsed } from '../services/time'
import { type CommentType } from '../types'
import './Reply.css'

interface Props {
  comment: CommentType
  createComment: ({ content }: { content: string }) => Promise<void>
  updateTransform: ({ newTransform }: { newTransform: ({ content }: { content: string }) => Promise<void> }) => void
  focusInput: () => void
  updateInputValue: ({ newValue }: { newValue: string }) => void
}

export function Reply({ comment, createComment, updateTransform, focusInput, updateInputValue }: Props) {
  const { user } = useUserByID({ userID: comment.userID })
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
    <div className='reply'>
      <div className='reply__main'>
        <img className='reply__avatar' src={user.avatar} alt={`Avatar image of ${user.username}`} />
        <div className='reply__text'>
          <p className='reply__content'>
            <strong className='reply__username'>{user.username} </strong>
            {comment.content}
          </p>
          <footer className='reply__footer'>
            <small className='reply__small'>{elapsedTime}</small>
            {likes !== 0 &&
              <small className='reply__small'>
                {likes} {likes === 1 ? 'Like' : 'Likes'}
              </small>
            }
            <button
              className='reply__small comment__small--button'
              onClick={handleReply}>Reply</button>
          </footer>
        </div>
        <button className='reply__likeButton' onClick={handleToggleLike}>
          {liked
            ? <Heart size={12} />
            : <EmptyHeart size={12} />
          }
        </button>
      </div>
    </div >
  )
}
