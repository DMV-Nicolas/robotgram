import { useEffect, useState } from 'react'
import { type ListCommentsResponse, type CommentType } from '../types'
import { toast } from 'sonner'

export function useComments({ targetID }: { targetID: string }) {
  const [comments, setComments] = useState<CommentType[]>([])

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
      setComments(comments)
    }

    fetchComments()
  }, [targetID])

  return { comments }
}
