// @generated by protoc-gen-es v1.6.0 with parameter "target=ts"
// @generated from file umgmt/v1alpha1/inventory.proto (package umgmt.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { FieldMask, Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { PostalAddress } from "../../google/type/postal_address_pb.js";
import { Sensor } from "./sensor_pb.js";

/**
 * @generated from enum umgmt.v1alpha1.State
 */
export enum State {
  /**
   * @generated from enum value: STATE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: STATE_ON = 1;
   */
  ON = 1,

  /**
   * @generated from enum value: STATE_OFF = 2;
   */
  OFF = 2,

  /**
   * @generated from enum value: STATE_SUSPENDED = 3;
   */
  SUSPENDED = 3,

  /**
   * @generated from enum value: STATE_HIBERNATED = 4;
   */
  HIBERNATED = 4,

  /**
   * @generated from enum value: STATE_UNKNOWN = 5;
   */
  UNKNOWN = 5,
}
// Retrieve enum metadata with: proto3.getEnumType(State)
proto3.util.setEnumType(State, "umgmt.v1alpha1.State", [
  { no: 0, name: "STATE_UNSPECIFIED" },
  { no: 1, name: "STATE_ON" },
  { no: 2, name: "STATE_OFF" },
  { no: 3, name: "STATE_SUSPENDED" },
  { no: 4, name: "STATE_HIBERNATED" },
  { no: 5, name: "STATE_UNKNOWN" },
]);

/**
 * @generated from message umgmt.v1alpha1.Location
 */
export class Location extends Message<Location> {
  /**
   * @generated from field: google.type.PostalAddress address = 1;
   */
  address?: PostalAddress;

  /**
   * @generated from field: string room = 2;
   */
  room = "";

  /**
   * @generated from field: string rack = 3;
   */
  rack = "";

  /**
   * @generated from field: string slot = 4;
   */
  slot = "";

  /**
   * @generated from field: string notes = 5;
   */
  notes = "";

  constructor(data?: PartialMessage<Location>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "umgmt.v1alpha1.Location";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "address", kind: "message", T: PostalAddress },
    { no: 2, name: "room", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "rack", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "slot", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "notes", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Location {
    return new Location().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Location {
    return new Location().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Location {
    return new Location().fromJsonString(jsonString, options);
  }

  static equals(a: Location | PlainMessage<Location> | undefined, b: Location | PlainMessage<Location> | undefined): boolean {
    return proto3.util.equals(Location, a, b);
  }
}

/**
 * @generated from message umgmt.v1alpha1.HardwareInfo
 */
export class HardwareInfo extends Message<HardwareInfo> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: string type = 3;
   */
  type = "";

  /**
   * @generated from field: string serial_number = 4;
   */
  serialNumber = "";

  /**
   * @generated from field: string part_number = 5;
   */
  partNumber = "";

  /**
   * @generated from field: string sku = 6;
   */
  sku = "";

  /**
   * @generated from field: string manufacturer = 7;
   */
  manufacturer = "";

  /**
   * @generated from field: google.protobuf.Timestamp manufacture_date = 8;
   */
  manufactureDate?: Timestamp;

  /**
   * @generated from field: string asset_tag = 9;
   */
  assetTag = "";

  constructor(data?: PartialMessage<HardwareInfo>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "umgmt.v1alpha1.HardwareInfo";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "serial_number", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "part_number", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "sku", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "manufacturer", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "manufacture_date", kind: "message", T: Timestamp },
    { no: 9, name: "asset_tag", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): HardwareInfo {
    return new HardwareInfo().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): HardwareInfo {
    return new HardwareInfo().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): HardwareInfo {
    return new HardwareInfo().fromJsonString(jsonString, options);
  }

  static equals(a: HardwareInfo | PlainMessage<HardwareInfo> | undefined, b: HardwareInfo | PlainMessage<HardwareInfo> | undefined): boolean {
    return proto3.util.equals(HardwareInfo, a, b);
  }
}

/**
 * @generated from message umgmt.v1alpha1.Chassis
 */
export class Chassis extends Message<Chassis> {
  /**
   * @generated from field: umgmt.v1alpha1.HardwareInfo info = 1;
   */
  info?: HardwareInfo;

  /**
   * @generated from field: uint32 height_mm = 2;
   */
  heightMm = 0;

  /**
   * @generated from field: uint32 width_mm = 3;
   */
  widthMm = 0;

  /**
   * @generated from field: uint32 depth_mm = 4;
   */
  depthMm = 0;

  /**
   * @generated from field: uint32 weight_g = 5;
   */
  weightG = 0;

  /**
   * @generated from field: umgmt.v1alpha1.Location location = 6;
   */
  location?: Location;

  /**
   * @generated from field: umgmt.v1alpha1.System system = 7;
   */
  system?: System;

  /**
   * @generated from field: repeated umgmt.v1alpha1.Component components = 8;
   */
  components: Component[] = [];

  /**
   * @generated from field: repeated umgmt.v1alpha1.Sensor sensors = 9;
   */
  sensors: Sensor[] = [];

  /**
   * @generated from field: string notes = 10;
   */
  notes = "";

  constructor(data?: PartialMessage<Chassis>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "umgmt.v1alpha1.Chassis";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "info", kind: "message", T: HardwareInfo },
    { no: 2, name: "height_mm", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 3, name: "width_mm", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 4, name: "depth_mm", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 5, name: "weight_g", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 6, name: "location", kind: "message", T: Location },
    { no: 7, name: "system", kind: "message", T: System },
    { no: 8, name: "components", kind: "message", T: Component, repeated: true },
    { no: 9, name: "sensors", kind: "message", T: Sensor, repeated: true },
    { no: 10, name: "notes", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Chassis {
    return new Chassis().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Chassis {
    return new Chassis().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Chassis {
    return new Chassis().fromJsonString(jsonString, options);
  }

  static equals(a: Chassis | PlainMessage<Chassis> | undefined, b: Chassis | PlainMessage<Chassis> | undefined): boolean {
    return proto3.util.equals(Chassis, a, b);
  }
}

/**
 * @generated from message umgmt.v1alpha1.System
 */
export class System extends Message<System> {
  /**
   * @generated from field: umgmt.v1alpha1.HardwareInfo info = 1;
   */
  info?: HardwareInfo;

  /**
   * @generated from field: umgmt.v1alpha1.State state = 2;
   */
  state = State.UNSPECIFIED;

  /**
   * @generated from field: string notes = 3;
   */
  notes = "";

  constructor(data?: PartialMessage<System>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "umgmt.v1alpha1.System";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "info", kind: "message", T: HardwareInfo },
    { no: 2, name: "state", kind: "enum", T: proto3.getEnumType(State) },
    { no: 3, name: "notes", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): System {
    return new System().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): System {
    return new System().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): System {
    return new System().fromJsonString(jsonString, options);
  }

  static equals(a: System | PlainMessage<System> | undefined, b: System | PlainMessage<System> | undefined): boolean {
    return proto3.util.equals(System, a, b);
  }
}

/**
 * @generated from message umgmt.v1alpha1.Component
 */
export class Component extends Message<Component> {
  /**
   * @generated from field: umgmt.v1alpha1.HardwareInfo info = 1;
   */
  info?: HardwareInfo;

  /**
   * @generated from field: string notes = 2;
   */
  notes = "";

  constructor(data?: PartialMessage<Component>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "umgmt.v1alpha1.Component";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "info", kind: "message", T: HardwareInfo },
    { no: 2, name: "notes", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Component {
    return new Component().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Component {
    return new Component().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Component {
    return new Component().fromJsonString(jsonString, options);
  }

  static equals(a: Component | PlainMessage<Component> | undefined, b: Component | PlainMessage<Component> | undefined): boolean {
    return proto3.util.equals(Component, a, b);
  }
}

/**
 * @generated from message umgmt.v1alpha1.Inventory
 */
export class Inventory extends Message<Inventory> {
  /**
   * @generated from field: repeated umgmt.v1alpha1.Chassis chassis = 1;
   */
  chassis: Chassis[] = [];

  /**
   * @generated from field: repeated umgmt.v1alpha1.System systems = 2;
   */
  systems: System[] = [];

  /**
   * @generated from field: repeated umgmt.v1alpha1.Component components = 3;
   */
  components: Component[] = [];

  constructor(data?: PartialMessage<Inventory>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "umgmt.v1alpha1.Inventory";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "chassis", kind: "message", T: Chassis, repeated: true },
    { no: 2, name: "systems", kind: "message", T: System, repeated: true },
    { no: 3, name: "components", kind: "message", T: Component, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Inventory {
    return new Inventory().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Inventory {
    return new Inventory().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Inventory {
    return new Inventory().fromJsonString(jsonString, options);
  }

  static equals(a: Inventory | PlainMessage<Inventory> | undefined, b: Inventory | PlainMessage<Inventory> | undefined): boolean {
    return proto3.util.equals(Inventory, a, b);
  }
}

/**
 * @generated from message umgmt.v1alpha1.ListInventoryRequest
 */
export class ListInventoryRequest extends Message<ListInventoryRequest> {
  /**
   * @generated from field: google.protobuf.FieldMask field_mask = 1;
   */
  fieldMask?: FieldMask;

  constructor(data?: PartialMessage<ListInventoryRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "umgmt.v1alpha1.ListInventoryRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "field_mask", kind: "message", T: FieldMask },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListInventoryRequest {
    return new ListInventoryRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListInventoryRequest {
    return new ListInventoryRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListInventoryRequest {
    return new ListInventoryRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListInventoryRequest | PlainMessage<ListInventoryRequest> | undefined, b: ListInventoryRequest | PlainMessage<ListInventoryRequest> | undefined): boolean {
    return proto3.util.equals(ListInventoryRequest, a, b);
  }
}

/**
 * @generated from message umgmt.v1alpha1.ListInventoryResponse
 */
export class ListInventoryResponse extends Message<ListInventoryResponse> {
  /**
   * @generated from field: umgmt.v1alpha1.Inventory inventory = 1;
   */
  inventory?: Inventory;

  constructor(data?: PartialMessage<ListInventoryResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "umgmt.v1alpha1.ListInventoryResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "inventory", kind: "message", T: Inventory },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListInventoryResponse {
    return new ListInventoryResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListInventoryResponse {
    return new ListInventoryResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListInventoryResponse {
    return new ListInventoryResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListInventoryResponse | PlainMessage<ListInventoryResponse> | undefined, b: ListInventoryResponse | PlainMessage<ListInventoryResponse> | undefined): boolean {
    return proto3.util.equals(ListInventoryResponse, a, b);
  }
}

/**
 * @generated from message umgmt.v1alpha1.GetChassisInfoRequest
 */
export class GetChassisInfoRequest extends Message<GetChassisInfoRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: google.protobuf.FieldMask field_mask = 2;
   */
  fieldMask?: FieldMask;

  constructor(data?: PartialMessage<GetChassisInfoRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "umgmt.v1alpha1.GetChassisInfoRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "field_mask", kind: "message", T: FieldMask },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetChassisInfoRequest {
    return new GetChassisInfoRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetChassisInfoRequest {
    return new GetChassisInfoRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetChassisInfoRequest {
    return new GetChassisInfoRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetChassisInfoRequest | PlainMessage<GetChassisInfoRequest> | undefined, b: GetChassisInfoRequest | PlainMessage<GetChassisInfoRequest> | undefined): boolean {
    return proto3.util.equals(GetChassisInfoRequest, a, b);
  }
}

/**
 * @generated from message umgmt.v1alpha1.GetChassisInfoResponse
 */
export class GetChassisInfoResponse extends Message<GetChassisInfoResponse> {
  /**
   * @generated from field: umgmt.v1alpha1.Chassis chassis = 1;
   */
  chassis?: Chassis;

  constructor(data?: PartialMessage<GetChassisInfoResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "umgmt.v1alpha1.GetChassisInfoResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "chassis", kind: "message", T: Chassis },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetChassisInfoResponse {
    return new GetChassisInfoResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetChassisInfoResponse {
    return new GetChassisInfoResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetChassisInfoResponse {
    return new GetChassisInfoResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetChassisInfoResponse | PlainMessage<GetChassisInfoResponse> | undefined, b: GetChassisInfoResponse | PlainMessage<GetChassisInfoResponse> | undefined): boolean {
    return proto3.util.equals(GetChassisInfoResponse, a, b);
  }
}

/**
 * @generated from message umgmt.v1alpha1.ChangeSystemStateRequest
 */
export class ChangeSystemStateRequest extends Message<ChangeSystemStateRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: umgmt.v1alpha1.State state = 2;
   */
  state = State.UNSPECIFIED;

  constructor(data?: PartialMessage<ChangeSystemStateRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "umgmt.v1alpha1.ChangeSystemStateRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "state", kind: "enum", T: proto3.getEnumType(State) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ChangeSystemStateRequest {
    return new ChangeSystemStateRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ChangeSystemStateRequest {
    return new ChangeSystemStateRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ChangeSystemStateRequest {
    return new ChangeSystemStateRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ChangeSystemStateRequest | PlainMessage<ChangeSystemStateRequest> | undefined, b: ChangeSystemStateRequest | PlainMessage<ChangeSystemStateRequest> | undefined): boolean {
    return proto3.util.equals(ChangeSystemStateRequest, a, b);
  }
}

/**
 * @generated from message umgmt.v1alpha1.ChangeSystemStateResponse
 */
export class ChangeSystemStateResponse extends Message<ChangeSystemStateResponse> {
  constructor(data?: PartialMessage<ChangeSystemStateResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "umgmt.v1alpha1.ChangeSystemStateResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ChangeSystemStateResponse {
    return new ChangeSystemStateResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ChangeSystemStateResponse {
    return new ChangeSystemStateResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ChangeSystemStateResponse {
    return new ChangeSystemStateResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ChangeSystemStateResponse | PlainMessage<ChangeSystemStateResponse> | undefined, b: ChangeSystemStateResponse | PlainMessage<ChangeSystemStateResponse> | undefined): boolean {
    return proto3.util.equals(ChangeSystemStateResponse, a, b);
  }
}
