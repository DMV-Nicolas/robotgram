import { useState } from "react"
import { usePost } from "../hooks/usePost"
import { getTimeElapsed } from "../services/time"
import { Comment, EmptyHeart, Heart, Options, Save, Share } from "./Icons"
import { PostType } from "../types"
import "./Post.css"

type PostCardHeaderParams = {
    username: string
    userAvatar: string
    postCreatedAt: string
}

function PostHeader({ username, userAvatar, postCreatedAt }: PostCardHeaderParams) {
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

type PostCardBodyParams = {
    postImages: string[]
    postID: string,
    username: string
}

function PostBody({ postImages, postID, username }: PostCardBodyParams) {
    const [slide, setSlide] = useState(0)

    const prevSlide = () => {
        if (slide > 0) setSlide(slide - 1)
    }

    const nextSlide = () => {
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

function PostFooter({ username, postLikes, postDescription, liked }: PostCardFooterParams) {
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

type PostParams = {
    post: PostType
}

export function Post({ post }: PostParams) {
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
