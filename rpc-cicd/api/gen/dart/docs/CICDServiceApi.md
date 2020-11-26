# swagger.api.CICDServiceApi

## Load the API package
```dart
import 'package:swagger/api.dart';
```

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**cICDServiceDelTask**](CICDServiceApi.md#cICDServiceDelTask) | **DELETE** /v1/task/{id} | 
[**cICDServiceDelTemplate**](CICDServiceApi.md#cICDServiceDelTemplate) | **DELETE** /v1/template/{id} | 
[**cICDServiceDelVariable**](CICDServiceApi.md#cICDServiceDelVariable) | **DELETE** /v1/variable/{id} | 
[**cICDServiceGetTask**](CICDServiceApi.md#cICDServiceGetTask) | **GET** /v1/task/{id} | 
[**cICDServiceGetTemplate**](CICDServiceApi.md#cICDServiceGetTemplate) | **GET** /v1/template/{id} | 
[**cICDServiceGetVariable**](CICDServiceApi.md#cICDServiceGetVariable) | **GET** /v1/variable/{id} | 
[**cICDServiceListTask**](CICDServiceApi.md#cICDServiceListTask) | **GET** /v1/task | 
[**cICDServiceListTemplate**](CICDServiceApi.md#cICDServiceListTemplate) | **GET** /v1/template | 
[**cICDServiceListVariable**](CICDServiceApi.md#cICDServiceListVariable) | **GET** /v1/variable | 
[**cICDServicePutTask**](CICDServiceApi.md#cICDServicePutTask) | **POST** /v1/task | 
[**cICDServicePutTemplate**](CICDServiceApi.md#cICDServicePutTemplate) | **POST** /v1/template | 
[**cICDServicePutVariable**](CICDServiceApi.md#cICDServicePutVariable) | **POST** /v1/variable | 
[**cICDServiceRunTask**](CICDServiceApi.md#cICDServiceRunTask) | **POST** /v1/runTask | 
[**cICDServiceUpdateTask**](CICDServiceApi.md#cICDServiceUpdateTask) | **PUT** /v1/task/{task.id} | 
[**cICDServiceUpdateTemplate**](CICDServiceApi.md#cICDServiceUpdateTemplate) | **PUT** /v1/template/{template.id} | 
[**cICDServiceUpdateVariable**](CICDServiceApi.md#cICDServiceUpdateVariable) | **PUT** /v1/variable/{variable.id} | 


# **cICDServiceDelTask**
> ApiEmpty cICDServiceDelTask(id)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new CICDServiceApi();
var id = id_example; // String | 

try { 
    var result = api_instance.cICDServiceDelTask(id);
    print(result);
} catch (e) {
    print("Exception when calling CICDServiceApi->cICDServiceDelTask: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

### Return type

[**ApiEmpty**](ApiEmpty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceDelTemplate**
> ApiEmpty cICDServiceDelTemplate(id)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new CICDServiceApi();
var id = id_example; // String | 

try { 
    var result = api_instance.cICDServiceDelTemplate(id);
    print(result);
} catch (e) {
    print("Exception when calling CICDServiceApi->cICDServiceDelTemplate: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

### Return type

[**ApiEmpty**](ApiEmpty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceDelVariable**
> ApiEmpty cICDServiceDelVariable(id)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new CICDServiceApi();
var id = id_example; // String | 

try { 
    var result = api_instance.cICDServiceDelVariable(id);
    print(result);
} catch (e) {
    print("Exception when calling CICDServiceApi->cICDServiceDelVariable: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

### Return type

[**ApiEmpty**](ApiEmpty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceGetTask**
> ApiTask cICDServiceGetTask(id)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new CICDServiceApi();
var id = id_example; // String | 

try { 
    var result = api_instance.cICDServiceGetTask(id);
    print(result);
} catch (e) {
    print("Exception when calling CICDServiceApi->cICDServiceGetTask: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

### Return type

[**ApiTask**](ApiTask.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceGetTemplate**
> ApiTemplate cICDServiceGetTemplate(id)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new CICDServiceApi();
var id = id_example; // String | 

try { 
    var result = api_instance.cICDServiceGetTemplate(id);
    print(result);
} catch (e) {
    print("Exception when calling CICDServiceApi->cICDServiceGetTemplate: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

### Return type

[**ApiTemplate**](ApiTemplate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceGetVariable**
> ApiVariable cICDServiceGetVariable(id)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new CICDServiceApi();
var id = id_example; // String | 

try { 
    var result = api_instance.cICDServiceGetVariable(id);
    print(result);
} catch (e) {
    print("Exception when calling CICDServiceApi->cICDServiceGetVariable: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

### Return type

[**ApiVariable**](ApiVariable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceListTask**
> ApiListTaskRes cICDServiceListTask(offset, limit, brief)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new CICDServiceApi();
var offset = offset_example; // String | 
var limit = limit_example; // String | 
var brief = true; // bool | 

try { 
    var result = api_instance.cICDServiceListTask(offset, limit, brief);
    print(result);
} catch (e) {
    print("Exception when calling CICDServiceApi->cICDServiceListTask: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **String**|  | [optional] 
 **limit** | **String**|  | [optional] 
 **brief** | **bool**|  | [optional] 

### Return type

[**ApiListTaskRes**](ApiListTaskRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceListTemplate**
> ApiListTemplateRes cICDServiceListTemplate(offset, limit, brief)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new CICDServiceApi();
var offset = offset_example; // String | 
var limit = limit_example; // String | 
var brief = true; // bool | 

try { 
    var result = api_instance.cICDServiceListTemplate(offset, limit, brief);
    print(result);
} catch (e) {
    print("Exception when calling CICDServiceApi->cICDServiceListTemplate: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **String**|  | [optional] 
 **limit** | **String**|  | [optional] 
 **brief** | **bool**|  | [optional] 

### Return type

[**ApiListTemplateRes**](ApiListTemplateRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceListVariable**
> ApiListVariableRes cICDServiceListVariable(offset, limit, brief)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new CICDServiceApi();
var offset = offset_example; // String | 
var limit = limit_example; // String | 
var brief = true; // bool | 

try { 
    var result = api_instance.cICDServiceListVariable(offset, limit, brief);
    print(result);
} catch (e) {
    print("Exception when calling CICDServiceApi->cICDServiceListVariable: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **String**|  | [optional] 
 **limit** | **String**|  | [optional] 
 **brief** | **bool**|  | [optional] 

### Return type

[**ApiListVariableRes**](ApiListVariableRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServicePutTask**
> ApiEmpty cICDServicePutTask(body)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new CICDServiceApi();
var body = new ApiTask(); // ApiTask | 

try { 
    var result = api_instance.cICDServicePutTask(body);
    print(result);
} catch (e) {
    print("Exception when calling CICDServiceApi->cICDServicePutTask: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiTask**](ApiTask.md)|  | 

### Return type

[**ApiEmpty**](ApiEmpty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServicePutTemplate**
> ApiEmpty cICDServicePutTemplate(body)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new CICDServiceApi();
var body = new ApiTemplate(); // ApiTemplate | 

try { 
    var result = api_instance.cICDServicePutTemplate(body);
    print(result);
} catch (e) {
    print("Exception when calling CICDServiceApi->cICDServicePutTemplate: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiTemplate**](ApiTemplate.md)|  | 

### Return type

[**ApiEmpty**](ApiEmpty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServicePutVariable**
> ApiEmpty cICDServicePutVariable(body)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new CICDServiceApi();
var body = new ApiVariable(); // ApiVariable | 

try { 
    var result = api_instance.cICDServicePutVariable(body);
    print(result);
} catch (e) {
    print("Exception when calling CICDServiceApi->cICDServicePutVariable: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiVariable**](ApiVariable.md)|  | 

### Return type

[**ApiEmpty**](ApiEmpty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceRunTask**
> ApiRunTaskRes cICDServiceRunTask(body)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new CICDServiceApi();
var body = new ApiRunTaskReq(); // ApiRunTaskReq | 

try { 
    var result = api_instance.cICDServiceRunTask(body);
    print(result);
} catch (e) {
    print("Exception when calling CICDServiceApi->cICDServiceRunTask: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiRunTaskReq**](ApiRunTaskReq.md)|  | 

### Return type

[**ApiRunTaskRes**](ApiRunTaskRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceUpdateTask**
> ApiEmpty cICDServiceUpdateTask(taskId, body)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new CICDServiceApi();
var taskId = taskId_example; // String | 
var body = new ApiTask(); // ApiTask | 

try { 
    var result = api_instance.cICDServiceUpdateTask(taskId, body);
    print(result);
} catch (e) {
    print("Exception when calling CICDServiceApi->cICDServiceUpdateTask: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskId** | **String**|  | 
 **body** | [**ApiTask**](ApiTask.md)|  | 

### Return type

[**ApiEmpty**](ApiEmpty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceUpdateTemplate**
> ApiEmpty cICDServiceUpdateTemplate(templateId, body)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new CICDServiceApi();
var templateId = templateId_example; // String | 
var body = new ApiTemplate(); // ApiTemplate | 

try { 
    var result = api_instance.cICDServiceUpdateTemplate(templateId, body);
    print(result);
} catch (e) {
    print("Exception when calling CICDServiceApi->cICDServiceUpdateTemplate: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateId** | **String**|  | 
 **body** | [**ApiTemplate**](ApiTemplate.md)|  | 

### Return type

[**ApiEmpty**](ApiEmpty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceUpdateVariable**
> ApiEmpty cICDServiceUpdateVariable(variableId, body)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new CICDServiceApi();
var variableId = variableId_example; // String | 
var body = new ApiVariable(); // ApiVariable | 

try { 
    var result = api_instance.cICDServiceUpdateVariable(variableId, body);
    print(result);
} catch (e) {
    print("Exception when calling CICDServiceApi->cICDServiceUpdateVariable: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableId** | **String**|  | 
 **body** | [**ApiVariable**](ApiVariable.md)|  | 

### Return type

[**ApiEmpty**](ApiEmpty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

