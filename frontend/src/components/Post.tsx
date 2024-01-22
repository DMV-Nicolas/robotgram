import { useState } from 'react'
import { usePost } from '../hooks/usePost'
import { getTimeElapsed } from '../services/time'
import { Comment, EmptyHeart, Heart, Options, Save, Share } from './Icons'
import { type PostType } from '../types'
import './Post.css'

interface PostCardHeaderProps {
  username: string
  userAvatar: string
  postCreatedAt: string
}

function PostHeader({ username, userAvatar, postCreatedAt }: PostCardHeaderProps) {
  const elapsedTime = getTimeElapsed(postCreatedAt)
  return (
    <header className="postHeader">
      <div className="user">
        <img className="avatar" src={userAvatar} alt={`Avatar image of ${username}`} />
        <strong className="username">{username} </strong>
        <span className="elapsedTime">• {elapsedTime}</span>
      </div>
      <div className="options">
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
  const [slide, setSlide] = useState(0)

  const prevSlide = (): void => {
    if (slide > 0) setSlide(slide - 1)
  }

  const nextSlide = (): void => {
    if (slide < postImages.length - 1) setSlide(slide + 1)
  }

  return (
    <div className="postBody">
      {slide > 0 &&
        <span className="leftArrow instagramIcons" onClick={prevSlide}></span>
      }
      <img className="image" src={postImages[slide]} alt={`Post image of ${username}`} />
      {slide < postImages.length - 1 &&
        <span className="rightArrow instagramIcons" onClick={nextSlide}></span>
      }
      <div className="indicators">
        {
          postImages.map((_, idx) => (
            <span key={`${postID}-${idx}`} className={`indicator ${slide === idx ? 'indicatorSelected' : ''}`} onClick={() => { setSlide(idx) }}></span>
          ))
        }
      </div>
    </div>
  )
}

interface PostCardFooterProps {
  username: string
  postLikes: number
  postDescription: string
  liked: boolean
}

function PostFooter({ username, postLikes, postDescription, liked }: PostCardFooterProps) {
  return (
    <footer className="postFooter">
      <section className="actions">
        <div className="leftActions">
          {liked
            ? <Heart />
            : <EmptyHeart />
          }
          <Comment />
          <Share />
        </div>
        <div className="rightActions">
          <Save />
        </div>
      </section>
      <section className="likeCount">
        <p>{postLikes} Me gusta</p>
      </section>
      <section className="description">
        <p><strong>{username}</strong> {postDescription}</p>
      </section>
    </footer>
  )
}

interface PostProps {
  post: PostType
}

export function Post({ post }: PostProps) {
  const { user, likes } = usePost({ post })

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
        liked={false}
      />
    </article>
  )
}
