import { useNavigate } from 'react-router-dom'
import { useLikes } from '../hooks/useLikes'
import { Close } from './Icons'
import { type PostType, type UserType } from '../types'
import './PostModal.css'

interface PostModalLeftProps {
  postImages: string[]
  username: string
}

function PostModalLeft({ postImages, username }: PostModalLeftProps) {
  return (
    <div className='postModalLeft'>
    </div>
  )
}

interface PostModalRightProps {
  username: string
  userAvatar: string
  postDescription: string
  postCreatedAt: string
  postLikes: number
  postLiked: boolean
  postToggleLike: () => void
}

function PostModalRight({ username, userAvatar, postDescription, postCreatedAt, postLikes, postLiked, postToggleLike }: PostModalRightProps) {
  return (
    <div className='postModalRight'>
    </div>
  )
}

interface PostModalProps {
  post: PostType
  user: UserType
}

export function PostModal({ user, post }: PostModalProps) {
  const { likes, liked, toggleLike } = useLikes({ targetID: post.id })
  const navigate = useNavigate()

  const handleToogleLike = () => {
    toggleLike()
  }

  const handleGoBack = () => {
    navigate(-1)
  }

  return (
    <div className="postModalContainer">
      <div className="postModal">
        <PostModalLeft
          postImages={post.images}
          username={user.username}
        />
        <PostModalRight
          username={user.username}
          userAvatar={user.avatar}
          postDescription={post.description}
          postCreatedAt={post.createdAt}
          postLikes={likes}
          postLiked={liked}
          postToggleLike={handleToogleLike}
        />
        <button className='postModal__close' onClick={handleGoBack}>
          <Close />
        </button>
      </div>
    </div>
  )
}
