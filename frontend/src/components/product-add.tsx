import { useQueryClient } from "@tanstack/react-query"
import { useTranslation } from "react-i18next"
import { useApi, type Product } from "@/lib/api"
import { Button } from "@chakra-ui/react"
import { useNavigate } from "react-router-dom"
import { FiPlus } from "react-icons/fi"
import { ProductCurrencyEnum } from "@ltbots/api"

export interface ProductAddProps {
    botId: string
}

export const ProductAdd = ({ botId }: ProductAddProps) => {
    const { t } = useTranslation()
    const { productApi } = useApi()
    const queryClient = useQueryClient()
    const navigate = useNavigate()

    const onClick = () => {
        productApi.createProduct({
            name: t('product-add.example-name'),
            description: '',
            price: '0',
            payLink: '',
            imageUrl: '',
            botId: botId,
            currency: 'XTR' as ProductCurrencyEnum,
            useInvoice: true,
            paymentToken: '',
        }).then((newProduct : Product | undefined) => {
            queryClient.invalidateQueries({ queryKey: ['products', botId] })

            if (!newProduct) {
                return
            }

            navigate('/product', { state: { productId: newProduct.productId } })
        })
    }

    return (
        <Button onClick={onClick} colorPalette="blue" w="full">
            <FiPlus />
            {t('product-add.add')}
        </Button>
    )
}
