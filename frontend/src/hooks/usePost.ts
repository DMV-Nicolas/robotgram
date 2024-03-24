import { useEffect, useState } from 'react'
import { DEFAULT_POST } from '../constants'
import { type PostType, type PostResponse } from '../types'
import { toast } from 'sonner'

export function usePost({ postID }: { postID: string }) {
  const [post, setPost] = useState(DEFAULT_POST)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchGetPost = async () => {
      const res = await fetch(`http://localhost:5000/v1/posts/${postID}`)

      if (!res.ok) {
        toast.error('cannot get post data')
      }

      const data: PostResponse = await res.json()
      const post: PostType = {
        id: data.id,
        userID: data.user_id,
        images: data.images,
        description: data.description,
        createdAt: data.created_at
      }
      setPost(post)
      setLoading(false)
    }

    if (postID.length !== 24) {
      return
    }

    fetchGetPost()
  }, [postID])

  return { post, loading }
}
