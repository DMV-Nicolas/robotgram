import { PostCard, user, post } from './PostCard'
import './App.css'


function App() {
  let p: post = {
    _id: "123",
    user_id: "123",
    images: ["https://img.freepik.com/vector-gratis/juguete-robot-vintage-sobre-fondo-blanco_1308-77501.jpg", "https://www.lavanguardia.com/andro4all/hero/2023/10/fabrica-de-robots.png", "c"],
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
