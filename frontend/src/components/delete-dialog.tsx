import { 
    Dialog, 
    Portal,
    Button, 
    Text
} from "@chakra-ui/react"
import { FiTrash } from "react-icons/fi"
import { useTranslation } from "react-i18next"
import { useRef, useImperativeHandle, useEffect, forwardRef } from "react"

export type DeleteDialogRef = {
    setOnApprove: (onApprove: () => void) => void
    setOnDiscard: (onDiscard: () => void) => void
}

export interface DeleteDialogProps {
    open?: boolean
    description: string
    onApprove?: () => void
    onDiscard?: () => void
}

export const DeleteDialog = forwardRef<DeleteDialogRef, DeleteDialogProps>(function DeleteDialog({ open, description, onApprove, onDiscard }, ref) {
    const { t } = useTranslation()

    const approveRef = useRef<() => void>(() => { })
    const discardRef = useRef<() => void>(() => { })

    useImperativeHandle(ref, () => ({
        setOnApprove: (fn: () => void) => { approveRef.current = fn },
        setOnDiscard: (fn: () => void) => { discardRef.current = fn },
    }), [])

    useEffect(() => {
        if (onApprove) approveRef.current = onApprove
        if (onDiscard) discardRef.current = onDiscard
    }, [onApprove, onDiscard])

    return (
        <Dialog.Root open={open} placement="center">
            <Portal>
                <Dialog.Positioner background="blackAlpha.600">
                    <Dialog.Content w="90vw" maxW="420px">
                        <Dialog.Header>
                            <Dialog.Title>{t('delete-dialog.title')}</Dialog.Title>
                        </Dialog.Header>
                        <Dialog.Body>
                            <Text>{description}</Text>
                        </Dialog.Body>
                        <Dialog.Footer>
                            <Button colorPalette="red" onClick={() => approveRef.current()}>
                                <FiTrash />
                                {t('delete-dialog.approve')}
                            </Button>
                            <Dialog.ActionTrigger asChild>
                                <Button variant="outline" onClick={() => discardRef.current()}>
                                    {t('delete-dialog.discard')}
                                </Button>
                            </Dialog.ActionTrigger>
                        </Dialog.Footer>
                    </Dialog.Content>
                </Dialog.Positioner>
            </Portal>
        </Dialog.Root>
    )
})
