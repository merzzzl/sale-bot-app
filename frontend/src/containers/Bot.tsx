import { Stack } from "@chakra-ui/react"
import { BotAvatar } from "@/components/bot-avatar"
import { BotTabs } from "@/components/bot-tabs"
import { useLocation, useNavigate } from "react-router-dom"

export const Bot = () => {
    const { state } = useLocation() as { state?: { botId?: string } }
    const navigate = useNavigate()
    const botId = state?.botId

    if (!botId) {
        navigate('/')
        return null
    }

    return (
        <Stack gap={4}>
            <BotAvatar botId={botId} />
            <BotTabs botId={botId} />
        </Stack>
    )
}

export default Bot