import { Bot } from "grammy"

export interface BotDetails {
    name: string
    username: string
    photo_url: string
}

export const useTgApi = () => {
    const tgApi = {
        async getBotDetails(token: string) {
            const info: BotDetails = {
                name: "",
                username: "",
                photo_url: "",
            }

            if (!token) return info

            const bot = new Bot(token, { client: { apiRoot: "/tg" } })
            const me = await bot.api.getMe()

            info.name = `${me.first_name || ""} ${me.last_name || ""}`.trim() || (me.username ? `@${me.username}` : "")
            info.username = me.username || ""

            const photos = await bot.api.getUserProfilePhotos(me.id, { limit: 1 })
            const first = photos.photos?.[0]?.[0]
            if (first?.file_id) {
                const file = await bot.api.getFile(first.file_id)
                if (file.file_path) {
                    info.photo_url = `/tg/file/bot${token}/${file.file_path}`
                }
            }

            return info
        }
    }

    return { tgApi }
}