import { type Bot } from "@/lib/api"
import { useMemo, useRef, useState } from "react"
import { HorizontalCard, HorizontalCardSkeleton } from "./ui/horizontal-card"
import { FiBarChart, FiPause, FiPlay, FiTrash } from "react-icons/fi"
import { useQuery, useQueries, type UseQueryResult, useQueryClient } from "@tanstack/react-query"
import { useApi } from "@/lib/api"
import { useTgApi, type BotDetails } from "@/lib/grammy"
import { Stack, type ColorPalette } from "@chakra-ui/react"
import { useTranslation } from "react-i18next"
import { DeleteDialog, type DeleteDialogRef } from "./delete-dialog"
import { useNavigate } from "react-router-dom"

export const BotList = () => {
    const { t } = useTranslation()
    const { botApi } = useApi()
    const { tgApi } = useTgApi()
    const queryClient = useQueryClient()    
    const navigate = useNavigate()

    const deleteDialog = useRef<DeleteDialogRef>(null)
    const [deleteDialogOpen, setDeleteDialogOpen] = useState(false)

    const bots = useQuery({
        queryKey: ["bots"],
        queryFn: () => botApi.getBots(),
    })

    const botsDetails = useQueries({
        queries: (bots.data ?? []).map((bot: Bot) => ({
            queryKey: ["botDetails", bot.botId],
            queryFn: () => tgApi.getBotDetails(bot.token),
        })),
    }) as UseQueryResult<BotDetails, Error>[]

    const onActivate = (botId: string) => {
        botApi.activateBot(botId).then(() => {
            queryClient.invalidateQueries({ queryKey: ["bots"] })
        })
    }

    const onDeactivate = (botId: string) => {
        botApi.deactivateBot(botId).then(() => {
            queryClient.invalidateQueries({ queryKey: ["bots"] })
        })
    }

    const onDelete = (botId: string) => {
        botApi.deleteBot(botId).then(() => {
            queryClient.invalidateQueries({ queryKey: ["bots"] })
        })
    }

    const onDeleteDialog = (botId: string) => {
        deleteDialog.current?.setOnApprove(() => {
            onDelete(botId)
            setDeleteDialogOpen(false)
        })
        deleteDialog.current?.setOnDiscard(() => setDeleteDialogOpen(false))
        setDeleteDialogOpen(true)
    }

    const onClick = (botId: string) => {
        navigate('/bot', { state: { botId } })
    }

    const onClickStats = (botId: string) => {
        navigate('/stats', { state: { botId } })
    }

    const botItems = useMemo(() => {
        const list = bots.data ?? []
        return list.map((bot: Bot, i: number) => {
            const details = botsDetails[i]
            const badges = !bot.active
                ? [{
                    notification: "Inactive",
                    colorPalette: "red" as ColorPalette,
                }]
                : []
            const actions = !bot.active
                ? [{
                    icon: <FiPlay />,
                    label: t('bot-list.activate'),
                    onClick: () => onActivate(bot.botId),
                }] : [{
                    icon: <FiPause />,
                    label: t('bot-list.deactivate'),
                    onClick: () => onDeactivate(bot.botId),
                }]

            actions.push({
                icon: <FiBarChart />,
                label: t('bot-list.stats'),
                onClick: () => onClickStats(bot.botId),
            })

            actions.push({
                icon: <FiTrash />,
                label: t('bot-list.delete'),
                onClick: () => onDeleteDialog(bot.botId),
            })

            let description = details.data?.username
            if (!description) {
                description = t('bot-list.error-username')
            } else {
                description = "@" + description
            }


            return (
                <HorizontalCard
                    key={bot.botId}
                    isLoading={details.isLoading}
                    onClick={() => onClick(bot.botId)}
                    img={details.data?.photo_url || ""}
                    title={details.data?.name || t('bot-list.error-name')}
                    description={description}
                    badges={badges}
                    actions={actions}
                />
            )
        })
    }, [bots.data, botsDetails, onClick])

    return (
        <Stack gap={4} >
            {bots.isLoading ? (
                <Stack gap={2} borderRadius={8} backgroundColor="bg.subtle" p={0}>
                    {[...Array(3)].map((_, i) => (
                        <HorizontalCardSkeleton key={i} withActions withBadges />
                    ))}
                </Stack>
            ) : (
                <Stack gap={2} backgroundColor="bg.subtle" borderRadius="md">
                    {botItems}
                </Stack>
            )}
            <DeleteDialog
                ref={deleteDialog}
                open={deleteDialogOpen}
                description={t('bot-list.delete-description')}
            />
        </Stack>
    )
}
