import ReactDOM from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import { TokenProvider } from './context/token.tsx'
import App from './App.tsx'
import './index.css'

const root = document.getElementById('root')
if (root instanceof HTMLElement) {
  ReactDOM.createRoot(root).render(
    <BrowserRouter>
      <TokenProvider>
        <App />
      </TokenProvider>
    </BrowserRouter>
  )
}
