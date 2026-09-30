import { Stack } from "@chakra-ui/react"
import { ProductEdit } from "@/components/product-edit"
import { useLocation, useNavigate } from "react-router-dom"

export const Product = () => {
    const { state } = useLocation() as { state?: { productId?: string } }
    const navigate = useNavigate()
    const productId = state?.productId

    if (!productId) {
        navigate('/')
        return null
    }

    return (
        <Stack gap={4}>
            <ProductEdit productId={productId} />
        </Stack>
    )
}

export default Product