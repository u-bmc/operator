// @generated by protoc-gen-connect-es v1.3.0 with parameter "target=ts"
// @generated from file umgmt/v1alpha1/umgmt.proto (package umgmt.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { DeleteUserRequest, DeleteUserResponse, GetUserInfoRequest, GetUserInfoResponse, ListUsersRequest, ListUsersResponse, UpdateUserRequest, UpdateUserResponse } from "./user_pb.js";
import { MethodKind } from "@bufbuild/protobuf";
import { ChangeSystemStateRequest, ChangeSystemStateResponse, GetChassisInfoRequest, GetChassisInfoResponse, ListInventoryRequest, ListInventoryResponse } from "./inventory_pb.js";
import { StreamConsoleRequest, StreamConsoleResponse } from "./console_pb.js";

/**
 * @generated from service umgmt.v1alpha1.UmgmtService
 */
export const UmgmtService = {
  typeName: "umgmt.v1alpha1.UmgmtService",
  methods: {
    /**
     * @generated from rpc umgmt.v1alpha1.UmgmtService.ListUsers
     */
    listUsers: {
      name: "ListUsers",
      I: ListUsersRequest,
      O: ListUsersResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc umgmt.v1alpha1.UmgmtService.GetUserInfo
     */
    getUserInfo: {
      name: "GetUserInfo",
      I: GetUserInfoRequest,
      O: GetUserInfoResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc umgmt.v1alpha1.UmgmtService.UpdateUser
     */
    updateUser: {
      name: "UpdateUser",
      I: UpdateUserRequest,
      O: UpdateUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc umgmt.v1alpha1.UmgmtService.DeleteUser
     */
    deleteUser: {
      name: "DeleteUser",
      I: DeleteUserRequest,
      O: DeleteUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc umgmt.v1alpha1.UmgmtService.ListInventory
     */
    listInventory: {
      name: "ListInventory",
      I: ListInventoryRequest,
      O: ListInventoryResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc umgmt.v1alpha1.UmgmtService.GetChassisInfo
     */
    getChassisInfo: {
      name: "GetChassisInfo",
      I: GetChassisInfoRequest,
      O: GetChassisInfoResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc umgmt.v1alpha1.UmgmtService.ChangeSystemState
     */
    changeSystemState: {
      name: "ChangeSystemState",
      I: ChangeSystemStateRequest,
      O: ChangeSystemStateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc umgmt.v1alpha1.UmgmtService.StreamConsole
     */
    streamConsole: {
      name: "StreamConsole",
      I: StreamConsoleRequest,
      O: StreamConsoleResponse,
      kind: MethodKind.BiDiStreaming,
    },
  }
} as const;

