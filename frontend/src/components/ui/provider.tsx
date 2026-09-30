"use client"

import { ChakraProvider, createSystem, defaultConfig, defineConfig } from "@chakra-ui/react"
import { ColorModeProvider, type ColorModeProviderProps } from "./color-mode"

export function Provider(props: ColorModeProviderProps) {
  const config = defineConfig({
    theme: {
      semanticTokens: {
        colors: {
          app: {
            bg: {
              DEFAULT: {
                value: {
                  _light: "{colors.white}",
                  _dark: "{colors.black}",
                },
              },
            },
          },
        },
      },
    },
    globalCss: {
      "html, body": {
        bg: "{colors.app.bg}",
      },
    },
  })


  const system = createSystem(defaultConfig, config)

  return (
    <ChakraProvider value={system}>
      <ColorModeProvider {...props} />
    </ChakraProvider>
  )
}
