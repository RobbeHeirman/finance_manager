# AuthApi

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**authGoogleAuthPost**](#authgoogleauthpost) | **POST** /auth/google_auth | Authenticate using a Google token|

# **authGoogleAuthPost**
> RestUserResponse authGoogleAuthPost(request)

Exchanges a Google OAuth token for an app-specific JWT

### Example

```typescript
import {
    AuthApi,
    Configuration,
    RestTokenRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new AuthApi(configuration);

let request: RestTokenRequest; //The google token request. Probably received from google oAuth

const { status, data } = await apiInstance.authGoogleAuthPost(
    request
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **request** | **RestTokenRequest**| The google token request. Probably received from google oAuth | |


### Return type

**RestUserResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

