import { useLaunchParams, type User } from "@tma.js/sdk-react"
import { AvatarHeader } from "./ui/avatar-header"
import { useTranslation } from "react-i18next"

export const UserAvatar = () => {
    const { t } = useTranslation()
    
    const launchParams = useLaunchParams()
    const user: User = launchParams.tgWebAppData!.user!

    let name = user.first_name + " " + user.last_name
    const username = user.username || t('user-avatar.error-username')

    if (name.trim() === "") {
        name = t('user-avatar.error-name')
    }
    
    return (
        <AvatarHeader
            photo_url={user.photo_url}
            name={name}
            username={username}
            active={true}
            showActive={false}
        />
    )
}
