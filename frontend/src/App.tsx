import { Route, Routes } from 'react-router-dom'
import { Toaster } from 'sonner'
import { LoginPage } from './pages/Login'
import { HomePage } from './pages/Home'
import { NotFoundPage } from './pages/NotFound'
import { SignupPage } from './pages/Signup'
import { ProfilePage } from './pages/Profile'
import { PostModalPage } from './pages/PostModal'
import { CreatePostPage } from './pages/CreatePost'
import { useModals } from './hooks/useModals'
import './App.css'

function App() {
  const { previousLocation } = useModals({ modalPages: ['/post/:postID', '/create'] })

  return (
    <main className="app">
      <Routes location={{ pathname: previousLocation }}>
        <Route path='/' element={<HomePage />} ></Route>
        <Route path='/login' element={<LoginPage />} />
        <Route path='/signup' element={<SignupPage />} />
        <Route path='/user/:userID' element={<ProfilePage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
      <Routes>
        <Route path='post/:postID' element={<PostModalPage />} />
        <Route path='create' element={<CreatePostPage />} />
        <Route path='*' element={<></>} />
      </Routes>
      <Toaster richColors />
    </main>
  )
}

export default App
