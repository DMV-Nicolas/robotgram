import { useLikes } from '../hooks/useLikes'
import { type UserType, type PostType } from '../types'
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

export function PostModal({ post, user }: PostModalProps) {
  const { likes, liked, toggleLike } = useLikes({ targetID: post.id })

  const handleToogleLike = () => {
    toggleLike()
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
      </div>
    </div>
  )
}
