// @generated by protoc-gen-connect-query v0.6.0 with parameter "target=ts"
// @generated from file ipc/v1alpha1/ipc.proto (package ipc.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { PublishRequest, PublishResponse, SubscribeRequest, SubscribeResponse } from "./ipc_pb.js";
import { MethodKind } from "@bufbuild/protobuf";
import { createQueryService, createUnaryHooks, UnaryFunctionsWithHooks } from "@connectrpc/connect-query";

export const typeName = "ipc.v1alpha1.IPCService";

/**
 * @generated from service ipc.v1alpha1.IPCService
 */
export const IPCService = {
  typeName: "ipc.v1alpha1.IPCService",
  methods: {
    /**
     * @generated from rpc ipc.v1alpha1.IPCService.Publish
     */
    publish: {
      name: "Publish",
      I: PublishRequest,
      O: PublishResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc ipc.v1alpha1.IPCService.Subscribe
     */
    subscribe: {
      name: "Subscribe",
      I: SubscribeRequest,
      O: SubscribeResponse,
      kind: MethodKind.ServerStreaming,
    },
  }
} as const;

const $queryService = createQueryService({  service: IPCService,});

/**
 * @generated from rpc ipc.v1alpha1.IPCService.Publish
 */
export const publish: UnaryFunctionsWithHooks<PublishRequest, PublishResponse> = {   ...$queryService.publish,  ...createUnaryHooks($queryService.publish)};
