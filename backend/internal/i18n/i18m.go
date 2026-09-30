package i18n

var translations = map[string]map[string]string{
	"en": {
		"send_invoice_title":       "Application balance top up",
		"send_invoice_description": "Your balance will be topped up by %d stars, they will be added to your balance in 5 minutes after payment, after that you will be able to use them to pay for messages in the bot.",
		"send_invoice_price_label": "Application balance top up",
		"notify_message":           "This chat will be used for notifications about new messages",
		"notify_success_message":   "Payment successfully processed",
		"prompt_language_message":  "Your answer must be in the english language if the user did not write in another language.",
		"handler_pay_button":       "Pay %d %s",
		"payment_notify_message":   "Payment successfully processed\n\nProduct: %s\nPrice: %d %s\n\nTelegram: @%s",
	},
	"ru": {
		"send_invoice_title":       "Пополнение баланса приложения",
		"send_invoice_description": "Ваш баланс будет пополнен на %d звезд, они будут добавлены на ваш баланс в течении 5 минут после оплаты, после этого вы сможете использовать их для оплаты сообщений в боте.",
		"send_invoice_price_label": "Пополнение баланса приложения",
		"notify_message":           "Этот чат будет использоваться для уведомлений о новых сообщениях",
		"notify_success_message":   "Платеж успешно обработан",
		"prompt_language_message":  "Твой ответ должен быть на русском языке если пользователь не писал на другом языке.",
		"handler_pay_button":       "Заплатить %d %s",
		"payment_notify_message":   "Платеж успешно обработан\n\nПродукт: %s\nЦена: %d %s\n\nTelegram: @%s",
	},
}

func Localize(languageCode string, key string) string {
	if _, ok := translations[languageCode]; !ok {
		languageCode = "en"
	}

	return translations[languageCode][key]
}
