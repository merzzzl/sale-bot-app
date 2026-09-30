import './App.css'
import { Box, Container } from '@chakra-ui/react'
import { Main } from '@/containers/Main'
import { useTranslation } from 'react-i18next'
import './i18n'
import { useEffect, useState } from 'react'
import { Bot } from '@/containers/Bot'
import { MemoryRouter, Routes, Route, useLocation, useNavigate } from 'react-router-dom'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { Toaster } from '@/components/ui/toaster'
import { Product } from '@/containers/Product'
import { Stats } from '@/containers/Stats'
import { TelegramChromeSync } from './components/ui/telegram-chrome-sync'

function BackButtonBridge() {
  const location = useLocation()
  const navigate = useNavigate()

  useEffect(() => {
    const tg = (window as any)?.Telegram?.WebApp
    if (!tg) return

    const isRoot = location.pathname === '/'
    if (isRoot) {
      tg.BackButton.hide()
      return
    }
    const onBack = () => navigate(-1)
    tg.BackButton.show()
    tg.BackButton.onClick(onBack)
    return () => {
      try { tg.BackButton.offClick?.(onBack) } catch { }
    }
  }, [location.pathname, navigate])

  return null
}

function App() {
  const { i18n } = useTranslation()
  const [queryClient] = useState(() => new QueryClient())

  useEffect(() => {
    i18n.changeLanguage('ru')
  }, [])

  return (
    <Box w="100vw" h="100vh">
      <QueryClientProvider client={queryClient}>
        <Toaster />
        <TelegramChromeSync />
        <Container maxW="container.md" py={8}>
          <MemoryRouter>
            <BackButtonBridge />
            <Routes>
              <Route path="/" element={<Main />} />
              <Route path="/bot" element={<Bot />} />
              <Route path="/product" element={<Product />} />
              <Route path="/stats" element={<Stats />} />
            </Routes>
          </MemoryRouter>
        </Container>
      </QueryClientProvider>
    </Box>
  )
}

export default App