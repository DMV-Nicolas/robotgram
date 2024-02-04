import { Navbar } from '../components/Navbar'
import { PostList } from '../components/PostList'
import { IsLoggedMiddleware } from './IsLoggedMiddleware'

export function HomePage() {
  return (
    <>
      <IsLoggedMiddleware>
        <Navbar />
        <PostList />
      </IsLoggedMiddleware>
    </>
  )
}
