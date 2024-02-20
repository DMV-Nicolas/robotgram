import { useLikes } from '../hooks/useLikes'
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

  const handleToogleLike = () => {
    toggleLike()
  }

  return (
    <div className="postModalContainer">
      <div className="postModal">
        <h1>{likes}</h1>
        <h1>{liked ? 'liked' : 'not liked'}</h1>
        <h1>{user.username}</h1>
        <h1>{post.id}</h1>
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
