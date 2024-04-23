import { useNavigate } from 'react-router-dom'
import { useUserByID } from '../hooks/useUserByID'
import { useLikes } from '../hooks/useLikes'
import { getTimeElapsed } from '../services/time'
import { Comment, EmptyHeart, Heart, Options, Save, Share } from './Icons'
import { type PostType } from '../types'
import { Slider } from './Slider'
import './Post.css'

interface PostHeaderProps {
  username: string
  userAvatar: string
  postCreatedAt: string
}

export function PostHeader({ username, userAvatar, postCreatedAt }: PostHeaderProps) {
  const elapsedTime = getTimeElapsed(postCreatedAt)
  return (
    <header className="postHeader">
      <div className="postHeader__user">
        <img className="postHeader__avatar" src={userAvatar} alt={`Avatar image of ${username}`} />
        <strong className="postHeader__username">{username} </strong>
        <span className="postHeader__elapsedTime">• {elapsedTime}</span>
      </div>
      <div className="postHeader__options">
        <Options />
      </div>
    </header>
  )
}

interface PostBodyProps {
  postImages: string[]
  postID: string
  username: string
}

function PostBody({ postImages, postID, username }: PostBodyProps) {
  return (
    <Slider
      id={postID}
      username={username}
      images={postImages}
      forceLimitHeight={false}
    />
  )
}

interface PostFooterProps {
  username: string
  postLikes: number
  postDescription: string
  liked: boolean
  toggleLike: () => void
  commentAction: () => void
}

export function PostFooter({ username, postLikes, postDescription, liked, toggleLike, commentAction }: PostFooterProps) {
  return (
    <footer className="postFooter">
      <section className="postFooter__actions">
        <div className="postFooter__leftActions">
          <button
            className='postFooter__button'
            onClick={toggleLike}
          >
            {liked
              ? <Heart size={24} />
              : <EmptyHeart size={24} />
            }
          </button>
          <button className='postFooter__button' onClick={commentAction}>
            <Comment />
          </button>
          <Share />
        </div>
        <div className="postFooter__rightActions">
          <Save />
        </div>
      </section>
      <section className="postFooter__likeCount">
        <p>{postLikes} {postLikes === 1 ? 'Like' : 'Likes'}</p>
      </section>
      <section className="postFooter__description">
        <p><strong>{username}</strong> {postDescription}</p>
      </section>
    </footer>
  )
}

interface PostProps {
  post: PostType
}

export function Post({ post }: PostProps) {
  const { user } = useUserByID({ userID: post.userID })
  const { toggleLike, liked, likes } = useLikes({ targetID: post.id })
  const navigate = useNavigate()

  const handleToggleLike = () => {
    toggleLike()
  }

  const handleClickComment = () => {
    navigate(`post/${post.id}`)
  }

  return (
    <article className="post">
      <PostHeader
        username={user.username}
        userAvatar={user.avatar}
        postCreatedAt={post.createdAt}
      />
      <PostBody
        postImages={post.images}
        postID={post.id}
        username={user.username}
      />
      <PostFooter
        username={user.username}
        postLikes={likes}
        postDescription={post.description}
        liked={liked}
        toggleLike={handleToggleLike}
        commentAction={handleClickComment}
      />
    </article>
  )
}
