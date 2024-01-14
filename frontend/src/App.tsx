import { Post } from './components/Post'
import { usePosts } from './hooks/usePosts'
import "./App.css"

function App() {
    const { posts } = usePosts()

    return (
        <main>
            <ul>
                {
                    posts.map((post) => (
                        <li key={post.id}>
                            <Post post={post} />
                        </li>
                    ))
                }
            </ul>
        </main>
    )
}

export default App
