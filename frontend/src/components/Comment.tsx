import { useLikes } from '../hooks/useLikes'
import { useUserByID } from '../hooks/useUserByID'
import { useComments } from '../hooks/useComments'
import { EmptyHeart, Heart } from './Icons'
import { Reply } from './Reply'
import { getTimeElapsed } from '../services/time'
import { type CommentType } from '../types'
import './Comment.css'

interface Props {
  comment: CommentType
  isDescriptionStyle: boolean
  isReplyStyle: boolean
  updateTransform: ({ newTransform }: { newTransform: ({ content }: { content: string }) => Promise<void> }) => void
  focusInput: () => void
  updateInputValue: ({ newValue }: { newValue: string }) => void
}

export function Comment({ comment, isDescriptionStyle, isReplyStyle, updateTransform, focusInput, updateInputValue }: Props) {
  const { user } = useUserByID({ userID: comment.userID })
  const { likes, liked, toggleLike } = useLikes({ targetID: comment.id })
  const { comments, createComment } = useComments({ targetID: comment.id })
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
      <div className='comment__main'>
        <img className="comment__avatar" src={user.avatar} alt={`Avatar image of ${user.username}`} />
        <div className='comment__text'>
          <p className='comment__content'>
            <strong className="comment__username">{user.username} </strong>
            {comment.content}
          </p>
          {!isDescriptionStyle &&
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
          }
        </div>
        {!isDescriptionStyle &&
          <button className='comment__likeButton' onClick={handleToggleLike}>
            {liked
              ? <Heart size={12} />
              : <EmptyHeart size={12} />
            }
          </button>
        }
      </div>
      {!isDescriptionStyle &&
        <ul className='comment__replys' style={{ paddingLeft: isReplyStyle ? '0' : '42px' }} >
          {
            comments.map((comment) => (
              <li className='comment__reply' key={comment.id}>
                <Reply
                  comment={comment}
                  createComment={createComment}
                  focusInput={focusInput}
                  updateInputValue={updateInputValue}
                  updateTransform={updateTransform}
                />
              </li>
            ))
          }
        </ul>
      }
    </div >
  )
}
