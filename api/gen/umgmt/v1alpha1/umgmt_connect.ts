// @generated by protoc-gen-connect-es v1.1.4 with parameter "target=ts"
// @generated from file umgmt/v1alpha1/umgmt.proto (package umgmt.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { ChangeMachineStateRequest, ChangeMachineStateResponse, ConfigureThermalFanProfilesRequest, ConfigureThermalFanProfilesResponse, ConfigureThermalSetPointsRequest, ConfigureThermalSetPointsResponse, GetInventoryRequest, GetInventoryResponse, GetMachineInfoRequest, GetMachineInfoResponse, GetMachineStateRequest, GetMachineStateResponse, GetSensorDataRequest, GetSensorDataResponse, GetSensorListRequest, GetSensorListResponse, GetUserInfoRequest, GetUserInfoResponse, GetUsersRequest, GetUsersResponse, StreamHostConsoleRequest, StreamHostConsoleResponse, UpdateUserRequest, UpdateUserResponse } from "./umgmt_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service umgmt.v1alpha1.UmgmtService
 */
export const UmgmtService = {
  typeName: "umgmt.v1alpha1.UmgmtService",
  methods: {
    /**
     * @generated from rpc umgmt.v1alpha1.UmgmtService.GetUsers
     */
    getUsers: {
      name: "GetUsers",
      I: GetUsersRequest,
      O: GetUsersResponse,
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
     * @generated from rpc umgmt.v1alpha1.UmgmtService.GetInventory
     */
    getInventory: {
      name: "GetInventory",
      I: GetInventoryRequest,
      O: GetInventoryResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc umgmt.v1alpha1.UmgmtService.GetMachineInfo
     */
    getMachineInfo: {
      name: "GetMachineInfo",
      I: GetMachineInfoRequest,
      O: GetMachineInfoResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc umgmt.v1alpha1.UmgmtService.GetMachineState
     */
    getMachineState: {
      name: "GetMachineState",
      I: GetMachineStateRequest,
      O: GetMachineStateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc umgmt.v1alpha1.UmgmtService.ChangeMachineState
     */
    changeMachineState: {
      name: "ChangeMachineState",
      I: ChangeMachineStateRequest,
      O: ChangeMachineStateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc umgmt.v1alpha1.UmgmtService.GetSensorList
     */
    getSensorList: {
      name: "GetSensorList",
      I: GetSensorListRequest,
      O: GetSensorListResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc umgmt.v1alpha1.UmgmtService.GetSensorData
     */
    getSensorData: {
      name: "GetSensorData",
      I: GetSensorDataRequest,
      O: GetSensorDataResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc umgmt.v1alpha1.UmgmtService.StreamHostConsole
     */
    streamHostConsole: {
      name: "StreamHostConsole",
      I: StreamHostConsoleRequest,
      O: StreamHostConsoleResponse,
      kind: MethodKind.BiDiStreaming,
    },
    /**
     * @generated from rpc umgmt.v1alpha1.UmgmtService.ConfigureThermalSetPoints
     */
    configureThermalSetPoints: {
      name: "ConfigureThermalSetPoints",
      I: ConfigureThermalSetPointsRequest,
      O: ConfigureThermalSetPointsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc umgmt.v1alpha1.UmgmtService.ConfigureThermalFanProfiles
     */
    configureThermalFanProfiles: {
      name: "ConfigureThermalFanProfiles",
      I: ConfigureThermalFanProfilesRequest,
      O: ConfigureThermalFanProfilesResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

