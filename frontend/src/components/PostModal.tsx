import { useNavigate } from 'react-router-dom'
import { useLikes } from '../hooks/useLikes'
import { useComments } from '../hooks/useComments'
import { Close } from './Icons'
import { Slider } from './Slider'
import { type CommentType, type PostType, type UserType } from '../types'
import './PostModal.css'

interface PostModalLeftProps {
  postID: string
  username: string
  postImages: string[]
}

function PostModalLeft({ postID, username, postImages }: PostModalLeftProps) {
  return (
    <div className='postModalLeft'>
      <Slider
        id={postID}
        username={username}
        images={postImages}
      />
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
  postComments: CommentType[]
}

function PostModalRight({ username, userAvatar, postDescription, postCreatedAt, postLikes, postLiked, postToggleLike, postComments }: PostModalRightProps) {
  return (
    <div className='postModalRight'>
      {postComments.map((comment, idx) => (
        <div className='postModalRight__comment' key={`${comment.id}-${idx}`}>
          <p>{comment.content}</p>
        </div>
      ))}
    </div>
  )
}

interface PostModalProps {
  post: PostType
  user: UserType
  loading: boolean
}

export function PostModal({ user, post, loading }: PostModalProps) {
  const { likes, liked, toggleLike } = useLikes({ targetID: post.id })
  const { comments } = useComments({ targetID: post.id })
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
        {!loading &&
          <>
            <PostModalLeft
              postID={post.id}
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
              postComments={comments}
            />
          </>
        }
      </div>
      <button className='postModalContainer__close' onClick={handleGoBack}>
        <Close />
      </button>
    </div>
  )
}
