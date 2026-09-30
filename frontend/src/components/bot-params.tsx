import {
    Button, Portal, Span, Select, Stack,
    Textarea, createListCollection, Skeleton,
    Field,
} from "@chakra-ui/react"
import { useTranslation } from "react-i18next"
import { useApi } from "@/lib/api"
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { toaster } from "@/components/ui/toaster"
import { FiSave } from "react-icons/fi"
import { useEffect, useState } from "react"

export const BotParams = ({ botId }: { botId: string }) => {
    const { t } = useTranslation()
    const { botApi, settingsApi } = useApi()
    const queryClient = useQueryClient()

    const bot = useQuery({
        queryKey: ['bot', botId],
        queryFn: () => botApi.getBot(botId),
    })

    const [instructions, setInstructions] = useState(bot.data?.prompt || '')
    const [presetId, setPresetId] = useState(bot.data?.presetId || '')
    const [presetCollection, setPresetCollection] = useState(createListCollection({
        items: [],
        itemToValue: (item: any) => item.presetId,
        itemToString: (item: any) => item.name,
    }))

    const promptPresets = useQuery({
        queryKey: ['promptPresets'],
        queryFn: () => settingsApi.getPromptPresets(),
    })

    useEffect(() => {
        if (bot.data && promptPresets.data) {
            setPresetCollection(createListCollection({
                items: promptPresets.data,
                itemToValue: (item: any) => item.presetId,
                itemToString: (item: any) => item.name,
            }))

            setInstructions(bot.data.prompt)
            setPresetId(bot.data.presetId)
        }
    }, [bot.data, promptPresets.data])

    const onSave = () => {
        if (!bot.data) {
            return
        }

        botApi.updateBot(bot.data?.botId, { prompt: instructions, presetId: presetId }).then(() => {
            queryClient.invalidateQueries({ queryKey: ['bot', botId] })
            toaster.create({ title: t('bot-params.save-success'), type: 'success' })
        }).catch(() => {
            toaster.create({ title: t('error.unknown'), type: 'error' })
        })
    }

    if (bot.isError || promptPresets.isError) {
        toaster.create({
            title: t('error.unknown'),
            type: 'error',
        })
    }

    if (bot.isLoading || promptPresets.isLoading) {
        return <BotParamsSkeleton />
    }

    return (
        <Stack gap="4">
            <Field.Root>
                <Field.Label>{t('bot-params.preset-label')}</Field.Label>
                <Select.Root collection={presetCollection} defaultValue={[presetId]} value={[presetId]} onValueChange={(value) => setPresetId(value.items[0].presetId)}>
                    <Select.Control>
                        <Select.Trigger>
                            <Select.ValueText placeholder={t('bot-params.preset-placeholder')} />
                        </Select.Trigger>
                        <Select.IndicatorGroup>
                            <Select.Indicator />
                        </Select.IndicatorGroup>
                    </Select.Control>
                    <Portal>
                        <Select.Positioner>
                            <Select.Content>
                                {presetCollection.items.map((preset: any) => (
                                    <Select.Item item={preset} key={preset.presetId}>
                                        <Stack gap="0">
                                            <Select.ItemText>{preset.name}</Select.ItemText>
                                            <Span color="fg.muted" textStyle="xs">
                                                {preset.description}
                                            </Span>
                                        </Stack>
                                        <Select.ItemIndicator />
                                    </Select.Item>
                                ))}
                            </Select.Content>
                        </Select.Positioner>
                    </Portal>
                </Select.Root>
                <Field.HelperText textAlign="left">{t('bot-params.preset-description')}</Field.HelperText>
            </Field.Root>
            <Field.Root>
                <Field.Label>{t('bot-params.instructions-label')}</Field.Label>
                <Textarea variant="subtle" placeholder={t('bot-params.instructions-placeholder')} value={instructions} onChange={(e) => setInstructions(e.target.value)} resize="vertical" h="150px" />
                <Field.HelperText textAlign="left">{t('bot-params.instructions-description')}</Field.HelperText>
            </Field.Root>
            <Button colorPalette="blue" onClick={() => onSave()}>
                <FiSave />
                {t('bot-params.save-button')}
            </Button>
        </Stack>
    )
}

export const BotParamsSkeleton = () => {
    return (
        <Stack gap="4">
            <Skeleton height="90px" />
            <Skeleton height="230px" />
            <Skeleton height="40px" />
        </Stack>
    )
}
