import { Stack } from '@chakra-ui/react'
import { ColorModeButton } from '@/components/ui/color-mode'
import { BotRegister } from '@/components/bot-register'
import { BotList } from '@/components/bot-list'
import { UserAvatar } from '@/components/user-avatar'
import { UserBalance } from '@/components/user-balance'

export const Main = () => {
    return (
        <Stack gap={4}>
            <UserBalance position="absolute" top={4} left={4} />
            <ColorModeButton position="absolute" top={4} right={4}/>
            <UserAvatar />
            <BotRegister />
            <BotList />
        </Stack>
    )
}

export default Main