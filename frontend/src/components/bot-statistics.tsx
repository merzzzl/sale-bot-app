import { useApi } from "@/lib/api"
import { useQuery } from "@tanstack/react-query"
import { Area, AreaChart } from "recharts"
import { Chart, useChart } from "@chakra-ui/charts"
import { Box, Stack, Text } from "@chakra-ui/react"
import { timeFrom, timeParse } from "@/lib/timesec"
import { Card, Stat, SegmentGroup } from "@chakra-ui/react"
import { useTranslation } from "react-i18next"
import { FiCreditCard, FiMessageSquare, FiUsers } from "react-icons/fi"
import { useState } from "react"
import { timeNow } from "@/lib/timesec"

export interface BotStatisticsProps {
    botId: string
}

export const BotStatistics = ({ botId }: BotStatisticsProps) => {
    const { botApi } = useApi()
    const { t } = useTranslation()

    const [startDate, setStartDate] = useState(timeNow() - timeFrom({ days: 7 }))
    const [endDate, setEndDate] = useState(timeNow() - timeFrom({ minutes: 5 }))

    const statistics = useQuery({
        queryKey: ['botStatistics', botId, startDate, endDate],
        queryFn: () => botApi.getBotStatistics(botId, startDate.toString(), endDate.toString()),
    })

    const messages = statistics.data?.filter((item) => item.type === 'MESSAGES') ?? []
    const chats = statistics.data?.filter((item) => item.type === 'CHATS') ?? []
    const payments = statistics.data?.filter((item) => item.type === 'PAYMENTS') ?? []

    const STEP = timeFrom({ hours: 8 })

    const buildSlots = (from: number, to: number, step: number): number[] => {
        const start = Math.floor(from / step) * step
        const end = Math.floor(to / step) * step
        const slots: number[] = []
        for (let t = start; t <= end; t += step) {
            slots.push(t)
        }
        return slots
    }

    const slots = buildSlots(startDate, endDate, STEP)

    const aggregateBySlot = (records: Array<{ timestamp: string, value: string }>) => {
        const sums: Record<number, number> = {}
        for (const item of records) {
            const ts = parseInt(item.timestamp)
            const val = Number(item.value) || 0
            const slot = Math.floor(ts / STEP) * STEP
            sums[slot] = (sums[slot] ?? 0) + val
        }
        return slots.map((s) => ({
            name: timeParse(s).toISOString().replace('T', ' ').slice(0, 16), // 'YYYY-MM-DD HH:MM'
            value: sums[s] ?? 0,
        }))
    }

    const messagesData = aggregateBySlot(messages as any)
    const chatsData = aggregateBySlot(chats as any)
    const paymentsData = aggregateBySlot(payments as any)

    const messagesChart = useChart({
        data: messagesData,
        series: [{ color: 'teal.solid' }],
    })

    const chatsChart = useChart({
        data: chatsData,
        series: [{ color: 'teal.solid' }],
    })

    const paymentsChart = useChart({
        data: paymentsData,
        series: [{ color: 'teal.solid' }],
    })

    const onSegmentChange = ({ value }: { value: string | null }) => {
        if (value === t('bot-statistics.7-days')) {
            setStartDate(timeNow() - timeFrom({ days: 7 }))
            setEndDate(timeNow() - timeFrom({ minutes: 5 }))
        } else if (value === t('bot-statistics.30-days')) {
            setStartDate(timeNow() - timeFrom({ days: 30 }))
            setEndDate(timeNow() - timeFrom({ minutes: 5 }))
        } else if (value === t('bot-statistics.90-days')) {
            setStartDate(timeNow() - timeFrom({ days: 90 }))
            setEndDate(timeNow() - timeFrom({ minutes: 5 }))
        }
    }

    return (
        <Stack gap={4}>
            <Text fontSize="sm" color="fg.muted">{t('bot-statistics.note')}</Text>
            <Box>
                <SegmentGroup.Root w="fit-content" size="sm" defaultValue={t('bot-statistics.7-days')} onValueChange={onSegmentChange}>
                    <SegmentGroup.Indicator />
                    <SegmentGroup.Items items={[t('bot-statistics.7-days'), t('bot-statistics.30-days'), t('bot-statistics.90-days')]} />
                </SegmentGroup.Root>
            </Box>
            <Card.Root size="sm" overflow="hidden">
                <Card.Body>
                    <Stat.Root>
                        <Stat.Label>
                            <FiMessageSquare /> {t('bot-statistics.messages')}
                        </Stat.Label>
                        <Stat.ValueText>{messagesChart.getTotal('value')}</Stat.ValueText>
                    </Stat.Root>
                </Card.Body>
                <Chart.Root height={8} chart={messagesChart}>
                    <AreaChart data={messagesChart.data} margin={{ top: 0, right: 0, left: 0, bottom: 10 }}>
                        {messagesChart.series.map((item) => (
                            <Area
                                type="natural"
                                key={item.name}
                                isAnimationActive={false}
                                dataKey={messagesChart.key(item.name)}
                                fill={messagesChart.color(item.color)}
                                fillOpacity={0.2}
                                stroke={messagesChart.color(item.color)}
                                strokeWidth={2}
                            />
                        ))}
                    </AreaChart>
                </Chart.Root>
            </Card.Root>
            <Text fontSize="sm" color="fg.muted" textAlign="left">{t('bot-statistics.messages-description')}</Text>
            <Card.Root size="sm" overflow="hidden">
                <Card.Body>
                    <Stat.Root>
                        <Stat.Label>
                            <FiUsers /> {t('bot-statistics.chats')}
                        </Stat.Label>
                        <Stat.ValueText>{chatsChart.getTotal('value')}</Stat.ValueText>
                    </Stat.Root>
                </Card.Body>
                <Chart.Root height={8} chart={chatsChart}>
                    <AreaChart
                        data={chatsChart.data}
                        margin={{ top: 0, right: 0, left: 0, bottom: 10 }}
                    >
                        {chatsChart.series.map((item) => (
                            <Area
                                key={item.name}
                                isAnimationActive={false}
                                dataKey={chatsChart.key(item.name)}
                                fill={chatsChart.color(item.color)}
                                fillOpacity={0.2}
                                stroke={chatsChart.color(item.color)}
                                strokeWidth={2}
                            />
                        ))}
                    </AreaChart>
                </Chart.Root>
            </Card.Root>
            <Text fontSize="sm" color="fg.muted" textAlign="left">{t('bot-statistics.chats-description')}</Text>
            <Card.Root size="sm" overflow="hidden">
                <Card.Body>
                    <Stat.Root>
                        <Stat.Label>
                            <FiCreditCard /> {t('bot-statistics.payments')}
                        </Stat.Label>
                        <Stat.ValueText>{paymentsChart.getTotal('value')}</Stat.ValueText>
                    </Stat.Root>
                </Card.Body>
                <Chart.Root height={8} chart={paymentsChart}>
                    <AreaChart
                        data={paymentsChart.data}
                        margin={{ top: 0, right: 0, left: 0, bottom: 10 }}
                    >
                        {paymentsChart.series.map((item) => (
                            <Area
                                key={item.name}
                                isAnimationActive={false}
                                dataKey={paymentsChart.key(item.name)}
                                fill={paymentsChart.color(item.color)}
                                fillOpacity={0.2}
                                stroke={paymentsChart.color(item.color)}
                                strokeWidth={2}
                            />
                        ))}
                    </AreaChart>
                </Chart.Root>
            </Card.Root>
            <Text fontSize="sm" color="fg.muted" textAlign="left">{t('bot-statistics.payments-description')}</Text>
        </Stack>
    )
}