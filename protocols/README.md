## Controller

### [GET] /api/v1/bots
**List all telegram bots on platform**

#### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 |  | **application/json**: [BotListResponse.Bots](#botlistresponsebots)<br> |
| default |  |  |

### [POST] /api/v1/bots
**Register telegram bot on platform**

#### Request Body

| Required | Schema |
| -------- | ------ |
|  Yes | **application/json**: [BotCreateRequest](#botcreaterequest)<br> |

#### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 |  | **application/json**: [Bot](#bot)<br> |
| default |  |  |

### [DELETE] /api/v1/bots/{bot_id}
**Delete telegram bot from platform**

#### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| bot_id | path |  | Yes | string (int64) |

#### Responses

| Code | Description |
| ---- | ----------- |
| 204 |  |
| default |  |

### [GET] /api/v1/bots/{bot_id}
**Get telegram bot on platform**

#### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| bot_id | path |  | Yes | string (int64) |

#### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 |  | **application/json**: [Bot](#bot)<br> |
| default |  |  |

### [PUT] /api/v1/bots/{bot_id}
**Update telegram bot on platform**

#### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| bot_id | path |  | Yes | string (int64) |

#### Request Body

| Required | Schema |
| -------- | ------ |
|  Yes | **application/json**: [BotUpdateRequest](#botupdaterequest)<br> |

#### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 |  | **application/json**: [Bot](#bot)<br> |
| default |  |  |

### [POST] /api/v1/bots/{bot_id}/activate
**Activate telegram bot**

#### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| bot_id | path |  | Yes | string (int64) |

#### Responses

| Code | Description |
| ---- | ----------- |
| 204 |  |
| default |  |

### [POST] /api/v1/bots/{bot_id}/deactivate
**Deactivate telegram bot**

#### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| bot_id | path |  | Yes | string (int64) |

#### Responses

| Code | Description |
| ---- | ----------- |
| 204 |  |
| default |  |

### [GET] /api/v1/products
**List all products for telegram bot**

#### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| bot_id | query |  | Yes | string (int64) |

#### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 |  | **application/json**: [ProductListResponse.Products](#productlistresponseproducts)<br> |
| default |  |  |

### [POST] /api/v1/products
**Create product for telegram bot**

#### Request Body

| Required | Schema |
| -------- | ------ |
|  Yes | **application/json**: [ProductCreateRequest](#productcreaterequest)<br> |

#### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 |  | **application/json**: [Product](#product)<br> |
| default |  |  |

### [DELETE] /api/v1/products/{product_id}
**Delete product for telegram bot**

#### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| product_id | path |  | Yes | string (int64) |

#### Responses

| Code | Description |
| ---- | ----------- |
| 204 |  |
| default |  |

### [GET] /api/v1/products/{product_id}
**Get product for telegram bot**

#### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| product_id | path |  | Yes | string (int64) |

#### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 |  | **application/json**: [Product](#product)<br> |
| default |  |  |

### [PUT] /api/v1/products/{product_id}
**Update product for telegram bot**

#### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| product_id | path |  | Yes | string (int64) |

#### Request Body

| Required | Schema |
| -------- | ------ |
|  Yes | **application/json**: [ProductUpdateRequest](#productupdaterequest)<br> |

#### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 |  | **application/json**: [Product](#product)<br> |
| default |  |  |

### [GET] /api/v1/settings/prompt-presets
**List of promt presets for telegram bot**

#### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 |  | **application/json**: [PromptPresetListResponse.PromptPresets](#promptpresetlistresponsepromptpresets)<br> |
| default |  |  |

### [GET] /api/v1/statistics
**Get statistics for user**

#### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| bot_id | query |  | Yes | string (int64) |
| start_time | query |  | Yes | string (int64) |
| end_time | query |  | Yes | string (int64) |

#### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 |  | **application/json**: [StatisticsGetResponse.Records](#statisticsgetresponserecords)<br> |
| default |  |  |

### [GET] /api/v1/transactions
**Get transactions**

#### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 |  | **application/json**: [TransactionsListResponse.Transactions](#transactionslistresponsetransactions)<br> |
| default |  |  |

### [POST] /api/v1/transactions/bill
**Deposit transaction**

#### Request Body

| Required | Schema |
| -------- | ------ |
|  Yes | **application/json**: [TransactionBillRequest](#transactionbillrequest)<br> |

#### Responses

| Code | Description |
| ---- | ----------- |
| 204 |  |
| default |  |

---
### Schemas

#### Bot

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| active | boolean |  | Yes |
| botId | string (int64) |  | Yes |
| presetId | string (int64) |  | Yes |
| prompt | string |  | Yes |
| token | string |  | Yes |

#### BotCreateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| botToken | string |  | Yes |

#### BotListResponse.Bots

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| BotListResponse.Bots | array |  |  |

#### BotUpdateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| presetId | string (int64) |  | Yes |
| prompt | string |  | Yes |

#### Product

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| botId | string (int64) |  | Yes |
| currency | string, <br>**Available values:** "RUB", "USD", "EUR", "XTR" | *Enum:* `"RUB"`, `"USD"`, `"EUR"`, `"XTR"` | Yes |
| description | string |  | Yes |
| imageUrl | string |  | Yes |
| name | string |  | Yes |
| payLink | string |  | Yes |
| price | string (int64) |  | Yes |
| productId | string (int64) |  | Yes |
| useInvoice | boolean |  | Yes |

#### ProductCreateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| botId | string (int64) |  | Yes |
| currency | string, <br>**Available values:** "RUB", "USD", "EUR", "XTR" | *Enum:* `"RUB"`, `"USD"`, `"EUR"`, `"XTR"` | Yes |
| description | string |  | Yes |
| imageUrl | string |  | Yes |
| name | string |  | Yes |
| payLink | string |  | Yes |
| paymentToken | string |  | Yes |
| price | string (int64) |  | Yes |
| useInvoice | boolean |  | Yes |

#### ProductListResponse.Products

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| ProductListResponse.Products | array |  |  |

#### ProductUpdateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| currency | string, <br>**Available values:** "RUB", "USD", "EUR", "XTR" | *Enum:* `"RUB"`, `"USD"`, `"EUR"`, `"XTR"` | Yes |
| description | string |  | Yes |
| imageUrl | string |  | Yes |
| name | string |  | Yes |
| payLink | string |  | Yes |
| paymentToken | string |  | Yes |
| price | string (int64) |  | Yes |
| useInvoice | boolean |  | Yes |

#### PromptPresetListResponse.PromptPresets

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| PromptPresetListResponse.PromptPresets | array |  |  |

#### StatisticsGetResponse.Records

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| StatisticsGetResponse.Records | array |  |  |

#### TransactionBillRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| amount | string (int64) |  | Yes |

#### TransactionsListResponse.Transactions

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| TransactionsListResponse.Transactions | array |  |  |
