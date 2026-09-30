import { 
    type ColorPalette,
    type StackProps,
    IconButton, 
    Avatar, 
    Stack, 
    HStack, 
    Text,
    Badge,
    Menu,
    Portal,
    Box,
    SkeletonCircle,
    SkeletonText,
    Skeleton 
} from "@chakra-ui/react"
import { FiMoreVertical } from "react-icons/fi"
import React from "react"

export type HorizontalCardProps = StackProps & {
    isLoading?: boolean
    onClick: () => void
    img: string
    title: string
    description: string
    badges?: {
        notification: string
        colorPalette: ColorPalette
    }[]
    actions?: {
        icon: React.ReactNode
        label: string
        onClick: () => void
    }[]
}

export const HorizontalCard = ({ isLoading, onClick, img, title, description, badges, actions, ...props }: HorizontalCardProps) => {
    if (isLoading) {
        return <HorizontalCardSkeleton />
    }

    return (
        <HStack justify="flex-start" align="center" p={2} cursor="pointer" {...props} onClick={() => onClick?.()}>
            <Avatar.Root size="xl" bg="bg.subtle">
                <Avatar.Image src={img} />
            </Avatar.Root>
            <Stack gap={0} justify="flex-start" align="flex-start" pl={2} flex="1" minW="0">
                <Stack direction="row" align="center" gap={2} maxW="100%">
                    <Text color="fg" fontWeight="bold" truncate textAlign="left">
                        {title}
                    </Text>
                    {badges?.map((badge, i) => (
                        <Badge key={i} colorPalette={badge.colorPalette}>
                            {badge.notification}
                        </Badge>
                    ))}
                </Stack>
                <Text color={description ? 'fg.subtle' : 'transparent'} truncate maxW="100%" textAlign="left">
                    {description ? description : '-'}
                </Text>
            </Stack>
            <Menu.Root>
                <Menu.Trigger asChild>
                    <IconButton variant="ghost" size="sm" onClick={(e) => { e.stopPropagation() }}>
                        <FiMoreVertical />
                    </IconButton>
                </Menu.Trigger>
                <Portal>
                    <Menu.Positioner>
                        <Menu.Content>
                            {actions?.map((action, i) => (
                                <Menu.Item key={i} value={action.label} onClick={(e) => { e.stopPropagation(); action.onClick() }}>
                                    {action.icon}
                                    <Box flex="1">{action.label}</Box>
                                </Menu.Item>
                            ))}
                        </Menu.Content>
                    </Menu.Positioner>
                </Portal>
            </Menu.Root>
        </HStack>
    )
}

export const HorizontalCardSkeleton = ({ withBadges, withActions }: { withBadges?: boolean, withActions?: boolean }) => {
    return (
        <HStack justify="flex-start" align="center" p={2} cursor="pointer">
            <SkeletonCircle>
                <Avatar.Root size="xl" />
            </SkeletonCircle>
            <Stack gap={1} justify="flex-start" align="flex-start" pl={2} flex="1" minW="0">
                <Stack direction="row" align="center" gap={2}>
                    <SkeletonText noOfLines={1} width="120px" />
                    {withBadges && <SkeletonText noOfLines={1} width="80px" />}
                </Stack>
                <SkeletonText noOfLines={1} width="120px" />    
            </Stack>
            {withActions && <Skeleton>
                <IconButton size="sm">
                    <FiMoreVertical />
                </IconButton>
            </Skeleton>}
        </HStack>
    )
}
