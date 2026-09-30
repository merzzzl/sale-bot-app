import { useTranslation } from "react-i18next"
import { useApi } from "@/lib/api"
import { useQueryClient, useQuery } from "@tanstack/react-query"
import { useState } from "react"
import { toaster } from "@/components/ui/toaster"
import { Field, Alert, Switch, IconButton, Dialog, Portal, Stack, Skeleton, SwitchIndicator, Button, Text } from "@chakra-ui/react"
import { FiPause, FiPlay, FiTrash } from "react-icons/fi"
import { useNavigate } from "react-router-dom"

export const BotActions = ({ botId }: { botId: string }) => {
    const { t } = useTranslation()
    const { botApi } = useApi()
    const queryClient = useQueryClient()
    const navigate = useNavigate()

    const [confirmOpen, setConfirmOpen] = useState(false)

    const bot = useQuery({
        queryKey: ['bot', botId],
        queryFn: () => botApi.getBot(botId),
    })

    const isActive = !!bot.data?.active

    const onActivate = () => {
        botApi.activateBot(botId).then(() => {
            queryClient.invalidateQueries({ queryKey: ['bot', botId] })
            toaster.create({ title: t('bot-actions.activate-success'), type: 'success' })
        }).catch(() => {
            toaster.create({ title: t('error.unknown'), type: 'error' })
        })
    }

    const onDeactivate = () => {
        botApi.deactivateBot(botId).then(() => {
            queryClient.invalidateQueries({ queryKey: ['bot', botId] })
            toaster.create({ title: t('bot-actions.deactivate-success'), type: 'success' })
        }).catch(() => {
            toaster.create({ title: t('error.unknown'), type: 'error' })
        })
    }

    const onDelete = () => {
        botApi.deleteBot(botId).then(() => {
            queryClient.invalidateQueries({ queryKey: ['bot', botId] })
            navigate('/')
        }).catch(() => {
            toaster.create({ title: t('error.unknown'), type: 'error' })
        })
    }

    const handleToggle = () => (!!bot.data?.active ? onDeactivate() : onActivate())

    if (bot.isLoading) {
        return <BotActionsSkeleton />
    }

    return (
        <Stack gap="4">
            <Field.Root>
                <Field.Label>{t('bot-actions.status-label')}</Field.Label>
                <Alert.Root status="info" variant="subtle" justifyContent="space-between" alignItems="center">
                    <Alert.Description textAlign="left">{t('bot-actions.status-description')}</Alert.Description>
                    <Switch.Root size="lg" checked={isActive} onCheckedChange={() => handleToggle()}>
                        <Switch.HiddenInput />
                        <Switch.Control alignSelf="flex-end">
                            <Switch.Thumb />
                        </Switch.Control>
                        <SwitchIndicator fallback={<FiPause />}>
                            <FiPlay />
                        </SwitchIndicator>
                    </Switch.Root>
                </Alert.Root>
            </Field.Root>
            <Field.Root>
                <Field.Label>{t('bot-actions.delete-label')}</Field.Label>
                <Alert.Root status="warning" variant="subtle" justifyContent="space-between" alignItems="center">
                    <Alert.Description textAlign="left">{t('bot-actions.delete-description')}</Alert.Description>
                    <IconButton colorPalette="red" onClick={() => setConfirmOpen(true)}>
                        <FiTrash />
                    </IconButton>
                </Alert.Root>
                <Dialog.Root open={confirmOpen} onOpenChange={(e) => setConfirmOpen(e.open)} placement="center">
                    <Portal>
                        <Dialog.Positioner background="blackAlpha.600">
                            <Dialog.Content w="90vw" maxW="420px">
                                <Dialog.Header>
                                    <Dialog.Title>{t('bot-actions.delete-button')}</Dialog.Title>
                                </Dialog.Header>
                                <Dialog.Body>
                                    <Text>{t('bot-actions.delete-dialog-description')}</Text>
                                </Dialog.Body>
                                <Dialog.Footer>
                                    <Button colorPalette="red" onClick={() => onDelete()}>
                                        <FiTrash />
                                        {t('bot-actions.delete-button')}
                                    </Button>
                                    <Dialog.ActionTrigger asChild>
                                        <Button variant="outline">{t('bot-actions.delete-discard-button')}</Button>
                                    </Dialog.ActionTrigger>
                                </Dialog.Footer>
                            </Dialog.Content>
                        </Dialog.Positioner>
                    </Portal>
                </Dialog.Root>
            </Field.Root>
        </Stack>
    )
}

export const BotActionsSkeleton = () => {
    return (
        <Stack gap="4">
            <Skeleton height="90px" />
            <Skeleton height="90px" />
        </Stack>
    )
}