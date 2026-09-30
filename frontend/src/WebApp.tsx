import './App.css'
import { AbsoluteCenter, Box, EmptyStateContent, EmptyStateDescription, EmptyStateIndicator, EmptyStateRoot, EmptyStateTitle, VStack } from '@chakra-ui/react'
import { useTranslation } from 'react-i18next'
import './i18n'
import { useEffect } from 'react'
import { BsTelegram } from "react-icons/bs"

function WebApp() {
  const { t, i18n } = useTranslation()

  useEffect(() => {
    i18n.changeLanguage('ru')
  }, [])

  return (
    <Box w="100vw" h="100vh">
      <AbsoluteCenter>
        <EmptyStateRoot maxWidth={370}>
          <EmptyStateContent>
            <EmptyStateIndicator>
              <BsTelegram />
            </EmptyStateIndicator>
            <VStack>
              <EmptyStateTitle>
                {t('browser.title')}
              </EmptyStateTitle>
              <EmptyStateDescription textAlign="center">
                {t('browser.body')}
              </EmptyStateDescription>
            </VStack>
          </EmptyStateContent>
        </EmptyStateRoot>
      </AbsoluteCenter>
    </Box>
  )
}

export default WebApp