import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { Provider } from "./components/ui/provider"
import './index.css'
import App from './App.tsx'
import { init, isTMA, miniApp, themeParams } from '@tma.js/sdk-react';
import WebApp from './WebApp.tsx'

const initializeTelegramSDK = () => {
  try {
    init()
    themeParams.mount();
    miniApp.mount();
    console.log('Telegram SDK initialized')
  } catch (error) {
    console.error('Failed to initialize Telegram SDK:', error)
  }
}

const inTelegram = isTMA()

if (inTelegram) {
  initializeTelegramSDK()
}

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Provider>
      {inTelegram ? <App /> : <WebApp />}
    </Provider>
  </StrictMode>,
)
