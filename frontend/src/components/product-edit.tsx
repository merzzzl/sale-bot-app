import { Button, Portal, createListCollection, Field, Input, Select, Textarea, NumberInput, Stack, HStack, Fieldset, SegmentGroup } from "@chakra-ui/react"
import { PasswordInput } from "./ui/password-input"
import { useTranslation } from "react-i18next"
import { useEffect, useState, useMemo } from "react"
import { useQuery, useQueryClient } from "@tanstack/react-query"
import { useApi } from "@/lib/api"
import { useNavigate } from "react-router-dom"
import { ProductCurrencyEnum } from "@ltbots/api"

interface ProductEditProps {
    productId: string
}

export const ProductEdit = ({ productId }: ProductEditProps) => {
    const { t } = useTranslation()
    const { productApi } = useApi()
    const queryClient = useQueryClient()
    const navigate = useNavigate()

    const [currency, setCurrency] = useState<string[]>([])
    const [name, setName] = useState('')
    const [description, setDescription] = useState('')
    const [price, setPrice] = useState(0)
    const [payLink, setPayLink] = useState('')
    const [imageUrl, setImageUrl] = useState('')
    const [useInvoice, setUseInvoice] = useState(false)
    const [paymentToken, setPaymentToken] = useState('')

    const productQuery = useQuery({
        queryKey: ['product', productId],
        queryFn: () => productApi.getProduct(productId),
    })

    useEffect(() => {
        if (productQuery.data) {
            setCurrency([productQuery.data.currency])
            setName(productQuery.data.name)
            setDescription(productQuery.data.description)
            setPrice(Number(productQuery.data.price) / 100)
            setPayLink(productQuery.data.payLink)
            setImageUrl(productQuery.data.imageUrl)
            setUseInvoice(productQuery.data.useInvoice)
        }
    }, [productQuery.data])

    const handleSubmit = () => {
        productApi.updateProduct(productId, {
            name: name.trim(),
            description: description.trim(),
            price: (price * 100).toString(),
            payLink: payLink.trim(),
            imageUrl: imageUrl.trim(),
            currency: currency[0] as ProductCurrencyEnum,
            useInvoice: useInvoice || currency[0] === 'XTR',
            paymentToken: paymentToken.trim(),
        }).then(() => {
            queryClient.invalidateQueries({ queryKey: ['products', productQuery.data?.botId] })
            navigate(-1)
        })
    }

    if (productQuery.isLoading) {
        return <ProductEditSkeleton />
    }

    return (
        <Fieldset.Root size="lg" maxW="md">
            <Stack gap={4}>
                <Fieldset.Legend>
                    {t('product-edit.title')}
                </Fieldset.Legend>
                <Fieldset.HelperText>
                    {t('product-edit.description')}
                </Fieldset.HelperText>
            </Stack>    
            <Fieldset.Content>
                <Field.Root>
                    <Field.Label>{t('product-edit.name-label')}</Field.Label>
                    <Input value={name} onChange={(e) => setName(e.target.value)} />
                    <Field.HelperText textAlign="left">{t('product-edit.name-description')}</Field.HelperText>
                </Field.Root>
                <Field.Root>
                    <Field.Label>{t('product-edit.image-label')}</Field.Label>
                    <Input type="url" value={imageUrl} onChange={(e) => setImageUrl(e.target.value)} />
                    <Field.HelperText textAlign="left">{t('product-edit.image-description')}</Field.HelperText>
                </Field.Root>
                <Field.Root>
                    <Field.Label>{t('product-edit.description-label')}</Field.Label>
                    <Textarea value={description} onChange={(e) => setDescription(e.target.value)} />
                    <Field.HelperText textAlign="left">{t('product-edit.description-description')}</Field.HelperText>
                </Field.Root>
                <Field.Root>
                    <Field.Label>{t('product-edit.price-label')}</Field.Label>
                    <HStack w="full">
                        <NumberInput.Root w="3/4" asChild min={0}>
                            <NumberInput.Input value={price} onChange={(e) => setPrice(Number(e.target.value))} />
                        </NumberInput.Root>
                        <CurrencySelect value={currency[0]} onChange={(value) => setCurrency([value])} />
                    </HStack>
                    <Field.HelperText textAlign="left">{t('product-edit.price-description')}</Field.HelperText>
                </Field.Root>
                <Field.Root>
                    <Field.Label>{t('product-edit.use-invoice-label')}</Field.Label>
                    <SegmentGroup.Root size="sm" value={useInvoice || currency[0] === 'XTR' ? 'Telegram' : 'Link'} onValueChange={(e) => setUseInvoice(e.value === 'Telegram')}>
                        <SegmentGroup.Indicator />
                        <SegmentGroup.Items items={["Link", "Telegram"]} />
                    </SegmentGroup.Root>
                    <Field.HelperText textAlign="left">{t('product-edit.use-invoice-description')}</Field.HelperText>
                </Field.Root>
                <Field.Root hidden={useInvoice || currency[0] === 'XTR'}>
                    <Field.Label>{t('product-edit.pay-label')}</Field.Label>
                    <Input type="url" value={payLink} onChange={(e) => setPayLink(e.target.value)} />
                    <Field.HelperText textAlign="left">{t('product-edit.pay-description')}</Field.HelperText>
                </Field.Root>
                <Field.Root hidden={!useInvoice || currency[0] === 'XTR'}>
                    <Field.Label>{t('product-edit.payment-token-label')}</Field.Label>
                    <PasswordInput type="text" value={paymentToken} onChange={(e) => setPaymentToken(e.target.value)} />
                    <Field.HelperText textAlign="left">{t('product-edit.payment-token-description')}</Field.HelperText>
                </Field.Root>
            </Fieldset.Content>
            <Button colorPalette="blue" onClick={handleSubmit}>
                {t('product-edit.submit-button')}
            </Button>
        </Fieldset.Root>
    )
}

export const CurrencySelect = ({ value, onChange }: { value: string, onChange: (value: string) => void }) => {
    const currenciesCollection = useMemo(() => createListCollection({
        items: [
            { value: 'RUB', label: 'RUB' },
            { value: 'USD', label: 'USD' },
            { value: 'EUR', label: 'EUR' },
            { value: 'XTR', label: 'Stars' },
        ],
    }), [])

    return (
        <Select.Root w="1/4" collection={currenciesCollection} value={[value]} onValueChange={(value) => onChange(value.value[0])} >
            <Select.Control>
                <Select.Trigger>
                    <Select.ValueText />
                </Select.Trigger>
                <Select.IndicatorGroup>
                    <Select.Indicator />
                </Select.IndicatorGroup>
            </Select.Control>
            <Portal>
                <Select.Positioner>
                    <Select.Content>
                        {currenciesCollection.items.map((item) => (
                            <Select.Item item={item} key={item.value}>
                                {item.label}
                                <Select.ItemIndicator />
                            </Select.Item>
                        ))}
                    </Select.Content>
                </Select.Positioner>
            </Portal>
        </Select.Root>
    )
}

export const ProductEditSkeleton = () => {
    const { t } = useTranslation()

    return (
        <Fieldset.Root size="lg" maxW="md">
            <Stack gap={4}>
                <Fieldset.Legend>
                    {t('product-edit.title')}
                </Fieldset.Legend>
                <Fieldset.HelperText>
                    {t('product-edit.description')}
                </Fieldset.HelperText>
            </Stack>
            <Fieldset.Content>
                <Field.Root>
                    <Field.Label>{t('product-edit.name-label')}</Field.Label>
                    <Input disabled />
                    <Field.HelperText textAlign="left">{t('product-edit.name-description')}</Field.HelperText>
                </Field.Root>
                <Field.Root>
                    <Field.Label>{t('product-edit.image-label')}</Field.Label>
                    <Input type="url" disabled />
                    <Field.HelperText textAlign="left">{t('product-edit.image-description')}</Field.HelperText>
                </Field.Root>
                <Field.Root>
                    <Field.Label>{t('product-edit.description-label')}</Field.Label>
                    <Textarea disabled />
                    <Field.HelperText textAlign="left">{t('product-edit.description-description')}</Field.HelperText>
                </Field.Root>
                <Field.Root>
                    <Field.Label>{t('product-edit.price-label')}</Field.Label>
                    <HStack w="full">
                        <NumberInput.Root min={0} w="3/4">
                            <NumberInput.Control />
                            <NumberInput.Input disabled />
                        </NumberInput.Root>
                        <Input type="text" w="1/4" disabled />
                    </HStack>
                </Field.Root>
                <Field.Root>
                    <Field.Label>{t('product-edit.pay-label')}</Field.Label>
                    <Input type="url" disabled />
                    <Field.HelperText textAlign="left">{t('product-edit.pay-description')}</Field.HelperText>
                </Field.Root>
            </Fieldset.Content>
            <Button colorPalette="blue" loading>
                {t('product-edit.submit-button')}
            </Button>
        </Fieldset.Root>
    )
}