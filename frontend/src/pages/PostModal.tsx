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
  const { post, loading: postLoading } = usePost({ postID })
  const { user, loading: userLoading } = useUserByID({ userID: post.userID })

  return (
    <PostModal
      post={post}
      user={user}
      loading={postLoading && userLoading}
    />
  )
}
