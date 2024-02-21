import { Route, Routes } from 'react-router-dom'
import { Toaster } from 'sonner'
import { LoginPage } from './pages/Login'
import { HomePage } from './pages/Home'
import { NotFoundPage } from './pages/NotFound'
import { SignupPage } from './pages/Signup'
import { ProfilePage } from './pages/Profile'
import './App.css'
import { PostModalPage } from './pages/PostModal'

function App() {
  // TODO: create application router
  return (
    <main className="app">
      <Routes>
        <Route path='/' element={<HomePage />} >
          <Route path='post/:postID' element={<PostModalPage />} />
        </Route>
        <Route path='/login' element={<LoginPage />} />
        <Route path='/signup' element={<SignupPage />} />
        <Route path='/user/:userID' element={<ProfilePage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
      <Toaster richColors />
      {/* <Footer /> */}
    </main>
  )
}

export default App
