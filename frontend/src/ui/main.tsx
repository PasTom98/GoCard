import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'

const root = createRoot(document.getElementById('root')!)

root.render(
  <StrictMode>
    <App />
  </StrictMode>,
)

// Hot Module Replacement (HMR) - Enable React Fast Refresh
if (import.meta.hot) {
  import.meta.hot.accept()
}
