import { Avatar, Float, Circle, Text, Stack, SkeletonCircle, SkeletonText } from "@chakra-ui/react"


export interface AvatarHeaderProps {
    isLoading?: boolean
    photo_url?: string
    name: string
    username: string
    showActive?: boolean
    active: boolean
}

export const AvatarHeader = ({ isLoading, photo_url, name, username, active, showActive = true }: AvatarHeaderProps) => {
    if (isLoading) {
        return <AvatarHeaderSkeleton />
    }

    return (
        <Stack gap={2} alignItems="center">
            <Avatar.Root size="2xl" border="3px solid" borderColor="bg.inverted">
                {photo_url && <Avatar.Image src={photo_url} />}
                {!photo_url && <Avatar.Fallback>{name.slice(0, 1)}</Avatar.Fallback>}
                <Float placement="bottom-end" offsetX="1" offsetY="1">
                    {showActive && <Circle
                        bg={active ? "border.success" : "border.error"}
                        size="14px"
                        outline="0.15em solid"
                        outlineColor="bg"
                    />}
                </Float>
            </Avatar.Root>
            <Stack gap={0}>
                <Text fontWeight="bold">
                    {name}
                </Text>
                <Text color="fg.subtle">
                    @{username}
                </Text>
            </Stack>
        </Stack>
    )
}

export const AvatarHeaderSkeleton = () => {
    return (
        <Stack gap={2} alignItems="center">
            <SkeletonCircle>
                <Avatar.Root size="2xl">
                    <Avatar.Fallback />
                </Avatar.Root>
            </SkeletonCircle>
            <Stack gap={1}>
                <SkeletonText noOfLines={1} width="120px" />
                <SkeletonText noOfLines={1} width="120px" />
            </Stack>
        </Stack>
    )
}
