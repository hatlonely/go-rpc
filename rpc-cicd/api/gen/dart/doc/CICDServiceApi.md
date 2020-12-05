# openapi.api.CICDServiceApi

## Load the API package
```dart
import 'package:openapi/api.dart';
```

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**cICDServiceDelJob**](CICDServiceApi.md#cICDServiceDelJob) | **DELETE** /v1/job/{id} | 
[**cICDServiceDelTask**](CICDServiceApi.md#cICDServiceDelTask) | **DELETE** /v1/task/{id} | 
[**cICDServiceDelTemplate**](CICDServiceApi.md#cICDServiceDelTemplate) | **DELETE** /v1/template/{id} | 
[**cICDServiceDelVariable**](CICDServiceApi.md#cICDServiceDelVariable) | **DELETE** /v1/variable/{id} | 
[**cICDServiceGetJob**](CICDServiceApi.md#cICDServiceGetJob) | **GET** /v1/job/{id} | 
[**cICDServiceGetTask**](CICDServiceApi.md#cICDServiceGetTask) | **GET** /v1/task/{id} | 
[**cICDServiceGetTemplate**](CICDServiceApi.md#cICDServiceGetTemplate) | **GET** /v1/template/{id} | 
[**cICDServiceGetTemplates**](CICDServiceApi.md#cICDServiceGetTemplates) | **POST** /v1/getTemplates | 
[**cICDServiceGetVariable**](CICDServiceApi.md#cICDServiceGetVariable) | **GET** /v1/variable/{id} | 
[**cICDServiceGetVariables**](CICDServiceApi.md#cICDServiceGetVariables) | **POST** /v1/getVariables | 
[**cICDServiceListJob**](CICDServiceApi.md#cICDServiceListJob) | **GET** /v1/job | 
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


# **cICDServiceDelJob**
> Object cICDServiceDelJob(id)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final id = id_example; // String | 

try { 
    final result = api_instance.cICDServiceDelJob(id);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServiceDelJob: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceDelTask**
> Object cICDServiceDelTask(id)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final id = id_example; // String | 

try { 
    final result = api_instance.cICDServiceDelTask(id);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServiceDelTask: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceDelTemplate**
> Object cICDServiceDelTemplate(id)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final id = id_example; // String | 

try { 
    final result = api_instance.cICDServiceDelTemplate(id);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServiceDelTemplate: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceDelVariable**
> Object cICDServiceDelVariable(id)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final id = id_example; // String | 

try { 
    final result = api_instance.cICDServiceDelVariable(id);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServiceDelVariable: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceGetJob**
> ApiJob cICDServiceGetJob(id)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final id = id_example; // String | 

try { 
    final result = api_instance.cICDServiceGetJob(id);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServiceGetJob: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

### Return type

[**ApiJob**](ApiJob.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceGetTask**
> ApiTask cICDServiceGetTask(id)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final id = id_example; // String | 

try { 
    final result = api_instance.cICDServiceGetTask(id);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServiceGetTask: $e\n');
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

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceGetTemplate**
> ApiTemplate cICDServiceGetTemplate(id)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final id = id_example; // String | 

try { 
    final result = api_instance.cICDServiceGetTemplate(id);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServiceGetTemplate: $e\n');
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

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceGetTemplates**
> ApiListTemplateRes cICDServiceGetTemplates(body)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final body = ApiGetTemplatesReq(); // ApiGetTemplatesReq | 

try { 
    final result = api_instance.cICDServiceGetTemplates(body);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServiceGetTemplates: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiGetTemplatesReq**](ApiGetTemplatesReq.md)|  | 

### Return type

[**ApiListTemplateRes**](ApiListTemplateRes.md)

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
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final id = id_example; // String | 

try { 
    final result = api_instance.cICDServiceGetVariable(id);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServiceGetVariable: $e\n');
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

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceGetVariables**
> ApiListVariableRes cICDServiceGetVariables(body)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final body = ApiGetVariablesReq(); // ApiGetVariablesReq | 

try { 
    final result = api_instance.cICDServiceGetVariables(body);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServiceGetVariables: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiGetVariablesReq**](ApiGetVariablesReq.md)|  | 

### Return type

[**ApiListVariableRes**](ApiListVariableRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceListJob**
> ApiListJobRes cICDServiceListJob(offset, limit, taskID)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final offset = offset_example; // String | 
final limit = limit_example; // String | 
final taskID = taskID_example; // String | 

try { 
    final result = api_instance.cICDServiceListJob(offset, limit, taskID);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServiceListJob: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **String**|  | [optional] 
 **limit** | **String**|  | [optional] 
 **taskID** | **String**|  | [optional] 

### Return type

[**ApiListJobRes**](ApiListJobRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceListTask**
> ApiListTaskRes cICDServiceListTask(offset, limit, brief)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final offset = offset_example; // String | 
final limit = limit_example; // String | 
final brief = true; // bool | 

try { 
    final result = api_instance.cICDServiceListTask(offset, limit, brief);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServiceListTask: $e\n');
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

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceListTemplate**
> ApiListTemplateRes cICDServiceListTemplate(offset, limit, brief)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final offset = offset_example; // String | 
final limit = limit_example; // String | 
final brief = true; // bool | 

try { 
    final result = api_instance.cICDServiceListTemplate(offset, limit, brief);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServiceListTemplate: $e\n');
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

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceListVariable**
> ApiListVariableRes cICDServiceListVariable(offset, limit, brief)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final offset = offset_example; // String | 
final limit = limit_example; // String | 
final brief = true; // bool | 

try { 
    final result = api_instance.cICDServiceListVariable(offset, limit, brief);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServiceListVariable: $e\n');
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

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServicePutTask**
> Object cICDServicePutTask(body)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final body = ApiTask(); // ApiTask | 

try { 
    final result = api_instance.cICDServicePutTask(body);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServicePutTask: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiTask**](ApiTask.md)|  | 

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServicePutTemplate**
> Object cICDServicePutTemplate(body)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final body = ApiTemplate(); // ApiTemplate | 

try { 
    final result = api_instance.cICDServicePutTemplate(body);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServicePutTemplate: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiTemplate**](ApiTemplate.md)|  | 

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServicePutVariable**
> Object cICDServicePutVariable(body)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final body = ApiVariable(); // ApiVariable | 

try { 
    final result = api_instance.cICDServicePutVariable(body);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServicePutVariable: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiVariable**](ApiVariable.md)|  | 

### Return type

[**Object**](Object.md)

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
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final body = ApiRunTaskReq(); // ApiRunTaskReq | 

try { 
    final result = api_instance.cICDServiceRunTask(body);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServiceRunTask: $e\n');
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
> Object cICDServiceUpdateTask(taskId, body)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final taskId = taskId_example; // String | 
final body = ApiTask(); // ApiTask | 

try { 
    final result = api_instance.cICDServiceUpdateTask(taskId, body);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServiceUpdateTask: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskId** | **String**|  | 
 **body** | [**ApiTask**](ApiTask.md)|  | 

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceUpdateTemplate**
> Object cICDServiceUpdateTemplate(templateId, body)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final templateId = templateId_example; // String | 
final body = ApiTemplate(); // ApiTemplate | 

try { 
    final result = api_instance.cICDServiceUpdateTemplate(templateId, body);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServiceUpdateTemplate: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateId** | **String**|  | 
 **body** | [**ApiTemplate**](ApiTemplate.md)|  | 

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cICDServiceUpdateVariable**
> Object cICDServiceUpdateVariable(variableId, body)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = CICDServiceApi();
final variableId = variableId_example; // String | 
final body = ApiVariable(); // ApiVariable | 

try { 
    final result = api_instance.cICDServiceUpdateVariable(variableId, body);
    print(result);
} catch (e) {
    print('Exception when calling CICDServiceApi->cICDServiceUpdateVariable: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableId** | **String**|  | 
 **body** | [**ApiVariable**](ApiVariable.md)|  | 

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

