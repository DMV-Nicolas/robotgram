function PostCardHeader() {
    return (
        <header className="rg-postCard-header">
            <img className="rg-postCard-avatar" src="https://unavatar.io/dmvnicolas" alt="avatar" />
            <div>
                <strong>Nicolas David</strong>
                <span>@dmvnicolas</span>
            </div>
        </header>
    )
}

function PostCardBody() {
    return (
        <body>
            <img src="https://png.pngtree.com/background/20230612/original/pngtree-wolf-animals-images-wallpaper-for-pc-384x480-picture-image_3180467.jpg" alt="post" />
        </body>
    )
}

function PostCardFooter() {
    return (
        <footer>
            <span>Like</span>
            <span>Comments</span>
        </footer>
    )
}

export function PostCard() {
    return (
        <article className="rg-postCard">
            <PostCardHeader />
            <PostCardBody />
            <PostCardFooter />
        </article>
    )
}