export const timeNow = (): number => {
    return Math.floor(Date.now() / 1000)
}

export const timeParse = (date: number): Date => {
    return new Date(date * 1000)
}

export const timeFormat = (date: Date): number => {
    return Math.floor(date.getTime() / 1000)
}

export const timeFrom = (params: { seconds?: number, minutes?: number, hours?: number, days?: number }): number => {
    const { seconds, minutes, hours, days } = params
    return (seconds ?? 0) + (minutes ?? 0) * 60 + (hours ?? 0) * 60 * 60 + (days ?? 0) * 60 * 60 * 24
}
