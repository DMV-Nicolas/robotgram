import { Navbar } from '../components/Navbar'
import { PostList } from '../components/PostList'
import { useUserID } from '../hooks/useUserID'
import { IsLoggedMiddleware } from './IsLoggedMiddleware'

export function HomePage() {
  const { userID } = useUserID()
  return (
    <>
      <IsLoggedMiddleware>
        <Navbar userID={userID} />
        <PostList />
      </IsLoggedMiddleware>
    </>
  )
}
