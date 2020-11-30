library swagger.api;

import 'dart:async';
import 'dart:convert';
import 'package:http/browser_client.dart';
import 'package:http/http.dart';

part 'api_client.dart';
part 'api_helper.dart';
part 'api_exception.dart';
part 'auth/authentication.dart';
part 'auth/api_key_auth.dart';
part 'auth/oauth.dart';
part 'auth/http_basic_auth.dart';

part 'api/cicd_service_api.dart';

part 'model/api_empty.dart';
part 'model/api_get_templates_req.dart';
part 'model/api_get_variables_req.dart';
part 'model/api_job.dart';
part 'model/api_list_job_res.dart';
part 'model/api_list_task_res.dart';
part 'model/api_list_template_res.dart';
part 'model/api_list_variable_res.dart';
part 'model/api_run_task_req.dart';
part 'model/api_run_task_res.dart';
part 'model/api_task.dart';
part 'model/api_template.dart';
part 'model/api_variable.dart';
part 'model/protobuf_any.dart';
part 'model/runtime_error.dart';
part 'model/template_script_template.dart';


ApiClient defaultApiClient = new ApiClient();
