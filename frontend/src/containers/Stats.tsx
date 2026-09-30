import { Stack } from "@chakra-ui/react"
import { BotAvatar } from "@/components/bot-avatar"
import { BotStatistics } from "@/components/bot-statistics"
import { useLocation, useNavigate } from "react-router-dom"

export const Stats = () => {
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
            <BotStatistics botId={botId} />
        </Stack>
    )
}

export default Stats