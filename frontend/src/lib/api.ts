import { useRawInitData } from '@tma.js/sdk-react';
import {
    Configuration,
    ControllerApi
} from '@ltbots/api'
import type {
    BotCreateRequest,
    Bot,
    BotUpdateRequest,
    ProductCreateRequest,
    ProductUpdateRequest,
    Product,
    PromptPresetListResponsePromptPresetsInner as PromptPreset,
    StatisticsGetResponseRecordsInner as StatisticsRecord,
    TransactionsListResponseTransactionsInner as Transaction,
} from '@ltbots/api'

const API_BASE_URL = ''

export type {
    Bot,
    Product,
    PromptPreset,
    Transaction,
}

const createConfiguration = (initDataRaw: string | undefined) => {
    return new Configuration({
        basePath: API_BASE_URL,
        baseOptions: {
            headers: {
                'Content-Type': 'application/json',
                'tg-token': initDataRaw,
            },
        },
    })
}

const createApiClient = (initDataRaw: string | undefined) => {
    return new ControllerApi(createConfiguration(initDataRaw))
}

export class ApiError extends Error {
    status?: number
    response?: any

    constructor(message: string, status?: number, response?: any) {
        super(message)
        this.name = 'ApiError'
        this.status = status
        this.response = response
    }
}

const handleApiError = (error: any): never => {
    if (error.response) {
        throw new ApiError(
            error.response.data?.message || 'Server error',
            error.response.status,
            error.response.data
        )
    } else if (error.request) {
        throw new ApiError('No connection to server')
    } else {
        throw new ApiError(error.message || 'Unknown error')
    }
}

export const useApi = () => {
    const initDataRaw = useRawInitData()

    const settingsApi = {
        async getPromptPresets() {
            try {
                const apiClient = createApiClient(initDataRaw)
                const response = await apiClient.controllerPromptPresetList()
                return response.data as PromptPreset[]
            } catch (error) {
                handleApiError(error)
            }
        },
    }

    const transactionsApi = {
        async getTransactions() {
            try {
                const apiClient = createApiClient(initDataRaw)
                const response = await apiClient.controllerTransactionsList()
                return response.data as Transaction[]
            } catch (error) {
                handleApiError(error)
            }
        },

        async createTransaction(amount: string) {
            try {
                const apiClient = createApiClient(initDataRaw)
                await apiClient.controllerTransactionBill({amount: amount})
            } catch (error) {
                handleApiError(error)
            }
        },
    }

    const botApi = {
        async getBotStatistics(botId: string, startTime: string, endTime: string) {
            try {
                const apiClient = createApiClient(initDataRaw)
                const response = await apiClient.controllerStatisticsGet(botId, startTime, endTime)
                return response.data as StatisticsRecord[]
            } catch (error) {
                handleApiError(error)
            }
        },
        
        async getBot(botId: string) {
            try {
                const apiClient = createApiClient(initDataRaw)
                const response = await apiClient.controllerBotGet(botId)
                return response.data as Bot
            } catch (error) {
                handleApiError(error)
            }
        },

        async getBots() {
            try {
                const apiClient = createApiClient(initDataRaw)
                const response = await apiClient.controllerBotList()
                return response.data as Bot[]
            } catch (error) {
                handleApiError(error)
            }
        },

        async createBot(data: BotCreateRequest) {
            try {
                const apiClient = createApiClient(initDataRaw)
                const response = await apiClient.controllerBotCreate(data)
                return response.data as Bot
            } catch (error) {
                handleApiError(error)
            }
        },

        async updateBot(botId: string, data: BotUpdateRequest) {
            try {
                const apiClient = createApiClient(initDataRaw)
                const response = await apiClient.controllerBotUpdate(botId, data)
                return response.data as Bot
            } catch (error) {
                handleApiError(error)
            }
        },

        async deleteBot(botId: string) {
            try {
                const apiClient = createApiClient(initDataRaw)
                await apiClient.controllerBotDelete(botId)
            } catch (error) {
                handleApiError(error)
            }
        },

        async activateBot(botId: string) {
            try {
                const apiClient = createApiClient(initDataRaw)
                await apiClient.controllerBotActivate(botId)
            } catch (error) {
                handleApiError(error)
            }
        },

        async deactivateBot(botId: string) {
            try {
                const apiClient = createApiClient(initDataRaw)
                await apiClient.controllerBotDeactivate(botId)
            } catch (error) {
                handleApiError(error)
            }
        },
    }

    const productApi = {
        async getProduct(productId: string) {
            try {
                const apiClient = createApiClient(initDataRaw)
                const response = await apiClient.controllerProductGet(productId)
                return response.data as Product
            } catch (error) {
                handleApiError(error)
            }
        },

        async getProducts(botId: string) {
            try {
                const apiClient = createApiClient(initDataRaw)
                const response = await apiClient.controllerProductList(botId)
                return response.data as Product[]
            } catch (error) {
                handleApiError(error)
            }
        },

        async createProduct(data: ProductCreateRequest) {
            try {
                const apiClient = createApiClient(initDataRaw)
                const response = await apiClient.controllerProductCreate(data)
                return response.data as Product
            } catch (error) {
                handleApiError(error)
            }
        },

        async updateProduct(productId: string, data: ProductUpdateRequest) {
            try {
                const apiClient = createApiClient(initDataRaw)
                const response = await apiClient.controllerProductUpdate(productId, data)
                return response.data as Product
            } catch (error) {
                handleApiError(error)
            }
        },

        async deleteProduct(productId: string) {
            try {
                const apiClient = createApiClient(initDataRaw)
                await apiClient.controllerProductDelete(productId)
            } catch (error) {
                handleApiError(error)
            }
        },
    }

    return { botApi, productApi, settingsApi, transactionsApi }
}
