///
//  Generated code. Do not modify.
//  source: master/master_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'master_service.pb.dart' as $0;
export 'master_service.pb.dart';

class MasterGRPCClient extends $grpc.Client {
  static final _$getAllRegulations =
      $grpc.ClientMethod<$0.Empty, $0.GetAllRegulationsResponse>(
          '/master.v1.MasterGRPC/GetAllRegulations',
          ($0.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.GetAllRegulationsResponse.fromBuffer(value));
  static final _$getAllChapters =
      $grpc.ClientMethod<$0.GetAllChaptersRequest, $0.GetAllChaptersResponse>(
          '/master.v1.MasterGRPC/GetAllChapters',
          ($0.GetAllChaptersRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.GetAllChaptersResponse.fromBuffer(value));
  static final _$getAllParagraphs = $grpc.ClientMethod<
          $0.GetAllParagraphsRequest, $0.GetAllParagraphsResponse>(
      '/master.v1.MasterGRPC/GetAllParagraphs',
      ($0.GetAllParagraphsRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.GetAllParagraphsResponse.fromBuffer(value));
  static final _$editParagraph =
      $grpc.ClientMethod<$0.EditParagraphRequest, $0.Empty>(
          '/master.v1.MasterGRPC/EditParagraph',
          ($0.EditParagraphRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.Empty.fromBuffer(value));
  static final _$editAbsent =
      $grpc.ClientMethod<$0.EditAbsentRequest, $0.Empty>(
          '/master.v1.MasterGRPC/EditAbsent',
          ($0.EditAbsentRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.Empty.fromBuffer(value));
  static final _$createChapter =
      $grpc.ClientMethod<$0.CreateChapterRequest, $0.CreateChapterResponse>(
          '/master.v1.MasterGRPC/CreateChapter',
          ($0.CreateChapterRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.CreateChapterResponse.fromBuffer(value));
  static final _$createParagraphs =
      $grpc.ClientMethod<$0.CreateParagraphsRequest, $0.Empty>(
          '/master.v1.MasterGRPC/CreateParagraphs',
          ($0.CreateParagraphsRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.Empty.fromBuffer(value));
  static final _$createRegulation = $grpc.ClientMethod<
          $0.CreateRegulationRequest, $0.CreateRegulationResponse>(
      '/master.v1.MasterGRPC/CreateRegulation',
      ($0.CreateRegulationRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.CreateRegulationResponse.fromBuffer(value));
  static final _$generateLinks =
      $grpc.ClientMethod<$0.GenerateLinksRequest, $0.GenerateLinksResponse>(
          '/master.v1.MasterGRPC/GenerateLinks',
          ($0.GenerateLinksRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.GenerateLinksResponse.fromBuffer(value));
  static final _$deleteRegulation =
      $grpc.ClientMethod<$0.DeleteRegulationRequest, $0.Empty>(
          '/master.v1.MasterGRPC/DeleteRegulation',
          ($0.DeleteRegulationRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.Empty.fromBuffer(value));

  MasterGRPCClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.GetAllRegulationsResponse> getAllRegulations(
      $0.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getAllRegulations, request, options: options);
  }

  $grpc.ResponseFuture<$0.GetAllChaptersResponse> getAllChapters(
      $0.GetAllChaptersRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getAllChapters, request, options: options);
  }

  $grpc.ResponseFuture<$0.GetAllParagraphsResponse> getAllParagraphs(
      $0.GetAllParagraphsRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getAllParagraphs, request, options: options);
  }

  $grpc.ResponseFuture<$0.Empty> editParagraph($0.EditParagraphRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$editParagraph, request, options: options);
  }

  $grpc.ResponseFuture<$0.Empty> editAbsent($0.EditAbsentRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$editAbsent, request, options: options);
  }

  $grpc.ResponseFuture<$0.CreateChapterResponse> createChapter(
      $0.CreateChapterRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$createChapter, request, options: options);
  }

  $grpc.ResponseFuture<$0.Empty> createParagraphs(
      $0.CreateParagraphsRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$createParagraphs, request, options: options);
  }

  $grpc.ResponseFuture<$0.CreateRegulationResponse> createRegulation(
      $0.CreateRegulationRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$createRegulation, request, options: options);
  }

  $grpc.ResponseFuture<$0.GenerateLinksResponse> generateLinks(
      $0.GenerateLinksRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$generateLinks, request, options: options);
  }

  $grpc.ResponseFuture<$0.Empty> deleteRegulation(
      $0.DeleteRegulationRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$deleteRegulation, request, options: options);
  }
}

abstract class MasterGRPCServiceBase extends $grpc.Service {
  $core.String get $name => 'master.v1.MasterGRPC';

  MasterGRPCServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Empty, $0.GetAllRegulationsResponse>(
        'GetAllRegulations',
        getAllRegulations_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($0.GetAllRegulationsResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.GetAllChaptersRequest,
            $0.GetAllChaptersResponse>(
        'GetAllChapters',
        getAllChapters_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.GetAllChaptersRequest.fromBuffer(value),
        ($0.GetAllChaptersResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.GetAllParagraphsRequest,
            $0.GetAllParagraphsResponse>(
        'GetAllParagraphs',
        getAllParagraphs_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.GetAllParagraphsRequest.fromBuffer(value),
        ($0.GetAllParagraphsResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.EditParagraphRequest, $0.Empty>(
        'EditParagraph',
        editParagraph_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.EditParagraphRequest.fromBuffer(value),
        ($0.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.EditAbsentRequest, $0.Empty>(
        'EditAbsent',
        editAbsent_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.EditAbsentRequest.fromBuffer(value),
        ($0.Empty value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$0.CreateChapterRequest, $0.CreateChapterResponse>(
            'CreateChapter',
            createChapter_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.CreateChapterRequest.fromBuffer(value),
            ($0.CreateChapterResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.CreateParagraphsRequest, $0.Empty>(
        'CreateParagraphs',
        createParagraphs_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.CreateParagraphsRequest.fromBuffer(value),
        ($0.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.CreateRegulationRequest,
            $0.CreateRegulationResponse>(
        'CreateRegulation',
        createRegulation_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.CreateRegulationRequest.fromBuffer(value),
        ($0.CreateRegulationResponse value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$0.GenerateLinksRequest, $0.GenerateLinksResponse>(
            'GenerateLinks',
            generateLinks_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.GenerateLinksRequest.fromBuffer(value),
            ($0.GenerateLinksResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.DeleteRegulationRequest, $0.Empty>(
        'DeleteRegulation',
        deleteRegulation_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.DeleteRegulationRequest.fromBuffer(value),
        ($0.Empty value) => value.writeToBuffer()));
  }

  $async.Future<$0.GetAllRegulationsResponse> getAllRegulations_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async {
    return getAllRegulations(call, await request);
  }

  $async.Future<$0.GetAllChaptersResponse> getAllChapters_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.GetAllChaptersRequest> request) async {
    return getAllChapters(call, await request);
  }

  $async.Future<$0.GetAllParagraphsResponse> getAllParagraphs_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.GetAllParagraphsRequest> request) async {
    return getAllParagraphs(call, await request);
  }

  $async.Future<$0.Empty> editParagraph_Pre($grpc.ServiceCall call,
      $async.Future<$0.EditParagraphRequest> request) async {
    return editParagraph(call, await request);
  }

  $async.Future<$0.Empty> editAbsent_Pre($grpc.ServiceCall call,
      $async.Future<$0.EditAbsentRequest> request) async {
    return editAbsent(call, await request);
  }

  $async.Future<$0.CreateChapterResponse> createChapter_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.CreateChapterRequest> request) async {
    return createChapter(call, await request);
  }

  $async.Future<$0.Empty> createParagraphs_Pre($grpc.ServiceCall call,
      $async.Future<$0.CreateParagraphsRequest> request) async {
    return createParagraphs(call, await request);
  }

  $async.Future<$0.CreateRegulationResponse> createRegulation_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.CreateRegulationRequest> request) async {
    return createRegulation(call, await request);
  }

  $async.Future<$0.GenerateLinksResponse> generateLinks_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.GenerateLinksRequest> request) async {
    return generateLinks(call, await request);
  }

  $async.Future<$0.Empty> deleteRegulation_Pre($grpc.ServiceCall call,
      $async.Future<$0.DeleteRegulationRequest> request) async {
    return deleteRegulation(call, await request);
  }

  $async.Future<$0.GetAllRegulationsResponse> getAllRegulations(
      $grpc.ServiceCall call, $0.Empty request);
  $async.Future<$0.GetAllChaptersResponse> getAllChapters(
      $grpc.ServiceCall call, $0.GetAllChaptersRequest request);
  $async.Future<$0.GetAllParagraphsResponse> getAllParagraphs(
      $grpc.ServiceCall call, $0.GetAllParagraphsRequest request);
  $async.Future<$0.Empty> editParagraph(
      $grpc.ServiceCall call, $0.EditParagraphRequest request);
  $async.Future<$0.Empty> editAbsent(
      $grpc.ServiceCall call, $0.EditAbsentRequest request);
  $async.Future<$0.CreateChapterResponse> createChapter(
      $grpc.ServiceCall call, $0.CreateChapterRequest request);
  $async.Future<$0.Empty> createParagraphs(
      $grpc.ServiceCall call, $0.CreateParagraphsRequest request);
  $async.Future<$0.CreateRegulationResponse> createRegulation(
      $grpc.ServiceCall call, $0.CreateRegulationRequest request);
  $async.Future<$0.GenerateLinksResponse> generateLinks(
      $grpc.ServiceCall call, $0.GenerateLinksRequest request);
  $async.Future<$0.Empty> deleteRegulation(
      $grpc.ServiceCall call, $0.DeleteRegulationRequest request);
}
