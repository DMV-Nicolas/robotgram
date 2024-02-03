import { Route, Routes } from 'react-router-dom'
import { LoginPage } from './pages/Login'
import { HomePage } from './pages/Home'
import { NotFoundPage } from './pages/NotFound'
import { SignupPage } from './pages/Signup'
import { Footer } from './components/Footer'
import './App.css'

function App() {
  // TODO: create application router
  return (
    <main className="app">
      <Routes>
        <Route path='/' element={<HomePage />} />
        <Route path='/login' element={<LoginPage />} />
        <Route path='/signup' element={<SignupPage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
      <Footer />
    </main>
  )
}

export default App
