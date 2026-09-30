import { useTgApi } from "@/lib/grammy"
import { useApi } from "@/lib/api"
import { useQuery } from "@tanstack/react-query"
import { toaster } from "@/components/ui/toaster"
import { useTranslation } from "react-i18next"
import { AvatarHeader } from "./ui/avatar-header"

export const BotAvatar = ({ botId }: { botId: string }) => {
    const { t } = useTranslation()
    const { botApi } = useApi()
    const { tgApi } = useTgApi()

    const bot = useQuery({
        queryKey: ['bot', botId],
        queryFn: () => botApi.getBot(botId),
    })

    const botDetails = useQuery({
        queryKey: ['botDetails', botId],
        queryFn: () => tgApi.getBotDetails(bot.data?.token as string),
        enabled: !!bot.data?.token,
    })

    if (bot.isError || botDetails.isError) {
        toaster.create({
            title: t('error.unknownError'),
            type: 'error',
        })
    }

    return (
        <AvatarHeader
            isLoading={bot.isLoading || botDetails.isLoading}
            photo_url={botDetails.data?.photo_url}
            name={botDetails.data?.name || t('bot-avatar.error-name')}
            username={botDetails.data?.username || t('bot-avatar.error-username')}
            active={bot.data?.active || false}
        />
    )
}

