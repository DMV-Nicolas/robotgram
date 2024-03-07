import { Link } from 'react-router-dom'
import { useUserByID } from '../hooks/useUserByID'
import { useLikes } from '../hooks/useLikes'
import { getTimeElapsed } from '../services/time'
import { Comment, EmptyHeart, Heart, Options, Save, Share } from './Icons'
import { type PostType } from '../types'
import './Post.css'
import { Slider } from './Slider'

interface PostCardHeaderProps {
  username: string
  userAvatar: string
  postCreatedAt: string
}

function PostHeader({ username, userAvatar, postCreatedAt }: PostCardHeaderProps) {
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

interface PostCardBodyProps {
  postImages: string[]
  postID: string
  username: string
}

function PostBody({ postImages, postID, username }: PostCardBodyProps) {
  return (
    <Slider
      id={postID}
      username={username}
      images={postImages}
    />
  )
}

interface PostCardFooterProps {
  username: string
  postID: string
  postLikes: number
  postDescription: string
  liked: boolean
  toggleLike: () => void
}

function PostFooter({ username, postID, postLikes, postDescription, liked, toggleLike }: PostCardFooterProps) {
  return (
    <footer className="postFooter">
      <section className="postFooter__actions">
        <div className="postFooter__leftActions">
          <button
            className='postFooter__button'
            onClick={toggleLike}
          >
            {liked
              ? <Heart />
              : <EmptyHeart />
            }
          </button>
          <Link
            className='postFooter__button'
            to={`/post/${postID}`}
          >
            <Comment />
          </Link>
          <Share />
        </div>
        <div className="postFooter__rightActions">
          <Save />
        </div>
      </section>
      <section className="postFooter__likeCount">
        <p>{postLikes} Me gusta</p>
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

  const handleToggleLike = () => {
    toggleLike()
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
        postID={post.id}
        postLikes={likes}
        postDescription={post.description}
        liked={liked}
        toggleLike={handleToggleLike}
      />
    </article>
  )
}
