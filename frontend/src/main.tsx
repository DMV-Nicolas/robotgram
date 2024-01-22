import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'

const root = document.getElementById('root')
if (root instanceof HTMLElement) {
  ReactDOM.createRoot(root).render(<App />)
}
