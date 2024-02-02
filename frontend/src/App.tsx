import { Route, Routes } from 'react-router-dom'
import { LoginPage } from './pages/Login'
import { HomePage } from './pages/Home'
import { NotFoundPage } from './pages/NotFound'
import './App.css'

function App() {
  // TODO: create application router

  return (
    <main className="app">
      <Routes>
        <Route path='/' element={<HomePage />} />
        <Route path='/login' element={<LoginPage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </main>
  )
}

export default App
