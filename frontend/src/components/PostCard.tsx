import { Comment, EmptyHeart, Heart, Options, Share } from "./Icons"
import { Post, User } from "../types"
import { getTimeElapsed } from "../services/time"
import "./PostCard.css"
import { useState } from "react"

type PostCardHeaderParams = {
    username: string
    userAvatar: string
    postCreatedAt: string
}

function PostCardHeader({ username, userAvatar, postCreatedAt }: PostCardHeaderParams) {
    const elapsedTime = getTimeElapsed(postCreatedAt)
    return (
        <header className="postCardHeader">
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

type PostCardBodyParams = {
    postImages: string[]
    postID: string,
    username: string
}

function PostCardBody({ postImages, postID, username }: PostCardBodyParams) {
    const [slide, setSlide] = useState(0)

    const prevSlide = () => {
        if (slide > 0) setSlide(slide - 1)
    }

    const nextSlide = () => {
        if (slide < postImages.length - 1) setSlide(slide + 1)
    }

    return (
        <div className="postCardBody">
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
                        <span key={`${postID}-${idx}`} className={`indicator ${slide === idx ? "indicatorSelected" : ""}`} onClick={() => setSlide(idx)}></span>
                    ))
                }
            </div>
        </div>
    )
}

type PostCardFooterParams = {
    username: string
    postLikes: number
    postDescription: string
    liked: boolean
}

function PostCardFooter({ username, postLikes, postDescription, liked }: PostCardFooterParams) {
    return (
        <footer className="postCardFooter">
            <section className="actions">
                {liked
                    ? <Heart />
                    : <EmptyHeart />
                }
                <Comment />
                <Share />
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

type PostCardParams = {
    post: Post
    user: User
    likes: number
}

export function PostCard({ post, user, likes }: PostCardParams) {
    return (
        <article className="postCard">
            <PostCardHeader
                username={user.username}
                userAvatar={user.avatar}
                postCreatedAt={post.createdAt}
            />
            <PostCardBody
                postImages={post.images}
                postID={post.id}
                username={user.username}
            />
            <PostCardFooter
                username={user.username}
                postLikes={likes}
                postDescription={post.description}
                liked={false}
            />
        </article>
    )
}
