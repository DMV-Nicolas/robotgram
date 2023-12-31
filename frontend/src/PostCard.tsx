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
    let genderEmoji: string
    if (user.gender == "male") {
        genderEmoji = "♂️"
    } else {
        genderEmoji = "♀️"
    }

    return (
        <article className="rg-postCard">
            <header className="rg-postCard-header">
                <img className="rg-postCard-avatar" src={user.avatar} alt="avatar" />
                <strong className="rg-postCard-username">{user.username} {genderEmoji}</strong>
            </header>
            <div className="rg-postCard-body">
                {post.images.map((url, i) => <img className="rg-postCard-image" src={url} rel="post_image" key={i + post._id} />)}
            </div>
            <footer className="rg-postCard-footer">
                <span>❤️ {likes}</span>
            </footer>
        </article>
    )
}