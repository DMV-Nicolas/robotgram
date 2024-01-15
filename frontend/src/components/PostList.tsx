import { usePosts } from "../hooks/usePosts"
import { Post } from "./Post"

export function PostList() {
    const { posts } = usePosts()

    return (
        <ul>
            {
                posts.map((post) => (
                    <li key={post.id}>
                        <Post post={post} />
                    </li>
                ))
            }
        </ul>

    )
}