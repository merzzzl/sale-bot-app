# ControllerApi

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**controllerBotActivate**](#controllerbotactivate) | **POST** /api/v1/bots/{bot_id}/activate | Activate telegram bot|
|[**controllerBotCreate**](#controllerbotcreate) | **POST** /api/v1/bots | Register telegram bot on platform|
|[**controllerBotDeactivate**](#controllerbotdeactivate) | **POST** /api/v1/bots/{bot_id}/deactivate | Deactivate telegram bot|
|[**controllerBotDelete**](#controllerbotdelete) | **DELETE** /api/v1/bots/{bot_id} | Delete telegram bot from platform|
|[**controllerBotGet**](#controllerbotget) | **GET** /api/v1/bots/{bot_id} | Get telegram bot on platform|
|[**controllerBotList**](#controllerbotlist) | **GET** /api/v1/bots | List all telegram bots on platform|
|[**controllerBotUpdate**](#controllerbotupdate) | **PUT** /api/v1/bots/{bot_id} | Update telegram bot on platform|
|[**controllerProductCreate**](#controllerproductcreate) | **POST** /api/v1/products | Create product for telegram bot|
|[**controllerProductDelete**](#controllerproductdelete) | **DELETE** /api/v1/products/{product_id} | Delete product for telegram bot|
|[**controllerProductGet**](#controllerproductget) | **GET** /api/v1/products/{product_id} | Get product for telegram bot|
|[**controllerProductList**](#controllerproductlist) | **GET** /api/v1/products | List all products for telegram bot|
|[**controllerProductUpdate**](#controllerproductupdate) | **PUT** /api/v1/products/{product_id} | Update product for telegram bot|
|[**controllerPromptPresetList**](#controllerpromptpresetlist) | **GET** /api/v1/settings/prompt-presets | List of promt presets for telegram bot|
|[**controllerStatisticsGet**](#controllerstatisticsget) | **GET** /api/v1/statistics | Get statistics for user|
|[**controllerTransactionBill**](#controllertransactionbill) | **POST** /api/v1/transactions/bill | Deposit transaction|
|[**controllerTransactionsList**](#controllertransactionslist) | **GET** /api/v1/transactions | Get transactions|

# **controllerBotActivate**
> controllerBotActivate()


### Example

```typescript
import {
    ControllerApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ControllerApi(configuration);

let botId: string; // (default to undefined)

const { status, data } = await apiInstance.controllerBotActivate(
    botId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **botId** | [**string**] |  | defaults to undefined|


### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**204** |  |  -  |
|**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **controllerBotCreate**
> Bot controllerBotCreate(botCreateRequest)


### Example

```typescript
import {
    ControllerApi,
    Configuration,
    BotCreateRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new ControllerApi(configuration);

let botCreateRequest: BotCreateRequest; //

const { status, data } = await apiInstance.controllerBotCreate(
    botCreateRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **botCreateRequest** | **BotCreateRequest**|  | |


### Return type

**Bot**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** |  |  -  |
|**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **controllerBotDeactivate**
> controllerBotDeactivate()


### Example

```typescript
import {
    ControllerApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ControllerApi(configuration);

let botId: string; // (default to undefined)

const { status, data } = await apiInstance.controllerBotDeactivate(
    botId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **botId** | [**string**] |  | defaults to undefined|


### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**204** |  |  -  |
|**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **controllerBotDelete**
> controllerBotDelete()


### Example

```typescript
import {
    ControllerApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ControllerApi(configuration);

let botId: string; // (default to undefined)

const { status, data } = await apiInstance.controllerBotDelete(
    botId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **botId** | [**string**] |  | defaults to undefined|


### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**204** |  |  -  |
|**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **controllerBotGet**
> Bot controllerBotGet()


### Example

```typescript
import {
    ControllerApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ControllerApi(configuration);

let botId: string; // (default to undefined)

const { status, data } = await apiInstance.controllerBotGet(
    botId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **botId** | [**string**] |  | defaults to undefined|


### Return type

**Bot**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** |  |  -  |
|**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **controllerBotList**
> Array<BotListResponseBotsInner> controllerBotList()


### Example

```typescript
import {
    ControllerApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ControllerApi(configuration);

const { status, data } = await apiInstance.controllerBotList();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**Array<BotListResponseBotsInner>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** |  |  -  |
|**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **controllerBotUpdate**
> Bot controllerBotUpdate(botUpdateRequest)


### Example

```typescript
import {
    ControllerApi,
    Configuration,
    BotUpdateRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new ControllerApi(configuration);

let botId: string; // (default to undefined)
let botUpdateRequest: BotUpdateRequest; //

const { status, data } = await apiInstance.controllerBotUpdate(
    botId,
    botUpdateRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **botUpdateRequest** | **BotUpdateRequest**|  | |
| **botId** | [**string**] |  | defaults to undefined|


### Return type

**Bot**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** |  |  -  |
|**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **controllerProductCreate**
> Product controllerProductCreate(productCreateRequest)


### Example

```typescript
import {
    ControllerApi,
    Configuration,
    ProductCreateRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new ControllerApi(configuration);

let productCreateRequest: ProductCreateRequest; //

const { status, data } = await apiInstance.controllerProductCreate(
    productCreateRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **productCreateRequest** | **ProductCreateRequest**|  | |


### Return type

**Product**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** |  |  -  |
|**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **controllerProductDelete**
> controllerProductDelete()


### Example

```typescript
import {
    ControllerApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ControllerApi(configuration);

let productId: string; // (default to undefined)

const { status, data } = await apiInstance.controllerProductDelete(
    productId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **productId** | [**string**] |  | defaults to undefined|


### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**204** |  |  -  |
|**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **controllerProductGet**
> Product controllerProductGet()


### Example

```typescript
import {
    ControllerApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ControllerApi(configuration);

let productId: string; // (default to undefined)

const { status, data } = await apiInstance.controllerProductGet(
    productId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **productId** | [**string**] |  | defaults to undefined|


### Return type

**Product**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** |  |  -  |
|**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **controllerProductList**
> Array<ProductListResponseProductsInner> controllerProductList()


### Example

```typescript
import {
    ControllerApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ControllerApi(configuration);

let botId: string; // (default to undefined)

const { status, data } = await apiInstance.controllerProductList(
    botId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **botId** | [**string**] |  | defaults to undefined|


### Return type

**Array<ProductListResponseProductsInner>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** |  |  -  |
|**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **controllerProductUpdate**
> Product controllerProductUpdate(productUpdateRequest)


### Example

```typescript
import {
    ControllerApi,
    Configuration,
    ProductUpdateRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new ControllerApi(configuration);

let productId: string; // (default to undefined)
let productUpdateRequest: ProductUpdateRequest; //

const { status, data } = await apiInstance.controllerProductUpdate(
    productId,
    productUpdateRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **productUpdateRequest** | **ProductUpdateRequest**|  | |
| **productId** | [**string**] |  | defaults to undefined|


### Return type

**Product**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** |  |  -  |
|**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **controllerPromptPresetList**
> Array<PromptPresetListResponsePromptPresetsInner> controllerPromptPresetList()


### Example

```typescript
import {
    ControllerApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ControllerApi(configuration);

const { status, data } = await apiInstance.controllerPromptPresetList();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**Array<PromptPresetListResponsePromptPresetsInner>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** |  |  -  |
|**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **controllerStatisticsGet**
> Array<StatisticsGetResponseRecordsInner> controllerStatisticsGet()


### Example

```typescript
import {
    ControllerApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ControllerApi(configuration);

let botId: string; // (default to undefined)
let startTime: string; // (default to undefined)
let endTime: string; // (default to undefined)

const { status, data } = await apiInstance.controllerStatisticsGet(
    botId,
    startTime,
    endTime
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **botId** | [**string**] |  | defaults to undefined|
| **startTime** | [**string**] |  | defaults to undefined|
| **endTime** | [**string**] |  | defaults to undefined|


### Return type

**Array<StatisticsGetResponseRecordsInner>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** |  |  -  |
|**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **controllerTransactionBill**
> controllerTransactionBill(transactionBillRequest)


### Example

```typescript
import {
    ControllerApi,
    Configuration,
    TransactionBillRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new ControllerApi(configuration);

let transactionBillRequest: TransactionBillRequest; //

const { status, data } = await apiInstance.controllerTransactionBill(
    transactionBillRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **transactionBillRequest** | **TransactionBillRequest**|  | |


### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**204** |  |  -  |
|**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **controllerTransactionsList**
> Array<TransactionsListResponseTransactionsInner> controllerTransactionsList()


### Example

```typescript
import {
    ControllerApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new ControllerApi(configuration);

const { status, data } = await apiInstance.controllerTransactionsList();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**Array<TransactionsListResponseTransactionsInner>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** |  |  -  |
|**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

