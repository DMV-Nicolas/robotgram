import { useEffect, useState } from 'react'
import { type PostType, type ListPostsResponse } from '../types'

export function usePosts({ userID }: { userID?: string }) {
  const [posts, setPosts] = useState<PostType[]>([])
  useEffect(() => {
    const fetchListPosts = async () => {
      let url = 'http://localhost:5000/v1/posts?offset=0&limit=5'
      if (userID !== undefined) {
        url += `&user_id=${userID}`
      }

      const res = await fetch(url)
      const data: ListPostsResponse = await res.json()
      if (data === null) {
        setPosts([])
        return
      }

      const posts = data.map((dataPost) => {
        const post: PostType = {
          id: dataPost.id,
          userID: dataPost.user_id,
          images: dataPost.images,
          description: dataPost.description,
          createdAt: dataPost.created_at
        }
        return post
      })
      setPosts(posts)
    }

    fetchListPosts()
  }, [userID])

  return { posts }
}
