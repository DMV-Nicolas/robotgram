import { useParams } from 'react-router-dom'
import { NotFound } from '../components/NotFound'
import { usePost } from '../hooks/usePost'
import { useUserByID } from '../hooks/useUserByID'
import { PostModal } from '../components/PostModal'

export function PostModalPage() {
  const { postID } = useParams()
  if (postID === undefined) {
    return <NotFound />
  }
  const { post } = usePost({ postID })
  const { user } = useUserByID({ userID: post.userID })

  return (
    <PostModal
      user={user}
      post={post}
    />
  )
}
