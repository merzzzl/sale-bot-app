import { useEffect } from "react";
import { miniApp } from "@tma.js/sdk-react";
import { useColorMode } from "./color-mode";

export function TelegramChromeSync() {
    const { colorMode } = useColorMode();

    useEffect(() => {
        const key = colorMode === "dark" ? "#0f0f0f" : "#f0f0f0";

        if (miniApp.setHeaderColor.isAvailable()) {
            miniApp.setHeaderColor(key);
        }
        if (miniApp.setBgColor.isAvailable()) {
            miniApp.setBgColor(key);
        }
        if (miniApp.setBottomBarColor.isAvailable()) {
            miniApp.setBottomBarColor(key);
        }
    }, [colorMode]);

    return null;
}
