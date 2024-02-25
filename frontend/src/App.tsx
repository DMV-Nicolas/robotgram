import { Route, Routes, matchPath, useLocation } from 'react-router-dom'
import { Toaster } from 'sonner'
import { LoginPage } from './pages/Login'
import { HomePage } from './pages/Home'
import { NotFoundPage } from './pages/NotFound'
import { SignupPage } from './pages/Signup'
import { ProfilePage } from './pages/Profile'
import { PostModalPage } from './pages/PostModal'
import './App.css'
import { useEffect, useState } from 'react'
import { Footer } from './components/Footer'

function App() {
  const location = useLocation()
  const [previousLocation, setPreviousLocation] = useState('')
  const modalPages = ['/post/:postID']

  const shouldUpdatePreviousLocation = () => {
    const f = modalPages.filter((path) => (
      matchPath(path, location.pathname) !== null
    ))
    return f.length !== 1
  }

  useEffect(() => {
    if (shouldUpdatePreviousLocation()) {
      setPreviousLocation(location.pathname)
    }
  }, [location.pathname])

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
        <Route path='*' element={<></>} />
      </Routes>
      <Toaster richColors />
      <Footer previousLocation={previousLocation} />
    </main>
  )
}

export default App
