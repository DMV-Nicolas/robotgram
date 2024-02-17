import { usePosts } from '../hooks/usePosts'
import { Post } from './Post'
import './PostList.css'

export function PostList() {
  const { posts } = usePosts({})

  return (
    <ul className="postList">
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
