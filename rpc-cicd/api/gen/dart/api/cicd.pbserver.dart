///
//  Generated code. Do not modify.
//  source: api/cicd.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'cicd.pb.dart' as $1;
import '../google/protobuf/empty.pb.dart' as $0;
import 'cicd.pbjson.dart';

export 'cicd.pb.dart';

abstract class CICDServiceBase extends $pb.GeneratedService {
  $async.Future<$1.RunTaskRes> runTask($pb.ServerContext ctx, $1.RunTaskReq request);
  $async.Future<$1.Task> getTask($pb.ServerContext ctx, $1.GetTaskReq request);
  $async.Future<$0.Empty> delTask($pb.ServerContext ctx, $1.DelTaskReq request);
  $async.Future<$0.Empty> putTask($pb.ServerContext ctx, $1.PutTaskReq request);
  $async.Future<$0.Empty> updateTask($pb.ServerContext ctx, $1.UpdateTaskReq request);
  $async.Future<$1.ListTaskRes> listTask($pb.ServerContext ctx, $1.ListTaskReq request);
  $async.Future<$1.Template> getTemplate($pb.ServerContext ctx, $1.GetTemplateReq request);
  $async.Future<$0.Empty> delTemplate($pb.ServerContext ctx, $1.DelTemplateReq request);
  $async.Future<$0.Empty> putTemplate($pb.ServerContext ctx, $1.PutTemplateReq request);
  $async.Future<$0.Empty> updateTemplate($pb.ServerContext ctx, $1.UpdateTemplateReq request);
  $async.Future<$1.ListTemplateRes> listTemplate($pb.ServerContext ctx, $1.ListTemplateReq request);
  $async.Future<$1.Variable> getVariable($pb.ServerContext ctx, $1.GetVariableReq request);
  $async.Future<$0.Empty> delVariable($pb.ServerContext ctx, $1.DelVariableReq request);
  $async.Future<$0.Empty> putVariable($pb.ServerContext ctx, $1.PutVariableReq request);
  $async.Future<$0.Empty> updateVariable($pb.ServerContext ctx, $1.UpdateVariableReq request);
  $async.Future<$1.ListVariableRes> listVariable($pb.ServerContext ctx, $1.ListVariableReq request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'RunTask': return $1.RunTaskReq();
      case 'GetTask': return $1.GetTaskReq();
      case 'DelTask': return $1.DelTaskReq();
      case 'PutTask': return $1.PutTaskReq();
      case 'UpdateTask': return $1.UpdateTaskReq();
      case 'ListTask': return $1.ListTaskReq();
      case 'GetTemplate': return $1.GetTemplateReq();
      case 'DelTemplate': return $1.DelTemplateReq();
      case 'PutTemplate': return $1.PutTemplateReq();
      case 'UpdateTemplate': return $1.UpdateTemplateReq();
      case 'ListTemplate': return $1.ListTemplateReq();
      case 'GetVariable': return $1.GetVariableReq();
      case 'DelVariable': return $1.DelVariableReq();
      case 'PutVariable': return $1.PutVariableReq();
      case 'UpdateVariable': return $1.UpdateVariableReq();
      case 'ListVariable': return $1.ListVariableReq();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'RunTask': return this.runTask(ctx, request);
      case 'GetTask': return this.getTask(ctx, request);
      case 'DelTask': return this.delTask(ctx, request);
      case 'PutTask': return this.putTask(ctx, request);
      case 'UpdateTask': return this.updateTask(ctx, request);
      case 'ListTask': return this.listTask(ctx, request);
      case 'GetTemplate': return this.getTemplate(ctx, request);
      case 'DelTemplate': return this.delTemplate(ctx, request);
      case 'PutTemplate': return this.putTemplate(ctx, request);
      case 'UpdateTemplate': return this.updateTemplate(ctx, request);
      case 'ListTemplate': return this.listTemplate(ctx, request);
      case 'GetVariable': return this.getVariable(ctx, request);
      case 'DelVariable': return this.delVariable(ctx, request);
      case 'PutVariable': return this.putVariable(ctx, request);
      case 'UpdateVariable': return this.updateVariable(ctx, request);
      case 'ListVariable': return this.listVariable(ctx, request);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => CICDServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => CICDServiceBase$messageJson;
}

