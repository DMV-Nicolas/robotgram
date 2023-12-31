import { PostCard, user, post } from './PostCard'
import './App.css'


function App() {
  let p: post = {
    _id: "123",
    user_id: "123",
    images: ["a", "b", "c"],
    description: "arroz con agua",
    created_at: "ayer"
  }

  let u: user = {
    _id: "123",
    username: "dmvnicolas",
    avatar: "https://unavatar.io/dmvnicolas",
    gender: "male"
  }

  return (
    <>
      <h1>Robotgram</h1>
      <PostCard post={p} user={u} likes={10} />
    </>
  )
}

export default App
