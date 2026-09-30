import { useMemo, useRef, useState } from "react"
import { useQuery, useQueryClient } from "@tanstack/react-query"
import { useTranslation } from "react-i18next"
import { useApi, type Product } from "@/lib/api"
import { HorizontalCard, HorizontalCardSkeleton } from "./ui/horizontal-card"
import { Stack } from "@chakra-ui/react"
import { DeleteDialog, type DeleteDialogRef } from "./delete-dialog"
import { FiTrash } from "react-icons/fi"
import { useNavigate } from "react-router-dom"

export interface ProductListProps {
    botId: string
}

export const ProductList = ({ botId }: ProductListProps) => {
    const { t } = useTranslation()
    const { productApi } = useApi()
    const queryClient = useQueryClient()
    const navigate = useNavigate()

    const deleteDialog = useRef<DeleteDialogRef>(null)
    const [deleteDialogOpen, setDeleteDialogOpen] = useState(false)

    const products = useQuery({
        queryKey: ['products', botId],
        queryFn: () => productApi.getProducts(botId),
    })

    const onDelete = (productId: string) => {
        productApi.deleteProduct(productId).then(() => {
            queryClient.invalidateQueries({ queryKey: ['products', botId] })
        })
    }

    const onDeleteDialog = (productId: string) => {
        deleteDialog.current?.setOnApprove(() => {
            onDelete(productId)
            setDeleteDialogOpen(false)
        })
        deleteDialog.current?.setOnDiscard(() => setDeleteDialogOpen(false))
        setDeleteDialogOpen(true)
    }

    const onClick = (productId: string) => {
        navigate('/product', { state: { productId } })
    }

    const productItems = useMemo(() => {
        return products.data?.map((product: Product) => (
            <HorizontalCard
                key={product.productId}
                onClick={() => onClick(product.productId)}
                img={product.imageUrl}
                title={product.name}
                description={product.description}
                badges={[{
                    notification: (Number(product.price) / 100).toString() + " " + product.currency,
                    colorPalette: "blue",
                }]}
                actions={[{
                    icon: <FiTrash />,
                    label: t('product-list.delete'),
                    onClick: () => onDeleteDialog(product.productId),
                }]}
            />
        ))
    }, [products.data])

    return (
        <Stack gap={4} >
            {products.isLoading ? (
                <Stack gap={2} borderRadius={8} backgroundColor="bg.subtle" p={0}>
                    {[...Array(3)].map((_, i) => (
                        <HorizontalCardSkeleton key={i} withActions withBadges />
                    ))}
                </Stack>
            ) : (
                <Stack gap={2} backgroundColor="bg.subtle" borderRadius="md">
                    {productItems}
                </Stack>
            )}
            <DeleteDialog
                ref={deleteDialog}
                open={deleteDialogOpen}
                description={t('product-list.delete-description')}
            />
        </Stack>
    )
}
