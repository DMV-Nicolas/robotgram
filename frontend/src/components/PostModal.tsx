import { useNavigate } from 'react-router-dom'
import { useLikes } from '../hooks/useLikes'
import { useComments } from '../hooks/useComments'
import { Close } from './Icons'
import { Slider } from './Slider'
import { PostFooter, PostHeader } from './Post'
import { Comment } from './Comment'
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
        forceLimitHeight={true}
      />
    </div>
  )
}

interface PostModalRightProps {
  userID: string
  username: string
  userAvatar: string
  postID: string
  postDescription: string
  postCreatedAt: string
  postLikes: number
  postLiked: boolean
  postToggleLike: () => void
  postComments: CommentType[]
}

function PostModalRight({ userID, username, userAvatar, postID, postDescription, postCreatedAt, postLikes, postLiked, postToggleLike, postComments }: PostModalRightProps) {
  return (
    <div className='postModalRight'>
      <div className='postModalRight__top'>
        <PostHeader
          username={username}
          userAvatar={userAvatar}
          postCreatedAt={postCreatedAt}
        />
        <ul className='postModalRight__comments'>
          <li className='postModalRight__comment'>
            <Comment
              comment={{
                id: '',
                targetID: '',
                userID,
                content: postDescription,
                createdAt: postCreatedAt
              }}
            />
          </li>
          {postComments.map((comment) => (
            <li className='postModalRight__comment' key={comment.id}>
              <Comment comment={comment} />
            </li>
          ))}
        </ul>
      </div>
      <div className='postModalRight__bottom'>
        <PostFooter
          username=""
          postID={postID}
          postLikes={postLikes}
          postDescription=""
          liked={postLiked}
          toggleLike={postToggleLike}
        />
        <form className='postModalRight__form'>
          <input
            className='postModalRight__input'
            type="text"
            placeholder='Add a comment...'
          />
          <button className='postModalRight__button'>Post</button>
        </form>
      </div>
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
              userID={user.id}
              username={user.username}
              userAvatar={user.avatar}
              postID={post.id}
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
