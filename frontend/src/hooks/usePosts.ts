import { useEffect, useState } from "react"
import { PostType, PostResponse } from "../types"

export function usePosts() {
    const [posts, setPosts] = useState<PostType[]>([])
    useEffect(() => {
        const getPosts = async () => {
            const res = await fetch("http://localhost:5000/v1/posts?offset=0&limit=1000")
            const data: PostResponse[] = await res.json()
            const posts = data.map((dataPost) => {
                const post: PostType = {
                    id: dataPost.id,
                    userID: dataPost.user_id,
                    images: dataPost.images,
                    description: dataPost.description,
                    createdAt: dataPost.created_at,
                }
                return post
            })
            setPosts(posts)
        }

        getPosts()
    }, [])

    return { posts }
}