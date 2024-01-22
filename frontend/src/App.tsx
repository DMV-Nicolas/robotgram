import { PostList } from './components/PostList'
import { Navbar } from './components/Navbar'
import './App.css'

function App() {
  // TODO: create application router

  return (
    <main className="app">
      <Navbar />
      <PostList />
    </main>
  )
}

export default App
