import { useEffect, useState } from 'react'
import { useToken } from './useToken'
import { toast } from 'sonner'
import { type ListCommentsResponse, type CommentType } from '../types'

export function useComments({ targetID }: { targetID: string }) {
  const [comments, setComments] = useState<CommentType[]>([])
  const [reListComments, setReListComments] = useState(false)
  const { accessToken, refreshAccessToken, updateAccessToken, updateRefreshToken } = useToken()

  const createComment = async ({ content }: { content: string }) => {
    const res = await fetch('http://localhost:5000/v1/comments', {
      method: 'POST',
      body: JSON.stringify({ target_id: targetID, content }),
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
        Authorization: `Bearer ${accessToken}`
      },
      credentials: 'include'
    })

    if (!res.ok) {
      toast.error('cannot create comment')
      const err = await refreshAccessToken()
      if (err instanceof Error) {
        updateAccessToken('')
        updateRefreshToken('')
      }
    }

    setReListComments((prevReListComments) => !prevReListComments)
  }

  useEffect(() => {
    const fetchComments = async () => {
      const res = await fetch(`http://localhost:5000/v1/comments/${targetID}?offset=0&limit=100`)

      if (!res.ok) {
        toast.error('cannot list comments')
        return
      }

      const data: ListCommentsResponse = await res.json()
      if (data === null) return

      const comments = data.map((dataComment) => {
        const comment: CommentType = {
          id: dataComment.id,
          userID: dataComment.user_id,
          targetID: dataComment.target_id,
          content: dataComment.content,
          createdAt: dataComment.created_at
        }
        return comment
      })
      setComments(comments.reverse())
    }

    fetchComments()
  }, [targetID, reListComments])

  return { comments, createComment }
}
