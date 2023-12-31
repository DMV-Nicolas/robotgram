import "./PostCard.css"

export type user = {
    _id: string
    username: string
    avatar: string
    gender: "male" | "female"
}

export type post = {
    _id: string
    user_id: string
    images: string[]
    description: string
    created_at: string
}

type postCardParams = {
    user: user
    post: post
    likes: number
}

export function PostCard({ post, user, likes }: postCardParams) {
    return (
        <article className="rg-postCard">
            <header className="rg-postCard-header">
                <img className="rg-postCard-avatar" src={user.avatar} alt="avatar" />
                <strong className="tg-postCard-username">@{user.username}</strong>
                <i>{user.gender}</i>
            </header>
            <div className="rg-postCard-body">
                {post.images.map((url) => <img src={url} rel="post_image" />)}
            </div>
            <footer className="rg-postCard-footer">
                <span>❤️ {likes}</span>
            </footer>
        </article>
    )
}