import { Text, Button, Menu, Portal, Box, type BoxProps } from "@chakra-ui/react"
import { useApi } from "@/lib/api"
import { useQuery, useQueryClient } from "@tanstack/react-query"
import { PiStarFill } from "react-icons/pi";
import { useEffect, useState, useMemo } from "react"
import { toaster } from "@/components/ui/toaster"
import { useTranslation } from "react-i18next"

interface UserBalanceProps extends BoxProps { }

export const UserBalance = (props: UserBalanceProps) => {
    const { transactionsApi } = useApi()
    const queryClient = useQueryClient()
    const { t } = useTranslation()
    const [balance, setBalance] = useState(0)

    const transactions = useQuery({
        queryKey: ['transactions'],
        queryFn: () => transactionsApi.getTransactions(),
        refetchInterval: 10000,
    })

    const createTransaction = (amount: number) => {
        transactionsApi.createTransaction((amount * 100).toString()).then(() => {
            queryClient.invalidateQueries({ queryKey: ['transactions'] })
        }).catch(() => {
            toaster.create({ title: t('error.unknown'), type: 'error' })
        })
    }

    const actions = useMemo(() => [
        {
            label: "100",
            icon: <PiStarFill color="#eab308" />,
            onClick: () => {
                createTransaction(100)
            }
        },
        {
            label: "200",
            icon: <PiStarFill color="#eab308" />,
            onClick: () => {
                createTransaction(200)
            }
        },
        {
            label: "500",
            icon: <PiStarFill color="#eab308" />,
            onClick: () => {
                createTransaction(500)
            }
        },
        {
            label: "1000",
            icon: <PiStarFill color="#eab308" />,
            onClick: () => {
                createTransaction(1000)
            }
        }
    ], [])

    useEffect(() => {
        if (transactions.data) {
            setBalance(Math.round(transactions.data.reduce((acc, transaction) => acc + Number(transaction.amount), 0) / 100))
        }
    }, [transactions.data])
    
    return (
        <Box {...props}>
            <Menu.Root>
                <Menu.Trigger asChild>
                    <Button variant="ghost" onClick={(e) => { e.stopPropagation() }}>
                        <Text fontWeight="bold" color="fg.muted">{balance}</Text>
                        <PiStarFill color="#eab308" />
                    </Button>
                </Menu.Trigger>
                <Portal>
                    <Menu.Positioner>
                        <Menu.Content>
                            {actions?.map((action, i) => (
                                <Menu.Item key={i} value={action.label} onClick={(e) => { e.stopPropagation(); action.onClick() }}>
                                    <Box flex="1">{action.label}</Box>
                                    {action.icon}
                                </Menu.Item>
                            ))}
                        </Menu.Content>
                    </Menu.Positioner>
                </Portal>
            </Menu.Root>
        </Box>
    )
}
