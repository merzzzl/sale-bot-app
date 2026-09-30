import { useEffect, useState } from "react"
import { useApi } from "@/lib/api"
import { Button, HStack, Link, Stack } from "@chakra-ui/react"
import { PasswordInput, PasswordStrengthMeter } from "./ui/password-input"
import { FiPlus } from "react-icons/fi"
import { useTranslation } from "react-i18next"
import { Text } from "@chakra-ui/react"
import { useQuery, useQueryClient } from "@tanstack/react-query"
import { useTgApi } from "@/lib/grammy"

export const BotRegister = () => {
    const { t } = useTranslation()
    const { botApi } = useApi()
    const { tgApi } = useTgApi() 
    const queryClient = useQueryClient()

    const [token, setToken] = useState<string>("")
    const [notification, setNotification] = useState<string>("")
    const [strength, setStrength] = useState<number>(1)
    const [username, setUsername] = useState<string>("")
    const [isOpen, setIsOpen] = useState<boolean>(false)
    const [isTokenValid, setIsTokenValid] = useState<boolean>(false)

    const validateTokenFormat = (token: string): boolean => {
        const tokenPattern = /^\d{8,10}:[A-Za-z0-9_-]{35}$/
        return tokenPattern.test(token)
    }

    const botDetails = useQuery({
        queryKey: ['bot-details', token],
        queryFn: () => tgApi.getBotDetails(token),
        enabled: isTokenValid,
        retry: false,
    })

    useEffect(() => {
        if (!validateTokenFormat(token)) {
            setIsTokenValid(false)
        } else {
            setIsTokenValid(true)
        }
    }, [token])

    useEffect(() => {
        if (botDetails.data && isTokenValid) {
            setUsername(botDetails.data.username)
            setStrength(2)
            setNotification("")
        } else {
            setUsername("")
            setStrength(1)
            setNotification(t('bot-register.token-invalid'))
        }
    }, [botDetails.data, isTokenValid])

    const handleOpenChange = (open: boolean) => {
        setIsOpen(open)
        if (!open) {
            setToken("")
            setUsername("")
            setNotification("")
            setStrength(1)
        }
    }

    const registerBot = () => {
        botApi.createBot({ botToken: token }).then(() => {
            handleOpenChange(false)
            queryClient.invalidateQueries({ queryKey: ['bots'] })
        }).catch(() => {
            setNotification(t('error.unknown'))
        })
    }

    return (
        <Stack gap={2}>
            <Button colorPalette={isOpen ? 'red' : 'blue'} size="md" borderRadius="md" onClick={() => handleOpenChange(!isOpen)}>
                {!isOpen && <FiPlus />}
                {isOpen ? t('bot-register.discard-title') : t('bot-register.title')}
            </Button>
            <Stack display={isOpen ? 'block' : 'none'} pb={4}>
                <Text textAlign="left" pb={2} color="gray.600" fontSize="sm" >{t('bot-register.description')} <Link href="https://t.me/BotFather" target="_blank">@BotFather</Link></Text>
                <HStack gap={2} align="flex-start">
                    <Stack width="full">
                        <PasswordInput placeholder={t('bot-register.token-placeholder')} size="sm" value={token} onChange={(e) => { setToken(e.target.value) }} />
                        <PasswordStrengthMeter value={strength} max={2} />
                    </Stack>
                    <Button colorPalette="blue" size="md" variant="subtle" borderRadius="md" onClick={registerBot} disabled={!token.trim() || strength < 2}>
                        {t('bot-register.save-button')}
                    </Button>
                </HStack>
                {notification && (
                    <Text color="red.500" mt="2" fontSize="sm" textAlign="left" width="full" >
                        {notification}
                    </Text>
                )}
                {username && (
                    <Text color="green.500" mt="2" fontSize="sm" textAlign="left" width="full" >
                        @{username}
                    </Text>
                )}
            </Stack>
        </Stack>
    )
}
