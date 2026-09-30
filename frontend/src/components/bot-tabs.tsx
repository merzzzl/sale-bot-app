import { Stack, Tabs } from "@chakra-ui/react"
import { useTranslation } from "react-i18next"
import { FiSettings, FiUser, FiShoppingCart } from "react-icons/fi"
import { ProductList } from "./product-list"
import { BotActions } from "./bot-actions"
import { BotParams } from "./bot-params"
import { ProductAdd } from "./product-add"

export const BotTabs = ({ botId }: { botId: string }) => {
    const { t } = useTranslation()

    return (
        <Tabs.Root defaultValue="products">
            <Tabs.List>
                <Tabs.Trigger value="instructions">
                    <FiUser />
                    {t('bot-tabs.instructions')}
                </Tabs.Trigger>
                <Tabs.Trigger value="products">
                    <FiShoppingCart />
                    {t('bot-tabs.products')}
                </Tabs.Trigger>
                <Tabs.Trigger value="actions">
                    <FiSettings />
                    {t('bot-tabs.actions')}
                </Tabs.Trigger>
            </Tabs.List>
            <Tabs.Content value="instructions">
                <BotParams botId={botId} />
            </Tabs.Content>
            <Tabs.Content value="actions">
                <BotActions botId={botId} />
            </Tabs.Content>
            <Tabs.Content value="products">
                <Stack gap={4}>
                    <ProductAdd botId={botId} />
                    <ProductList botId={botId} />
                </Stack>
            </Tabs.Content>
        </Tabs.Root>
    )
}
